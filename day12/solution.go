package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var springsLineRegex = regexp.MustCompile(`([.?#]+) (\d,\d,\d)`)

func hotSprings(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine := scanner.Text()
		parseSpringLine(currentLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sol1 := 0
	fmt.Printf("%d\n", sol1)

	sol2 := 0
	fmt.Printf("%d\n", sol2)

	return sol1, sol2

}

func parseSpringLine(line string) {

	parts := strings.Split(line, " ")

	springParts := regexp.MustCompile(`\.+`).Split(parts[0], math.MaxInt)
	fmt.Println(len(line))

	finalParts := make([]string, 0)
	for _, v := range springParts {
		if v != "" {
			finalParts = append(finalParts, v)
		}

	}

	ints := make([]int, 0)
	intParts := strings.Split(parts[1], ",")
	for _, i := range intParts {
		anInt, _ := strconv.Atoi(i)
		ints = append(ints, anInt)
	}

}
