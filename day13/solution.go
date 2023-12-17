package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/clarauso/cr-aoc-2023/utils"
)

func pointOfIncidence(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := make([][]rune, 0)
	sol1 := 0
	sol2 := 0

	for scanner.Scan() {
		currentLine := scanner.Text()

		switch len(currentLine) {
		case 0:
			left, above := evalPattern(pattern)
			sol1 += left
			sol1 += above
			pp := fixSmudge(pattern)
			l2, a2 := evalPattern(pp)
			sol2 += l2
			sol2 += a2
			// reset pattern
			pattern = make([][]rune, 0)
		default:
			pattern = append(pattern, utils.MapToRuneSlice(currentLine))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	left, above := evalPattern(pattern)
	sol1 += left
	sol1 += above
	pp := fixSmudge(pattern)
	l2, a2 := evalPattern(pp)
	sol2 += l2
	sol2 += a2

	// 27502
	fmt.Printf("%d\n", sol1)
	// 31102 too low
	fmt.Printf("%d\n", sol2)

	return sol1, sol2

}

func evalPattern(pattern [][]rune) (int, int) {

	// check on columns
	cols := len(pattern[0])
	rows := len(pattern)

	columns := make([]string, 0)
	for c := 0; c < cols; c++ {
		column := ""
		for j := 0; j < rows; j++ {
			column = column + string(pattern[j][c])
		}
		columns = append(columns, column)

	}

	val := 0
	for i := 1; i < cols; i++ {
		if columns[i-1] == columns[i] {

			found := true
			for j := 1; i+j < cols && i-j-1 >= 0; j++ {
				if columns[i+j] != columns[i-j-1] {
					found = false
					break
				}
			}
			if found {
				val = i
			}

		}
	}

	// check rows

	aboveVal := 0
	for i := 1; i < rows; i++ {
		if string(pattern[i-1]) == string(pattern[i]) {

			found := true
			for j := 1; i+j < rows && i-j-1 >= 0; j++ {
				if string(pattern[i+j]) != string(pattern[i-j-1]) {
					found = false
					break
				}
			}
			if found {
				aboveVal = i * 100
				break
			}

		}
	}

	fmt.Printf("This has %d %d\n", val, aboveVal)
	for _, p := range pattern {
		fmt.Println(string(p))
	}
	fmt.Println("===")

	return val, aboveVal

}

func fixSmudge(pattern [][]rune) [][]rune {

	rows := len(pattern)

	for i := 0; i < rows; i++ {
		for j := i + 1; j < rows; j++ {
			distances := patternDistance(pattern[i], pattern[j])
			if len(distances) == 1 {
				idx := distances[0]
				fmt.Printf("Pattern %s changing row %d[%d]\n", string(pattern[0]), i, idx)
				current := pattern[i][idx]
				if current == '.' {
					pattern[i][idx] = '#'
				} else {
					pattern[i][idx] = '.'
				}
				// exactly one smudge, so we can stop
				return pattern
			}
		}
	}

	cols := len(pattern[0])
	columns := make([][]rune, 0)
	for c := 0; c < cols; c++ {
		column := make([]rune, rows)
		for j := 0; j < rows; j++ {
			column[j] = pattern[j][c]
		}
		columns = append(columns, column)

	}

	fmt.Println("Processing cols")
	for i := 0; i < cols; i++ {
		for j := i + 1; j < cols; j++ {
			distances := patternDistance(columns[i], columns[j])
			if len(distances) == 1 {
				idx := distances[0]
				fmt.Printf("Pattern %s changing col %d[%d]\n", string(pattern[0]), i, idx)
				current := pattern[idx][i]
				if current == '.' {
					pattern[idx][i] = '#'
				} else {
					pattern[idx][i] = '.'
				}
				return pattern
			}
		}
	}

	return pattern
}

func patternDistance(a []rune, b []rune) []int {

	diffIndexes := make([]int, 0)
	for i, _ := range a {
		if a[i] != b[i] {
			diffIndexes = append(diffIndexes, i)
		}
	}

	return diffIndexes

}
