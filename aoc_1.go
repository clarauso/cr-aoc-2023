package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func trebuchet() {
	file, err := os.Open("input-files/aoc_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 54770 is the answer
	total := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		total += handleLineExtended(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total is %d\n", total)
}

func handleLine(line string) int {

	runes := make([]rune, 0)

	for _, e := range line {

		if !unicode.IsDigit(e) {
			continue
		}

		runes = append(runes, e)
	}

	n := len(runes)
	numString := string(runes[0]) + string(runes[n-1])

	value, _ := strconv.Atoi(numString)

	return value
}

// one, two, three, four, five, six, seven, eight, and nine
func handleLineExtended(line string) int {

	digits := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	stringDigits := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var firstDigit string
	minDigit := math.MaxInt
	minStrDigit := math.MaxInt
	for idx, e := range digits {
		digitIdx := strings.Index(line, e)
		strDigitIdx := strings.Index(line, stringDigits[idx])

		if digitIdx != -1 && digitIdx < minDigit && digitIdx < minStrDigit {
			minDigit = digitIdx
			firstDigit = e
		}

		if strDigitIdx != -1 && strDigitIdx < minStrDigit && strDigitIdx < minDigit {
			minStrDigit = strDigitIdx
			firstDigit = e
		}
	}

	var secondDigit string
	maxDigit := -1
	maxStrDigit := -1
	for idx, e := range digits {
		digitIdx := strings.LastIndex(line, e)
		strDigitIdx := strings.LastIndex(line, stringDigits[idx])

		if digitIdx != -1 && digitIdx > maxDigit && digitIdx > maxStrDigit {
			maxDigit = digitIdx
			secondDigit = e
		}

		if strDigitIdx != -1 && strDigitIdx > maxStrDigit && strDigitIdx > maxDigit {
			maxStrDigit = strDigitIdx
			secondDigit = e
		}
	}

	numString := firstDigit + secondDigit
	value, _ := strconv.Atoi(numString)

	fmt.Printf("Line [%s] is [%d]\n", line, value)

	return value

}
