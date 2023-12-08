package main

import (
	"log"
	"testing"
)

const inputPath string = "input-files/"

func runTest(t *testing.T, toTest func(string) (int, int), inputFile string, expected1 int, expected2 int) {

	log.Println(t.Name())
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
	filename := inputPath + "day07_input.txt"
	runTest(t, camelCards, filename, 248_836_197, 251_195_607)
}

func TestDay08Wasteland(t *testing.T) {
	filename := inputPath + "day08_input.txt"
	runTest(t, wasteland, filename, 19_099, 17_099_847_107_071)
}
