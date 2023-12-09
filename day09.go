package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func oasisReport(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sol1 := 0
	sol2 := 0

	for scanner.Scan() {
		currentLine := scanner.Text()
		sli := mapToArray(currentLine)

		compSeq := extrapolatedDiff(sli)
		a, b := extrapolate(sli, compSeq)

		sol1 += a
		sol2 += b
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// 1641934234
	fmt.Printf("%d\n", sol1)
	// 975
	fmt.Printf("%d\n", sol2)

	return sol1, sol2

}

func extrapolatedDiff(seq []int) []int {

	hasNotZeroVals := true
	sequences := make([][]int, 0)

	for hasNotZeroVals {
		hasNotZeroVals = false
		n := len(seq)
		newSeq := make([]int, n-1)
		for i := 1; i < n; i++ {
			newSeq[i-1] = seq[i] - seq[i-1]
			if newSeq[i-1] != 0 {
				hasNotZeroVals = true
			}
		}
		seq = newSeq
		// for part 1
		newSeq = append(newSeq, 0)
		// for part 2
		newSeq = append([]int{0}, newSeq...)
		sequences = append(sequences, newSeq)
	}

	n := len(sequences)

	for i := n - 2; i >= 0; i-- {
		m := len(sequences[i])
		sequences[i][m-1] = sequences[i][m-2] + sequences[i+1][m-2]
		sequences[i][0] = sequences[i][1] + (-sequences[i+1][0])
	}

	return sequences[0]

}

func extrapolate(a []int, b []int) (int, int) {

	n := len(a)
	m := len(b)

	newVal := a[n-1] + b[m-1]
	oldVal := a[0] - b[0]

	return newVal, oldVal
}
