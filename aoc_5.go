package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

func seeds() {

	file, err := os.Open("input-files/aoc_5_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	allMappingFn := map[string][]func(int) int{"seed-to-soil": make([]func(int) int, 0),
		"soil-to-fertilizer":      make([]func(int) int, 0),
		"fertilizer-to-water":     make([]func(int) int, 0),
		"water-to-light":          make([]func(int) int, 0),
		"light-to-temperature":    make([]func(int) int, 0),
		"temperature-to-humidity": make([]func(int) int, 0),
		"humidity-to-location":    make([]func(int) int, 0)}

	linesRead := 0
	var seedList []int
	numLine := regexp.MustCompile(`^\d{1,} \d{1,} \d{1,}$`)
	mapNameLine := regexp.MustCompile(`^([a-z]{1,}-to-[a-z]{1,}) map:$`)
	scanner := bufio.NewScanner(file)
	var mappingFunctions []func(int) int
	currentMap := ""
	for scanner.Scan() {
		currentLine := scanner.Text()

		if linesRead == 0 {
			seedList = mapSeedLine(currentLine)
		}

		if mapNameLine.MatchString(currentLine) {
			if currentMap != "" {
				allMappingFn[currentMap] = mappingFunctions
			}
			matches := mapNameLine.FindStringSubmatch(currentLine)
			currentMap = matches[1]
			mappingFunctions = allMappingFn[currentMap]
		}

		if numLine.MatchString(currentLine) {
			a, b, c := mapNumericLine(currentLine)
			mappingFunctions = append(mappingFunctions, aToB(a, b, c))
		}

		linesRead++
	}

	// to set last item
	if currentMap != "" {
		allMappingFn[currentMap] = mappingFunctions
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// 650599855 for part 1
	sol1 := find(seedList, allMappingFn)
	fmt.Printf("Min location (part 1) is %d\n", sol1)

	sol2 := findAll(seedList, allMappingFn)
	// 1240035 for part 2
	fmt.Printf("Min location (part 2) is %d\n", sol2)

}

func mapSeedLine(line string) []int {
	// seeds: 79 14 55 13
	line = strings.ReplaceAll(line, "seeds: ", "")
	return mapToArray(line)
}

func mapNumericLine(line string) (int, int, int) {
	// 50 98 2
	intArr := mapToArray(line)

	return intArr[0], intArr[1], intArr[2]
}

func aToB(destRangeStart int, sourceRangeStart int, rangeLength int) func(int) int {
	fn := func(a int) int {

		if a >= sourceRangeStart && a < sourceRangeStart+rangeLength {
			delta := a - sourceRangeStart
			return destRangeStart + delta
		}

		return a

	}

	return fn
}

func applyAll(toMap int, functions []func(int) int) int {
	mapped := toMap
	for _, myFn := range functions {
		mapped = myFn(toMap)
		if mapped != toMap {
			break
		}
	}

	return mapped
}

func find(seedList []int, allMappingFn map[string][]func(int) int) int {
	minValue := math.MaxInt
	mappings := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

	for _, seed := range seedList {

		value := seed
		for _, mappingKey := range mappings {
			value = applyAll(value, allMappingFn[mappingKey])
		}

		if value < minValue {
			minValue = value
		}
	}

	return minValue

}

func findAll(seedList []int, allMappingFn map[string][]func(int) int) int {
	minValue := math.MaxInt
	mappings := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

	n := len(seedList)

	for i := 0; i < n; i += 2 {
		fmt.Printf("Processing index %d out of %d\n", i, n)
		for j := 0; j < seedList[i+1]; j++ {
			value := seedList[i] + j
			for _, mappingKey := range mappings {
				value = applyAll(value, allMappingFn[mappingKey])
			}

			if value < minValue {
				minValue = value
			}
		}
	}

	return minValue

}
