package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// As you walk, the Elf shows you a small bag and some cubes which are either red, green, or blue.
// Each time you play this game, he will hide a secret number of cubes of each color in the bag, and your goal
// is to figure out information about the number of cubes.

// To get information, once a bag has been loaded with cubes, the Elf will reach into the bag, grab a handful of random cubes,
//  show them to you, and then put them back in the bag. He'll do this a few times per game.
//
// You play several games and record the information from each game (your puzzle input).
// Each game is listed with its ID number (like the 11 in Game 11: ...) followed by a semicolon-separated list of
// subsets of cubes that were revealed from the bag (like 3 red, 5 green, 4 blue).

// 3 type of cubes: red, green, blue

var allCubes = map[string]int{"red": 12, "green": 13, "blue": 14}

func cubeGame() {

	file, err := os.Open("input-files/aoc_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	totalPower := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		currentLine := scanner.Text()
		idx, possible := evalGame(currentLine)
		//fmt.Printf("Game %d possible %t\n", idx, possible)
		if possible {
			total += idx
		}

		idx, power := gamePower(currentLine)
		fmt.Printf("Game %d power %d\n", idx, power)
		totalPower += power
	}

	fmt.Printf("Result is %d, total power %d\n", total, totalPower)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func parseGame(line string) (int, []string) {
	components := strings.Split(line, ":")

	game := strings.TrimSpace(components[0])
	game = strings.ReplaceAll(game, "Game ", "")
	gameIdx, _ := strconv.Atoi(game)

	rounds := strings.TrimSpace(components[1])
	return gameIdx, strings.Split(rounds, ";")
}

// for question 1
func evalGame(line string) (int, bool) {

	gameIdx, rounds := parseGame(line)

	possible := true
	for _, val := range rounds {
		//fmt.Println(strings.TrimSpace(val))
		cubes := strings.Split(val, ",")

		cubeMap := make(map[string]int, 0)
		for _, cVal := range cubes {
			cVal = strings.TrimSpace(cVal)
			set := strings.Split(cVal, " ")
			num, _ := strconv.Atoi(set[0])
			cubeMap[set[1]] = num
		}

		for key, val := range cubeMap {

			total := allCubes[key]
			if total < val {
				//fmt.Printf("Game %d failed because of %s=%d\n", gameIdx, key, val)
				possible = false
				break
			}

		}

		if !possible {
			break
		}
	}

	return gameIdx, possible

}

// for question 2
func gamePower(line string) (int, int) {

	gameIdx, rounds := parseGame(line)

	requiredCubesMap := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, val := range rounds {
		//fmt.Println(strings.TrimSpace(val))
		cubes := strings.Split(val, ",")

		cubeMap := make(map[string]int, 0)
		// cubes by color
		for _, cVal := range cubes {
			cVal = strings.TrimSpace(cVal)
			set := strings.Split(cVal, " ")
			num, _ := strconv.Atoi(set[0])
			// color: number
			cubeMap[set[1]] = num
		}

		for key, val := range cubeMap {
			// current required cubes for the give color
			reqForColor := requiredCubesMap[key]

			if val > reqForColor {
				requiredCubesMap[key] = val
			}

		}

	}

	power := 1
	for _, val := range requiredCubesMap {
		power *= val
	}

	return gameIdx, power

}
