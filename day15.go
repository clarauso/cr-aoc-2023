package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func lensLibrary(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sol1 := 0
	sol2 := 0
	var steps []string

	for scanner.Scan() {
		currentLine := scanner.Text()

		steps = strings.Split(currentLine, ",")
		fmt.Println(currentLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, s := range steps {
		sol1 += stepHash(s)
	}

	// 513172
	fmt.Printf("%d\n", sol1)

	fmt.Printf("%d\n", sol2)

	return sol1, sol2

}

func stepHash(input string) int {

	out := 0

	for _, c := range input {
		out += int(c)
		out *= 17
		out %= 256
	}

	fmt.Printf("%s becomes %d\n", input, out)

	return out

}
