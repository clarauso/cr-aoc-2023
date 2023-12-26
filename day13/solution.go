package day13

import (
	"bufio"
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
	pattern := make([]string, 0)
	sol1 := 0
	sol2 := 0

	for scanner.Scan() {
		currentLine := scanner.Text()

		switch len(currentLine) {
		case 0:
			left, above := evalPattern(pattern, false)
			sol1 += left
			sol1 += above

			l2, a2 := evalPattern(pattern, true)
			sol2 += l2
			sol2 += a2
			// reset pattern
			pattern = make([]string, 0)
		default:
			pattern = append(pattern, currentLine)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// also eval last line
	left, above := evalPattern(pattern, false)
	sol1 += left
	sol1 += above

	l2, a2 := evalPattern(pattern, true)
	sol2 += l2
	sol2 += a2

	return sol1, sol2

}

func evalPattern(pattern []string, smudge bool) (int, int) {

	// if we are considering the pattern with the smudge
	// we can flip one position
	var totalFlipsAllowed int
	if smudge {
		totalFlipsAllowed = 1
	} else {
		totalFlipsAllowed = 0
	}
	// build column slice and check it
	columns := utils.TransposeStringSlice(pattern)
	leftVal := getValue(columns, totalFlipsAllowed, 1)
	// check rows
	aboveVal := getValue(pattern, totalFlipsAllowed, 100)

	return leftVal, aboveVal

}

func getValue(sli []string, totalFlipsAllowed int, multiplier int) int {
	output := 0
	rows := len(sli)
	var flips int
	for i := 1; i < rows; i++ {
		flips = totalFlipsAllowed
		distance := stringDistance(sli[i-1], sli[i])
		if distance <= flips {
			found := true
			flips -= distance

			for j := 1; i+j < rows && i-j-1 >= 0; j++ {
				distance = stringDistance(sli[i+j], sli[i-j-1])
				if distance > flips {
					found = false
					break
				}
				flips -= distance
			}
			if found && flips == 0 {
				output = i * multiplier
				break
			}

		}
	}

	return output
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
