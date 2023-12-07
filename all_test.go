package main

import (
	"log"
	"testing"
)

func runTest(t *testing.T, toTest func(string) (int, int), inputFile string, expected1 int, expected2 int) {

	sol1, sol2 := toTest(inputFile)

	log.Printf("Solution 1 %d, solution 2 %d", sol1, sol2)

	if sol1 != expected1 {
		t.Fatalf("Solution 1 must be %d", expected1)
	}

	if sol2 != expected2 {
		t.Fatalf("Solution 2 must be %d", expected2)
	}
}

func TestDay07CamelCards(t *testing.T) {
	runTest(t, camelCards, "input-files/day07_input.txt", 248_836_197, 251_195_607)
}
