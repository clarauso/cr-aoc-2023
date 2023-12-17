package utils

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

var SpaceRemRegex = regexp.MustCompile(` {2,}`)

func MapToArray(stringArray string) []int {

	arr := make([]int, 0)
	for _, v := range strings.Split(stringArray, " ") {
		num, _ := strconv.Atoi(v)
		arr = append(arr, num)
	}

	return arr
}

func MapToRuneSlice(input string) []rune {
	return []rune(input)
}

// greatest common divisor (greatestCommonDivisor) via Euclidean algorithm (taken from https://go.dev/play/p/SmzvkDjYlb)
func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (leastCommonMultiple) via GCD (taken from https://go.dev/play/p/SmzvkDjYlb)
func LeastCommonMultiple(a, b int, integers ...int) int {
	result := a * b / GreatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = LeastCommonMultiple(result, integers[i])
	}

	return result
}

// Gets an element from the input map also deleting it from the map.
func Pop[K string, V any](m map[K]V) V {

	var out V
	for key, value := range m {
		delete(m, key)
		out = value
		break
	}

	return out
}

func (p1 Point) ManhattanDistance(p2 Point) int {

	xDelta := math.Abs(float64(p1.X - p2.X))
	yDelta := math.Abs(float64(p1.Y - p2.Y))

	return int(xDelta + yDelta)

}
