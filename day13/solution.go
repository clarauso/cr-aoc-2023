package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/clarauso/cr-aoc-2023/utils"
)

func PointOfIncidence(inputFilePath string) (int, int) {

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

			l2, a2 := evalPatternWithSmudge(pattern)
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

	l2, a2 := evalPatternWithSmudge(pattern)
	sol2 += l2
	sol2 += a2

	// 27502
	fmt.Printf("%d\n", sol1)
	// 31102 too low, 37350 and 35846 too high, 31947...
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

	return val, aboveVal

}

func evalPatternWithSmudge(pattern [][]rune) (int, int) {

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
	var flips int
	for i := 1; i < cols; i++ {
		flips = 1
		distance := stringDistance(columns[i-1], columns[i])
		if distance <= flips {
			found := true
			flips -= distance

			for j := 1; i+j < cols && i-j-1 >= 0; j++ {
				distance = stringDistance(columns[i+j], columns[i-j-1])
				if distance > flips {
					found = false
					break
				}
				flips -= distance
			}
			if found && flips == 0 {
				val = i
			}

		}
	}

	// check rows

	aboveVal := 0
	for i := 1; i < rows; i++ {
		flips = 1
		distance := stringDistance(string(pattern[i-1]), string(pattern[i]))
		if distance <= flips {
			found := true
			flips -= distance

			for j := 1; i+j < rows && i-j-1 >= 0; j++ {
				distance = stringDistance(string(pattern[i+j]), string(pattern[i-j-1]))
				if distance > flips {
					found = false
					break
				}
				flips -= distance
			}
			if found && flips == 0 {
				aboveVal = i * 100
				break
			}

		}
	}

	return val, aboveVal

}

func stringDistance(a string, b string) int {

	distance := 0
	for i, _ := range a {
		if a[i] != b[i] {
			distance++

		}
	}

	return distance

}
