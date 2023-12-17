package day08

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/clarauso/cr-aoc-2023/utils"
)

type Node struct {
	label string
	left  *Node
	right *Node
}

func Wasteland(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []rune
	linesRead := 0
	nodesMap := make(map[string]*Node, 0)
	leftMap := make(map[string]string)
	rightMap := make(map[string]string)

	for scanner.Scan() {
		currentLine := scanner.Text()

		switch {
		case linesRead == 0:
			instructions = parseWastelandInstructions(currentLine)
		case len(currentLine) == 0:
			continue
		default:
			node, left, right := parseWastelandNode(currentLine)
			nodesMap[node.label] = &node
			leftMap[node.label] = left
			rightMap[node.label] = right
		}

		linesRead++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	setPointers(nodesMap, leftMap, rightMap)

	// sol 1
	sol1 := goTo(nodesMap["AAA"], "ZZZ", instructions)

	// sol 2
	startingSli := getStartingNodes(nodesMap)
	endingSli := getEndingNodes(nodesMap)
	sol2 := goToList(startingSli, endingSli, instructions)

	return sol1, sol2

}

func parseWastelandInstructions(line string) []rune {
	return utils.MapToRuneSlice(line)
}

func parseWastelandNode(line string) (Node, string, string) {

	wastelandInputLineRegex := regexp.MustCompile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`)

	sli := wastelandInputLineRegex.FindStringSubmatch(line)

	node := Node{label: sli[1]}

	return node, sli[2], sli[3]
}

func setPointers(nodesMap map[string]*Node, leftMap map[string]string, rightMap map[string]string) {

	for key, value := range nodesMap {
		value.left = nodesMap[leftMap[key]]
		value.right = nodesMap[rightMap[key]]
	}
}

func goTo(start *Node, endLabel string, instructions []rune) int {

	steps := 0

	current := start

	for current.label != endLabel {
		for _, r := range instructions {
			current = nextNode(current, r)
			steps++
		}
	}

	return steps
}

func goToList(start []*Node, end []*Node, instructions []rune) int {

	lengths := make([]int, 0)

	for _, s := range start {

		for _, e := range end {
			conn := areConnected(*s, *e)

			if conn {
				fmt.Printf("Going from %s to %s is possible\n", s.label, e.label)

				n := goTo(s, e.label, instructions)
				lengths = append(lengths, n)
			}

		}

	}

	steps := utils.LeastCommonMultiple(lengths[0], lengths[1], lengths[2:]...)

	return steps
}

func nextNode(current *Node, direction rune) *Node {

	switch direction {
	case 'L':
		current = current.left
	case 'R':
		current = current.right
	}

	return current
}

func filter(nodesMap map[string]*Node, predicate func(Node) bool) []*Node {

	out := make([]*Node, 0)
	for _, n := range nodesMap {
		if predicate(*n) {
			out = append(out, n)
		}
	}

	return out

}

func getStartingNodes(nodesMap map[string]*Node) []*Node {

	predicate := func(n Node) bool {
		return strings.LastIndex(n.label, "A") == len(n.label)-1
	}
	return filter(nodesMap, predicate)

}

func getEndingNodes(nodesMap map[string]*Node) []*Node {

	predicate := func(n Node) bool {
		return strings.LastIndex(n.label, "Z") == len(n.label)-1
	}
	return filter(nodesMap, predicate)

}

// Checks if a path exists between a and b. Based on https://stackoverflow.com/a/354366/2924050
func areConnected(a Node, b Node) bool {

	toDo := make(map[string]Node, 0)
	done := make(map[string]Node, 0)

	toDo[a.label] = a
	for len(toDo) > 0 {
		current := utils.Pop(toDo)
		done[current.label] = current
		if current.left.label == b.label || current.right.label == b.label {
			return true
		}
		_, x1 := done[current.left.label]
		if !x1 {
			toDo[current.left.label] = *current.left
		}
		_, x2 := done[current.right.label]
		if !x2 {
			toDo[current.right.label] = *current.right
		}

	}

	return false
}
