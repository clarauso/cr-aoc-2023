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
