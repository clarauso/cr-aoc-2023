package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func boatRace() {

	file, err := os.Open("input-files/aoc_6_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linesRead := 0
	var timeSli, distanceSli []int
	var singleTime, singleDistance int
	for scanner.Scan() {
		currentLine := scanner.Text()

		currentLine = spaceRemRegex.ReplaceAllString(currentLine, " ")
		currentLine = strings.ReplaceAll(currentLine, "Time: ", "")
		currentLine = strings.ReplaceAll(currentLine, "Distance: ", "")

		fmt.Println(currentLine)

		switch {
		case linesRead == 0:
			timeSli = mapToArray(currentLine)
			singleTime = merge(currentLine)
		case linesRead == 1:
			distanceSli = mapToArray(currentLine)
			singleDistance = merge(currentLine)
		}

		linesRead++

	}

	fmt.Printf("%d %d\n", singleTime, singleDistance)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sol1 := 1
	for i, _ := range timeSli {

		record := distanceSli[i]
		wins := 0
		for j := 0; j < timeSli[i]; j++ {
			d := distance(timeSli[i], j)
			if d > record {
				wins++
			}
		}

		fmt.Printf("For duration %d you have %d ways to win\n", timeSli[i], wins)
		sol1 *= wins
	}

	sol2 := 0
	record := singleDistance
	for j := 0; j < singleTime; j++ {
		d := distance(singleTime, j)
		if d > record {
			sol2++
		}
	}
	fmt.Printf("Total is %d\n", sol2)

}

// Your toy boat has a starting speed of zero millimeters per millisecond.
// For each whole millisecond you spend at the beginning of the race holding down the button, the boat's speed increases by one millimeter per millisecond.

func speed(holdTime int) int {
	startingSpeed := 0
	deltaPerMs := 1
	return startingSpeed + holdTime*deltaPerMs
}

func distance(raceDuration int, holdTime int) int {

	remainingTime := raceDuration - holdTime

	return speed(holdTime) * remainingTime
}

func merge(line string) int {

	line = strings.ReplaceAll(line, " ", "")

	num, _ := strconv.Atoi(line)

	return num

}
