package main

import (
	"regexp"
	"strconv"
	"strings"
)

var spaceRemRegex = regexp.MustCompile(` {2,}`)

func mapToArray(stringArray string) []int {

	arr := make([]int, 0)
	for _, v := range strings.Split(stringArray, " ") {
		num, _ := strconv.Atoi(v)
		arr = append(arr, num)
	}

	return arr
}

func mapToRuneSlice(input string) []rune {
	return []rune(input)
}

// greatest common divisor (greatestCommonDivisor) via Euclidean algorithm (taken from https://go.dev/play/p/SmzvkDjYlb)
func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (leastCommonMultiple) via GCD (taken from https://go.dev/play/p/SmzvkDjYlb)
func leastCommonMultiple(a, b int, integers ...int) int {
	result := a * b / greatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = leastCommonMultiple(result, integers[i])
	}

	return result
}

// Gets an element from the input map also deleting it from the map.
func pop[K string, V any](m map[K]V) V {

	var out V
	for key, value := range m {
		delete(m, key)
		out = value
		break
	}

	return out
}
