package main

import (
	"log"
	"testing"

	"github.com/clarauso/cr-aoc-2023/day07"
	"github.com/clarauso/cr-aoc-2023/day08"
	"github.com/clarauso/cr-aoc-2023/day09"
	"github.com/clarauso/cr-aoc-2023/day11"
	"github.com/clarauso/cr-aoc-2023/day13"
	"github.com/clarauso/cr-aoc-2023/day14"
	"github.com/clarauso/cr-aoc-2023/day15"
)

func runTest(t *testing.T, toTest func(string) (int, int), inputFile string, expected1 int, expected2 int) {

	log.Printf("[ Start %s\n", t.Name())
	defer log.Printf("] End %s\n", t.Name())

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
	filename := "day07/input.txt"
	runTest(t, day07.CamelCards, filename, 248_836_197, 251_195_607)
}

func TestDay08Wasteland(t *testing.T) {
	filename := "day08/input.txt"
	runTest(t, day08.Wasteland, filename, 19_099, 17_099_847_107_071)
}

func TestDay09Oasis(t *testing.T) {
	filename := "day09/input.txt"
	runTest(t, day09.OasisReport, filename, 1_641_934_234, 975)
}

func TestDay11CosmicExpansion(t *testing.T) {
	filename := "day11/input.txt"
	runTest(t, day11.CosmicExpansion, filename, 9_312_968, 597_714_117_556)
}

func TestDay13PointOfIncidence(t *testing.T) {
	filename := "day13/input.txt"
	runTest(t, day13.PointOfIncidence, filename, 27_502, 31_947)
}

func TestDay14ReflectorDish(t *testing.T) {
	filename := "day14/input.txt"
	runTest(t, day14.ReflectorDish, filename, 112_773, 98_894)
}

func TestDay15LensLibrary(t *testing.T) {
	filename := "day15/input.txt"
	runTest(t, day15.LensLibrary, filename, 513_172, 237_806)
}
