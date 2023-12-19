package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Part struct {
	x, m, a, s int
}

type Workflow struct {
	name  string
	rules []Rule
}

//px{a<2006:qkq,m>2090:A,rfg}
//pv{a>1716:R,A}
//lnx{m>1548:A,A}
//rfg{s<537:gd,x>2440:R,A}
//qs{s>3448:A,lnx}
//qkq{x<1416:A,crn}
//crn{x>2662:A,R}
//in{s<1351:px,qqz}
//qqz{s>2770:qs,m<1801:hdj,R}
//gd{a>3333:R,R}
//hdj{m>838:A,pv}

// field operator value, value operator OR outcome
// workflow name
// OUTCOME
type Rule struct {
	field        rune
	operator     rune
	threshold    int
	nextWorkflow string
	outcome      rune
}

var partRegex = regexp.MustCompile(`{x=([0-9]{1,}),m=([0-9]{1,}),a=([0-9]{1,}),s=([0-9]{1,})}`)
var workflowRegex = regexp.MustCompile(`([a-z]{2,3})\{(.*),([ARa-z]{1,})\}`)
var ruleRegex = regexp.MustCompile(`([xmas])([<>])([0-9]{1,}):([ARa-z]{1,})`)

func Aplenty(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	readParts := false

	workflows := make(map[string]Workflow, 0)
	parts := make([]Part, 0)

	for scanner.Scan() {
		currentLine := scanner.Text()

		switch {
		case len(currentLine) == 0:
			readParts = true
		case !readParts:
			wf := parseWorkflow(currentLine)
			workflows[wf.name] = wf
		case readParts:
			parts = append(parts, parsePart(currentLine))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sol1, sol2 := 0, 0
	startWf := workflows["in"]

	for _, p := range parts {
		if startWf.accept(p, workflows) {
			sol1 += p.ratingSum()
			fmt.Printf("Accepted %v\n", p)
		}
	}

	//grid := make([][]rune, 0)
	//for i,d := range directions {
	//}

	fmt.Printf("%d %d", sol1, sol2)
	return sol1, sol2

}

func parseWorkflow(line string) Workflow {

	wf := Workflow{}

	matches := workflowRegex.FindStringSubmatch(line)
	wf.name = matches[1]
	wf.rules = make([]Rule, 0)

	// a single string containing all rules
	allRules := matches[2]
	// splutted string rules
	strRulesSli := strings.Split(allRules, ",")

	for _, s := range strRulesSli {
		rule := Rule{}
		ruleParts := ruleRegex.FindStringSubmatch(s)

		rule.field = rune(ruleParts[1][0])
		rule.operator = rune(ruleParts[2][0])
		rule.threshold, _ = strconv.Atoi(ruleParts[3])

		switch ruleParts[4] {
		case "A", "R":
			rule.outcome = rune(ruleParts[4][0])
		default:
			rule.nextWorkflow = ruleParts[4]
		}

		wf.rules = append(wf.rules, rule)

	}

	lastRule := Rule{}
	lastIdx := len(matches) - 1
	switch matches[lastIdx] {
	case "A", "R":
		lastRule.outcome = rune(matches[lastIdx][0])
	default:
		lastRule.nextWorkflow = matches[lastIdx]
	}

	wf.rules = append(wf.rules, lastRule)

	return wf

}

func parsePart(line string) Part {

	matches := partRegex.FindStringSubmatch(line)
	part := Part{}

	for i := 1; i < len(matches); i++ {
		switch i {
		case 1:
			part.x, _ = strconv.Atoi(matches[i])
		case 2:
			part.m, _ = strconv.Atoi(matches[i])
		case 3:
			part.a, _ = strconv.Atoi(matches[i])
		case 4:
			part.s, _ = strconv.Atoi(matches[i])
		}
	}

	return part
}

func (w Workflow) accept(p Part, workflows map[string]Workflow) bool {

	// Rule "x>10:one": If the part's x is more than 10, send the part to the workflow named one.
	// Rule "m<20:two": Otherwise, if the part's m is less than 20, send the part to the workflow named two.
	// Rule "a>30:R": Otherwise, if the part's a is more than 30, the part is immediately rejected (R).
	// Rule "A": Otherwise, because no other rules matched the part, the part is immediately accepted (A).

	for _, rule := range w.rules {

		outcome, workflow := rule.evaluate(p)
		if outcome == 'A' {
			return true
		} else if outcome == 'R' {
			return false
		} else if workflow != "" {
			return workflows[workflow].accept(p, workflows)
		}
	}

	return false
}

func (r Rule) evaluate(p Part) (rune, string) {

	if r.operator == 0 {
		if r.outcome == 'A' || r.outcome == 'R' {
			return r.outcome, ""
		} else {
			return 'C', r.nextWorkflow
		}
	}

	value := p.extractField(r.field)

	accept := false
	switch r.operator {
	case '>':
		accept = value > r.threshold
	default:
		accept = value < r.threshold
	}

	if accept {
		if r.outcome == 'A' || r.outcome == 'R' {
			return r.outcome, ""
		} else {
			return 'C', r.nextWorkflow
		}
	}

	return 'C', ""
}

func (p Part) extractField(fieldName rune) int {

	switch fieldName {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	default:
		return p.s
	}

}

func (p Part) ratingSum() int {
	return p.x + p.m + p.a + p.s
}
