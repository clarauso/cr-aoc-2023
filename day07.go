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
	cards          []rune
	bid            int
	handValue      int
	handValueJolly int
}

var cardValuesFirst = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}

var cardValuesSecond = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2, 'J': 1}

func camelCards() {

	file, err := os.Open("input-files/day07_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := make([]Hand, 0)

	for scanner.Scan() {
		currentLine := scanner.Text()
		currentLine = spaceRemRegex.ReplaceAllString(currentLine, " ")
		h := parseHand(currentLine)
		hands = append(hands, h)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// sol 1
	sortHands(hands, sortingFieldGetter(false), cardValuesFirst)
	sol1 := 0
	for i, h := range hands {
		sol1 += (i + 1) * h.bid
	}
	// must be 248836197
	fmt.Printf("Solution 1 is %d\n", sol1)

	// sol 2
	sortHands(hands, sortingFieldGetter(true), cardValuesSecond)
	sol2 := 0
	for i, h := range hands {
		sol2 += (i + 1) * h.bid
	}
	// must be 251195607
	fmt.Printf("Solution 2 is %d\n", sol2)

}

func parseHand(line string) Hand {

	sli := strings.Split(line, " ")
	bid, _ := strconv.Atoi(sli[1])
	runes := make([]rune, len(sli[0]))

	for i, v := range sli[0] {
		runes[i] = v
	}

	val, valJolly := computeHandValue(runes)

	return Hand{cards: runes, bid: bid, handValue: val, handValueJolly: valJolly}

}

func computeHandValue(cards []rune) (int, int) {

	labelCount := make(map[rune]int, 5)

	for _, l := range cards {

		_, x := labelCount[l]
		if x {
			labelCount[l] += 1
		} else {
			labelCount[l] = 1
		}

	}

	// values for hand types
	// - Five of a kind, where all five cards have the same label
	// - Four of a kind, where four cards have the same label and one card has a different label
	// - Full house, where three cards have the same label, and the remaining two cards share a different label
	// - Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand
	// - Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label
	// - One pair, where two cards share one label, and the other three cards have a different label from the pair and each other
	const fiveOf = 6
	const fourOf = 5
	const fullHouse = 4
	const threeOf = 3
	const aPair = 1

	pairs := 0
	plainValue := 0
	numJolly := 0
	for k, v := range labelCount {

		if k == 'J' {
			numJolly = v
		}

		switch v {

		case 5:
			// 5 of a kind
			plainValue = fiveOf
		case 4:
			// 4 of a kind
			plainValue = fourOf
		case 3:
			// a tris, the hand can be full house OR 3 of a kind
			plainValue = threeOf
		case 2:
			// a pair, the hand can be one pair OR two pair
			pairs++
		}

	}

	// pairs can be 1 (plainValue == 3, becomes full house) or 2 (plainValue == 0, becomes a pair)
	val := plainValue + pairs
	// no J or all J: cannot add points
	if numJolly == 0 || numJolly == 5 {
		return val, val
	}

	// here we have at least 1 J
	var valWithJolly int
	switch val {
	case 5, 4:
		// 4 of a kind XXXX Y OR full house (1 tris + 1 pair) XXX YY
		// always becomes 5 of a kind
		valWithJolly = fiveOf
	case 3:
		// 3 of a kind XXX ZY
		if numJolly == 2 {
			// 2 J, becomes full house
			valWithJolly = fullHouse
		} else {
			// 1 or 3 J, becomes 4 of a kind
			valWithJolly = fourOf
		}
	case 2:
		// 2 pairs XX YY Z
		if numJolly == 2 {
			// 2 J, becomes 4 of a kind
			valWithJolly = fourOf
		} else {
			// 1 J, becomes a full house
			valWithJolly = fullHouse
		}
	case 1:
		// a pair XX YZW
		// pair becomes 3 of a kind
		valWithJolly = threeOf
	case 0:
		// XZYWQ
		// high card becomes a pair
		valWithJolly = aPair
	}

	return val, valWithJolly

}

// sort by hand value and card values (stopping at first different card)
func sortHands(sli []Hand, handValExtractor func(Hand) int, cardValues map[rune]int) {
	sort.Slice(sli, func(i, j int) bool {

		getVal := cardValueFn(cardValues)

		iVal := handValExtractor(sli[i])
		jVal := handValExtractor(sli[j])

		if iVal == jVal {

			var r bool
			for k, _ := range sli[i].cards {
				// same card, keep comparing
				if sli[i].cards[k] == sli[j].cards[k] {
					continue
				}
				r = getVal(sli[i].cards[k]) < getVal(sli[j].cards[k])
				break
			}

			return r

		} else {
			return iVal < jVal
		}

	})
}

func cardValueFn(cardValues map[rune]int) func(rune) int {

	return func(card rune) int {
		return cardValues[card]
	}

}

// Returns a function that reads the field to be considered for sorting
func sortingFieldGetter(considerJolly bool) func(Hand) int {

	if considerJolly {
		return func(h Hand) int { return h.handValueJolly }
	} else {
		return func(h Hand) int { return h.handValue }
	}

}
