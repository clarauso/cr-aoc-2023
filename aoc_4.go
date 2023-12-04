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

type Card struct {
	idx            int
	winningNumbers []int
	playerNumbers  []int
}

var spaceRemRegex = regexp.MustCompile(` {2,}`)

func scratchCards() {

	file, err := os.Open("input-files/aoc_4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalValue := 0
	// how many additional cards each card id gives
	allCopies := make([]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine := scanner.Text()

		card := parseCard(currentLine)
		// question 1
		totalValue += card.evaluateCard()
		// question 2
		copies := card.countCopiesWon()
		allCopies = append(allCopies, copies)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// answer is 17782
	fmt.Printf("Total is %d\n", totalValue)

	total := sumAllCopies(allCopies)
	// answer is 8477787
	fmt.Printf("Total is %d\n", total)
}

func parseCard(card string) Card {
	// clean the input to avoid parsing exceptions
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

	winning := mapToArray(winningString)
	myNumbers := mapToArray(myNumbersString)

	return Card{idx: idx, winningNumbers: winning, playerNumbers: myNumbers}
}

// for question 1
func (c Card) evaluateCard() int {

	// checks already done
	checks := make(map[int]bool)
	total := 0
	for _, n := range c.playerNumbers {
		if slices.Contains(c.winningNumbers, n) {

			_, exists := checks[n]
			if !exists {
				checks[n] = true
				total += 1
			}

		}
	}

	if total > 0 {
		pow := math.Pow(2, float64(total-1))
		total = int(pow)
	}

	return total

}

// how many copies of cards you win for the given one
func (c Card) countCopiesWon() int {

	// checks already done
	checks := make(map[int]bool)
	total := 0
	for _, n := range c.playerNumbers {
		if slices.Contains(c.winningNumbers, n) {

			_, exists := checks[n]
			if !exists {
				checks[n] = true
				total += 1
			}

		}
	}

	return total

}

func mapToArray(stringArray string) []int {

	arr := make([]int, 0)
	for _, v := range strings.Split(stringArray, " ") {
		num, _ := strconv.Atoi(v)
		arr = append(arr, num)
	}

	return arr
}

func sumAllCopies(allCopies []int) int {

	totalCards := len(allCopies)
	// final card counters set to 1
	out := make([]int, totalCards)
	for i := 0; i < totalCards; i++ {
		out[i] = 1
	}

	for i, v := range allCopies {

		for j := i + 1; j <= i+v; j++ {
			out[j] += out[i]
		}

	}

	total := 0
	for _, num := range out {
		total += num
	}

	return total
}
