package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/clarauso/cr-aoc-2023/utils"
)

func ReflectorDish(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sol2 := 0
	grid := make([][]rune, 0)

	for scanner.Scan() {
		currentLine := scanner.Text()
		row := utils.MapToRuneSlice(currentLine)
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tilt := tilt(grid)
	sol1 := totalLoad(tilt) // 112773

	fmt.Printf("%d %d", sol1, sol2)

	return sol1, sol2

}

func tilt(grid [][]rune) [][]rune {

	inputRows := len(grid)
	inputCols := len(grid[0])
	out := make([][]rune, inputRows)

	out[0] = grid[0]

	for i := inputRows - 1; i > 0; i-- {
		for j := 0; j < inputCols; j++ {
			if grid[i][j] == 'O' {
				pos := 0
				for k := i - 1; k >= 0; k-- {
					if grid[k][j] == '.' {
						pos++
					} else if grid[k][j] == 'O' {
						continue
					} else {
						break
					}
				}
				if out[i-pos] == nil {
					out[i-pos] = make([]rune, inputCols)
				}
				out[i-pos][j] = 'O'
			}

		}
	}

	return out
}

func totalLoad(grid [][]rune) int {

	totalLoad := 0
	inputRows := len(grid)

	for i, r := range grid {
		rocks := strings.Count(string(r), "O")
		totalLoad += rocks * (inputRows - i)
	}

	return totalLoad
}
