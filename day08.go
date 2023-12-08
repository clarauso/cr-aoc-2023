package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Node struct {
	label string
	left  *Node
	right *Node
}

func wasteland(inputFilePath string) (int, int) {

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
			fmt.Println(instructions)
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
	fmt.Println(sol1)

	// sol 2
	sol2 := 0

	return sol1, sol2

}

func parseWastelandInstructions(line string) []rune {
	return mapToRuneSlice(line)
}

func parseWastelandNode(line string) (Node, string, string) {

	wastelandInputLineRegex := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)

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

			switch r {
			case 'L':
				current = current.left
			case 'R':
				current = current.right
			}

			steps++

		}
	}

	return steps
}
