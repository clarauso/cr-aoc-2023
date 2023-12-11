package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Universe [][]rune

func cosmicExpansion(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	universe := make([][]rune, 0)

	for scanner.Scan() {
		currentLine := scanner.Text()
		row := mapToRuneSlice(currentLine)
		universe = append(universe, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	colsToExpand, rowsToExpand := expand(universe)
	galaxiesSol1 := galaxiesLocator(universe, colsToExpand, rowsToExpand, 1)
	galaxiesSol2 := galaxiesLocator(universe, colsToExpand, rowsToExpand, 999_999)

	sol1 := sumDistances(galaxiesSol1)
	fmt.Printf("%d\n", sol1)

	sol2 := sumDistances(galaxiesSol2)
	fmt.Printf("%d\n", sol2)

	return sol1, sol2

}

func expand(u Universe) ([]int, []int) {

	cols := len(u[0])
	rows := len(u)

	columnsToExpand := make([]int, 0)
	for c := 0; c < cols; c++ {
		column := ""
		for j := 0; j < rows; j++ {
			column = column + string(u[j][c])
		}

		if !strings.ContainsRune(column, '#') {
			columnsToExpand = append(columnsToExpand, c)
		}

	}

	rowsToExpand := make([]int, 0)
	for i, row := range u {
		if !strings.ContainsRune(string(row), '#') {
			rowsToExpand = append(rowsToExpand, i)
		}
	}

	return columnsToExpand, rowsToExpand

}

func galaxiesLocator(u Universe, columnsToExpand []int, rowsToExpand []int, modifier int) []Point {

	out := make([]Point, 0)
	for i, _ := range u {
		for j, v := range u[i] {

			if v == '#' {

				plusCol := 0
				for _, col := range columnsToExpand {
					if j > col {
						plusCol += modifier
					}
				}

				plusRow := 0
				for _, row := range rowsToExpand {
					if i > row {
						plusRow += modifier
					}
				}

				out = append(out, Point{x: j + plusCol, y: i + plusRow})
			}

		}
	}

	return out

}

func sumDistances(points []Point) int {

	total := 0
	for i, p := range points {

		for j, q := range points {
			if i <= j {
				continue
			}
			d := p.manhattanDistance(q)
			total += d

		}

	}

	return total

}
