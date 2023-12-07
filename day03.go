package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func gearRatios() {

	file, err := os.Open("input-files/aoc_3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	engine := make([][]rune, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		currentLine := scanner.Text()
		// the line as an array of runes
		lineArr := parseLine(currentLine)
		engine = append(engine, lineArr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// first map all numbers with their start and end indexes
	partsMap := toPartsMap(engine)

	// then process all symbols
	total := sumGearRatios(engine, partsMap)

	// answer is 87287096
	fmt.Printf("Total is %d\n", total)

}

func toPartsMap(engine [][]rune) map[string]int {
	partsMap := make(map[string]int)
	for i, row := range engine {

		var strNumber string
		for j, item := range row {

			if unicode.IsDigit(item) {
				strNumber = strNumber + string(item)
				// reached end of the row
				if j == len(row)-1 {
					updateMap(partsMap, i, j, strNumber)
				}
				continue
			}

			// point or symbol
			if strNumber != "" {
				updateMap(partsMap, i, j, strNumber)
				strNumber = ""
			}

		}

	}

	return partsMap
}

func sumGearRatios(engine [][]rune, partsMap map[string]int) int {

	total := 0
	for i, row := range engine {

		for j, item := range row {
			//if !unicode.IsDigit(item) && item != '.' { // this condition for part 1
			if item == '*' {
				arr := keysToCheck(i, j)

				lastMatch := -1
				numMatches := 0
				partial := 1
				for _, key := range arr {
					partNumber, exists := partsMap[key]
					// second part of this check is a workaround to avoid double counting
					if exists && lastMatch != partNumber {
						numMatches++
						lastMatch = partNumber
						partial *= partNumber
					}
				}
				// a gear ratio is valid only if the symbol is adjacent to exactly two numbers
				if numMatches == 2 {
					total += partial
				}
			}

		}

	}

	return total

}

func toKey(row int, col int) string {
	a := fmt.Sprintf("[%d]", row)
	b := fmt.Sprintf("[%d]", col)
	return a + b
}

func parseLine(line string) []rune {
	return []rune(line)
}

func updateMap(partsMap map[string]int, row int, col int, value string) {

	startKey := toKey(row, col-len(value))
	endKey := toKey(row, col-1)
	number, _ := strconv.Atoi(value)

	partsMap[startKey] = number
	partsMap[endKey] = number
}

func keysToCheck(row int, col int) []string {

	out := make([]string, 0)

	// previous row
	for c := col - 1; c <= col+1; c++ {
		key := toKey(row-1, c)
		out = append(out, key)
	}

	// same row
	for c := col - 1; c <= col+1; c++ {
		key := toKey(row, c)
		out = append(out, key)
	}

	// subsequent row
	for c := col - 1; c <= col+1; c++ {
		key := toKey(row+1, c)
		out = append(out, key)
	}

	return out

}
