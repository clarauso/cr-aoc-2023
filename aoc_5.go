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
	numLine := regexp.MustCompile(`^\d{1,} \d{1,} \d{1,}`)
	mapNameLine := regexp.MustCompile(`^([a-z]{1,}-to-[a-z]{1,}) map:`)
	scanner := bufio.NewScanner(file)
	var mappingFunctions []func(int) int
	currentMap := ""
	for scanner.Scan() {
		currentLine := scanner.Text()

		if linesRead == 0 {
			//seedList = mapSeedLine(currentLine)
			seedList = mapSeedLineWithRange(currentLine)
		}

		if mapNameLine.MatchString(currentLine) {
			if currentMap != "" {
				allMappingFn[currentMap] = mappingFunctions
			}
			matches := mapNameLine.FindStringSubmatch(currentLine)
			// TODO do not use hardcoded index
			currentMap = matches[1]
			mappingFunctions = allMappingFn[currentMap]
		}

		if numLine.MatchString(currentLine) {
			a, b, c := mapNumericLine(currentLine)
			//fmt.Printf("%d %d %d\n", a, b, c)
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

	minValue := math.MaxInt
	for _, seed := range seedList {

		value := seed
		value = iterativeMap(value, allMappingFn["seed-to-soil"])
		value = iterativeMap(value, allMappingFn["soil-to-fertilizer"])
		value = iterativeMap(value, allMappingFn["fertilizer-to-water"])
		value = iterativeMap(value, allMappingFn["water-to-light"])
		value = iterativeMap(value, allMappingFn["light-to-temperature"])
		value = iterativeMap(value, allMappingFn["temperature-to-humidity"])
		value = iterativeMap(value, allMappingFn["humidity-to-location"])

		if value < minValue {
			minValue = value
		}
	}

	fmt.Printf("Min location is %d\n", minValue)

}

// for part 1
func mapSeedLine(line string) []int {
	// seeds: 79 14 55 13
	line = strings.ReplaceAll(line, "seeds: ", "")
	return mapToArray(line)
}

// for part 2
func mapSeedLineWithRange(line string) []int {
	// seeds: 79 14 55 13

	baseArray := mapSeedLine(line)
	outArray := make([]int, 0)
	n := len(baseArray)

	for i := 0; i < n; i += 2 {
		fmt.Printf("Index %d of %d will add %d\n", i, n, baseArray[i+1])
		for j := 0; j < baseArray[i+1]; j++ {
			item := baseArray[i] + j
			outArray = append(outArray, item)
		}
	}

	return outArray
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

func iterativeMap(toMap int, functions []func(int) int) int {
	mapped := toMap
	for _, myFn := range functions {
		mapped = myFn(toMap)
		if mapped != toMap {
			break
		}
	}

	return mapped
}
