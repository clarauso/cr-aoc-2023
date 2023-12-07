package main

import (
	"log"
	"testing"
)

func TestDay07CamelCards(t *testing.T) {

	expected1 := 248_836_197
	expected2 := 251_195_607
	sol1, sol2 := camelCards("input-files/day07_input.txt")

	log.Printf("Solution 1 %d, solution 2 %d", sol1, sol2)

	if sol1 != expected1 {
		t.Fatalf("Solution 1 must be %d", expected1)
	}

	if sol2 != expected2 {
		t.Fatalf("Solution 2 must be %d", expected2)
	}

}
