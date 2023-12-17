package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/clarauso/cr-aoc-2023/utils"
)

type Universe [][]rune

func CosmicExpansion(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	universe := make([][]rune, 0)

	for scanner.Scan() {
		currentLine := scanner.Text()
		row := utils.MapToRuneSlice(currentLine)
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

func galaxiesLocator(u Universe, columnsToExpand []int, rowsToExpand []int, modifier int) []utils.Point {

	out := make([]utils.Point, 0)
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

				out = append(out, utils.Point{X: j + plusCol, Y: i + plusRow})
			}

		}
	}

	return out

}

func sumDistances(points []utils.Point) int {

	total := 0
	for i, p := range points {

		for j, q := range points {
			if i <= j {
				continue
			}
			d := p.ManhattanDistance(q)
			total += d

		}

	}

	return total

}
