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
	grid := make([][]rune, 0)

	for scanner.Scan() {
		currentLine := scanner.Text()
		row := utils.MapToRuneSlice(currentLine)
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tiltG := tiltWithDirection(grid, 'N')
	sol1 := totalLoad(tiltG) // 112773

	cycle := []rune{'N', 'W', 'S', 'E'}

	gridCopy := make([][]rune, len(grid))
	copy(gridCopy, grid)
	i := 0

	for i < 98 {
		for _, d := range cycle {
			grid = tiltWithDirection(grid, d)
		}
		i++
	}

	sol2 := totalLoad(grid) // 98894

	fmt.Printf("%d %d", sol1, sol2)
	return sol1, sol2

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

func tiltWithDirection(grid [][]rune, direction rune) [][]rune {

	inputRows := len(grid)
	inputCols := len(grid[0])
	out := make([][]rune, inputRows)

	var fn func(int, int, [][]rune, [][]rune)

	switch direction {
	case 'N':
		fn = tiltNorth
	case 'W':
		fn = tiltWest
	case 'S':
		fn = tiltSouth
	case 'E':
		fn = tiltEast
	}

	fn(inputRows, inputCols, grid, out)

	return out
}

func tiltNorth(inputRows int, inputCols int, grid [][]rune, out [][]rune) {

	out[0] = make([]rune, inputCols)
	copy(out[0], grid[0])

	for i := inputRows - 1; i > 0; i-- {
		for j := 0; j < inputCols; j++ {
			if grid[i][j] == '#' {
				if out[i] == nil {
					out[i] = make([]rune, inputCols)
				}
				out[i][j] = '#'
				continue
			}

			if grid[i][j] == 'O' {
				pos := 0
				for k := i - 1; k >= 0; k-- {
					if grid[k][j] == 'O' {
						continue
					} else if grid[k][j] == '#' {
						break
					} else {
						pos++
					}
				}
				if out[i-pos] == nil {
					out[i-pos] = make([]rune, inputCols)
				}
				out[i-pos][j] = 'O'
			}
		}
	}

}

func tiltWest(inputRows int, inputCols int, grid [][]rune, out [][]rune) {
	for i := 0; i < inputRows; i++ {
		if out[i] == nil {
			out[i] = make([]rune, inputCols)
		}
		if len(grid[i]) > 0 {
			out[i][0] = grid[i][0]
		}
	}

	for j := inputCols - 1; j > 0; j-- {
		for i := 0; i < inputRows; i++ {
			if out[i] == nil {
				out[i] = make([]rune, inputCols)
			}

			if len(grid[i]) == 0 {
				out[i] = make([]rune, inputCols)
				continue
			}

			if grid[i][j] == '#' {
				out[i][j] = '#'
				continue
			}

			if grid[i][j] == 'O' {
				pos := 0
				for k := j - 1; k >= 0; k-- {
					if grid[i][k] == 'O' {
						continue
					} else if grid[i][k] == '#' {
						break
					} else {
						pos++
					}
				}
				out[i][j-pos] = 'O'
			}
		}
	}

}

func tiltSouth(inputRows int, inputCols int, grid [][]rune, out [][]rune) {

	out[inputRows-1] = make([]rune, inputCols)
	copy(out[inputRows-1], grid[inputRows-1])

	for i := 0; i < inputRows; i++ {
		for j := 0; j < inputCols; j++ {
			if grid[i][j] == '#' {
				if out[i] == nil {
					out[i] = make([]rune, inputCols)
				}
				out[i][j] = '#'
				continue
			}

			if grid[i][j] == 'O' {
				pos := 0
				for k := i + 1; k < inputRows; k++ {
					if grid[k][j] == 'O' {
						continue
					} else if grid[k][j] == '#' {
						break
					} else {
						pos++
					}
				}
				if out[i+pos] == nil {
					out[i+pos] = make([]rune, inputCols)
				}
				out[i+pos][j] = 'O'
			}
		}
	}

}

func tiltEast(inputRows int, inputCols int, grid [][]rune, out [][]rune) {

	for j := 0; j < inputCols; j++ {
		for i := 0; i < inputRows; i++ {
			if out[i] == nil {
				out[i] = make([]rune, inputCols)
			}

			if grid[i][j] == '#' {
				out[i][j] = '#'
				continue
			}

			if grid[i][j] == 'O' {
				pos := 0
				for k := j + 1; k < inputCols; k++ {
					if grid[i][k] == 'O' {
						continue
					} else if grid[i][k] == '#' {
						break
					} else {
						pos++
					}
				}
				out[i][j+pos] = 'O'
			}
		}
	}
}
