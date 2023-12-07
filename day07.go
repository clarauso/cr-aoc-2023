package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards     []rune
	bid       int
	handValue int
}

var cardValues = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}

func camelCards() {

	file, err := os.Open("input-files/day07_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linesRead := 0
	hands := make([]Hand, 0)

	for scanner.Scan() {
		currentLine := scanner.Text()

		currentLine = spaceRemRegex.ReplaceAllString(currentLine, " ")

		fmt.Println(currentLine)

		h := parseHand(currentLine)

		hands = append(hands, h)

		linesRead++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// sol 1
	sortHands(hands)
	sol1 := 0
	for i, h := range hands {
		fmt.Printf("%d %s\n", i, string(h.cards))
		sol1 += (i + 1) * h.bid
	}
	// must be 248836197
	fmt.Printf("Solution 1 is %d\n", sol1)

}

func parseHand(line string) Hand {

	sli := strings.Split(line, " ")
	bid, _ := strconv.Atoi(sli[1])
	runes := make([]rune, len(sli[0]))

	for i, v := range sli[0] {
		runes[i] = v
	}

	val := computeHandValue(runes)

	return Hand{cards: runes, bid: bid, handValue: val}

}

func computeHandValue(cards []rune) int {

	labelCount := make(map[rune]int, 5)

	for _, l := range cards {

		_, x := labelCount[l]
		if x {
			labelCount[l] += 1
		} else {
			labelCount[l] = 1
		}

	}

	tris := false
	pairs := 0
	for _, v := range labelCount {

		switch v {
		case 5:
			return 6
		case 4:
			return 5
		case 3:
			tris = true
		case 2:
			pairs++
		}

	}

	if tris && pairs == 1 {
		return 4
	} else if tris {
		return 3
	}

	return pairs

}

// sort by value and first card
func sortHands(sli []Hand) {
	sort.Slice(sli, func(i, j int) bool {

		if sli[i].handValue == sli[j].handValue {

			var r bool
			for k, _ := range sli[i].cards {
				if sli[i].cards[k] == sli[j].cards[k] {
					continue
				}
				r = cardValue(sli[i].cards[k]) < cardValue(sli[j].cards[k])
				break
			}

			return r

		} else {
			return sli[i].handValue < sli[j].handValue
		}

	})
}

func cardValue(card rune) int {

	return cardValues[card]

}
