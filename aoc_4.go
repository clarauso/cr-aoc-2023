package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func scratchCards() {

	file, err := os.Open("input-files/aoc_4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		idx, winning, myNumbers := parseCard(currentLine)

		value := evaluateCard(winning, myNumbers)
		total += value
		fmt.Printf("Card %d has value %d\n", idx, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// answer is
	fmt.Printf("Total is %d\n", total)

}

func parseCard(card string) (int, []int, []int) {
	// clean the input to avoid parsing exceptions
	spaceRemRegex := regexp.MustCompile(` {2,}`)
	card = spaceRemRegex.ReplaceAllString(card, " ")

	//[Card 1]:[41 48 83 86 17 | 83 86  6 31 17  9 48 53]
	first := strings.Split(card, ":")

	// index of the card
	idxString := strings.ReplaceAll(first[0], "Card ", "")
	idx, _ := strconv.Atoi(idxString)

	body := strings.TrimSpace(first[1])

	//[41 48 83 86 17][83 86  6 31 17  9 48 53]
	second := strings.Split(body, " | ")
	// 41 48 83 86 17
	winningString := second[0]
	// 83 86 6 31 17  9 48 53
	myNumbersString := second[1]

	winning := make([]int, 0)
	for _, v := range strings.Split(winningString, " ") {
		num, _ := strconv.Atoi(v)
		winning = append(winning, num)
	}

	myNumbers := make([]int, 0)
	for _, v := range strings.Split(myNumbersString, " ") {
		num, _ := strconv.Atoi(v)
		myNumbers = append(myNumbers, num)
	}

	return idx, winning, myNumbers

}

func evaluateCard(winningNumbers []int, myNumbers []int) int {

	// checks already done
	checks := make(map[int]bool)
	total := 0
	for _, n := range myNumbers {
		if slices.Contains(winningNumbers, n) {

			_, exists := checks[n]
			if !exists {
				checks[n] = true
				total += 1
			}

		}
	}

	fmt.Println(checks)

	if total > 0 {
		pow := math.Pow(2, float64(total-1))
		total = int(pow)
	}

	return total

}
