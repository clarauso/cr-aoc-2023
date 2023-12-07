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

// values for hand types
const fiveOf int = 6    // five of a kind, where all five cards have the same label
const fourOf int = 5    // four of a kind, where four cards have the same label and one card has a different label
const fullHouse int = 4 // full house, where three cards have the same label, and the remaining two cards share a different label
const threeOf int = 3   // three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand
const twoPairs int = 2  // two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label
const aPair int = 1     // one pair, where two cards share one label, and the other three cards have a different label from the pair and each other
const highCard int = 0  // high card, all cards have different labels

func camelCards(inputFilePath string) {

	file, err := os.Open(inputFilePath)
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
	sortHands(hands, valueGetterFn(false), cardValuesFirst)
	sol1 := 0
	for i, h := range hands {
		sol1 += (i + 1) * h.bid
	}
	// must be 248836197
	fmt.Printf("Solution 1 is %d\n", sol1)

	// sol 2
	sortHands(hands, valueGetterFn(true), cardValuesSecond)
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

	val, valJolly := computeCardsValue(runes)

	return Hand{cards: runes, bid: bid, handValue: val, handValueJolly: valJolly}

}

func computeCardsValue(cards []rune) (int, int) {

	// count how many occurrences there are for each card
	cardCount := make(map[rune]int, 5)
	for _, l := range cards {

		_, x := cardCount[l]
		if x {
			cardCount[l] += 1
		} else {
			cardCount[l] = 1
		}

	}

	baseType := computeType(cardCount)
	numJolly := cardCount['J']
	typeWithJolly := computeTypeWithJolly(baseType, numJolly)

	return baseType, typeWithJolly

}

// Computes hand's type with the given number of jolly cards
func computeType(cardCount map[rune]int) int {
	baseType := highCard
	pairs := 0

	for _, v := range cardCount {

		switch v {

		case 5:
			// 5 of a kind
			baseType = fiveOf
		case 4:
			// 4 of a kind
			baseType = fourOf
		case 3:
			// a tris, the hand can be full house OR 3 of a kind
			baseType = threeOf
		case 2:
			// a pair, the hand can be one pair OR two pair
			pairs++
		}

	}

	// pairs can be 0 (do not modify type), 1 (plainValue == 3, becomes full house) or 2 (plainValue == 0, becomes a pair)
	return baseType + pairs
}

// Computes hand's type with the given number of jolly cards
func computeTypeWithJolly(baseType int, numJolly int) int {
	// no J or all J: cannot add points
	if numJolly == 0 || numJolly == 5 {
		return baseType
	}
	// 1 to 4 J here
	var typeWithJolly int
	switch baseType {
	case fourOf, fullHouse:
		// 4 of a kind XXXX Y OR full house XXX YY: both become 5 of a kind
		typeWithJolly = fiveOf
	case threeOf:
		// 3 of a kind XXX ZY, we can have 1 or 3 jolly: both become 4 of a kind
		typeWithJolly = fourOf
	case twoPairs:
		// 2 pairs XX YY Z
		if numJolly == 2 {
			// 2 J, becomes 4 of a kind
			typeWithJolly = fourOf
		} else {
			// 1 J, becomes a full house
			typeWithJolly = fullHouse
		}
	case aPair:
		// a pair XX YZW, becomes 3 of a kind
		typeWithJolly = threeOf
	case highCard:
		// high card becomes a pair
		typeWithJolly = aPair
	}

	return typeWithJolly
}

// Sorts by hand value and card values
func sortHands(sli []Hand, handValExtractor func(Hand) int, cardValues map[rune]int) {
	getVal := cardValueFn(cardValues)

	sort.Slice(sli, func(i, j int) bool {
		iVal := handValExtractor(sli[i])
		jVal := handValExtractor(sli[j])

		if iVal != jVal {

			return iVal < jVal

		} else {
			// if the values are equal, sort using card values
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
		}

	})
}

// Returns a function that provides the value of a given card using the input map
func cardValueFn(cardValues map[rune]int) func(rune) int {

	return func(card rune) int {
		return cardValues[card]
	}

}

// Returns a function that reads the value field of an hand
func valueGetterFn(withJolly bool) func(Hand) int {

	if withJolly {
		return func(h Hand) int { return h.handValueJolly }
	} else {
		return func(h Hand) int { return h.handValue }
	}

}
