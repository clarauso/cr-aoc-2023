package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func trebuchetAlt() {
	file, err := os.Open("input-files/aoc_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	// is 54770
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		total += handleLineExtendedAlt(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total is %d\n", total)
}

// one, two, three, four, five, six, seven, eight, and nine
func handleLineExtendedAlt(line string) int {

	stringDigits := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var firstDigit string
	var current string
	firstIdx := -1
	for idx, e := range line {

		if firstIdx != -1 {
			break
		}

		if unicode.IsDigit(e) {
			firstDigit = string(e)
			firstIdx = idx
			break
		}

		substring := false
		current = current + string(e)
		for digitIdx, digit := range stringDigits {

			if digit == current {
				firstDigit = string(fmt.Sprintf("%d", digitIdx))
				firstIdx = idx
				break
			} else if strings.Contains(digit, current) {
				substring = true
				break
			}

		}

		if !substring {
			current = string(e)
		}

	}
	// reset the string
	current = ""
	var secondDigit string
	secondIdx := -1

	for idx, e := range line[firstIdx+1:] {

		if unicode.IsDigit(e) {
			secondDigit = string(e)
			secondIdx = idx
			current = ""
			continue
		}

		substring := false
		current = current + string(e)
		for digitIdx, digit := range stringDigits {

			if digit == current {
				secondDigit = string(fmt.Sprintf("%d", digitIdx))
				secondIdx = idx
				current = ""
				break
			} else if strings.Contains(digit, current) {
				substring = true
				break
			}

		}

		if !substring && len(current) > 0 {
			current = current[1:]
		}

	}

	var numString string
	if secondIdx == -1 {
		numString = firstDigit + firstDigit
	} else {
		numString = firstDigit + secondDigit
	}

	value, _ := strconv.Atoi(numString)

	fmt.Printf("Line [%s] is [%d]\n", line, value)

	return value

}
