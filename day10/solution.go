package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type GraphNode[L string | rune] struct {
	start bool
	label L
	pos   Point
	in    []*GraphNode[L]
	out   []*GraphNode[L]
}

var rightMap = map[rune][]rune{
	'-': {'-', '7', 'J'},
	'L': {'-', '7', 'J'},
	'F': {'-', '7', 'J'},
	'S': {'-', '7', 'J'},
}

var downMap = map[rune][]rune{
	'|': {'|', 'L', 'J'},
	'7': {'|', 'L', 'J'},
	'F': {'|', 'L', 'J'},
	'S': {'|', 'L', 'J'},
}

var upMap = map[rune][]rune{
	'|': {'|', 'F', '7'},
	'L': {'|', 'F', '7'},
	'J': {'|', 'F', '7'},
	'S': {'|', 'L', 'J'},
}

var leftMap = map[rune][]rune{
	'-': {'-', 'F', 'L'},
	'J': {'-', 'F', 'L'},
	'7': {'-', 'F', 'L'},
	'S': {'-', 'F', 'L'},
}

// The pipes are arranged in a two-dimensional grid of tiles:
//
// '|' is a vertical pipe connecting north and south.
// '-' is a horizontal pipe connecting east and west.
// 'L' is a 90-degree bend connecting north and east.
// 'J' is a 90-degree bend connecting north and west.
// '7' is a 90-degree bend connecting south and west.
// 'F' is a 90-degree bend connecting south and east.
// '.' is ground; there is no pipe in this tile.
// 'S' is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
// Based on the acoustics of the animal's scurrying, you're confident the pipe that contains the animal is one large, continuous loop.

// Instead, you need to find the tile that would take the longest number of steps along the loop to reach from the
// starting point - regardless of which way around the loop the animal went.

func pipeMaze(inputFilePath string) (int, int) {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]rune, 0)
	sol1 := 0
	sol2 := 0

	for scanner.Scan() {
		currentLine := scanner.Text()
		row := mapToRuneSlice(currentLine)
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	animalPosition := findAnimal(grid)

	fmt.Printf("Animal position %o", animalPosition)

	startNode := GraphNode[rune]{label: grid[animalPosition.y][animalPosition.x], pos: animalPosition, start: true}

	explore(&startNode, grid, make(map[*GraphNode[rune]]bool))

	//
	fmt.Printf("%d\n", sol1)
	//
	fmt.Printf("%d\n", sol2)

	return sol1, sol2

}

func findAnimal(grid [][]rune) Point {
	var animalPosition Point
	for i, v := range grid {
		for _, r := range v {

			if idx := strings.Index(string(r), "S"); idx != -1 {
				animalPosition = Point{y: i, x: idx}
			}

			fmt.Printf("%s", string(r))
		}
		fmt.Printf("\n")
	}

	return animalPosition
}

func explore(graphNode *GraphNode[rune], grid [][]rune, visited map[*GraphNode[rune]]bool) {

	rows := len(grid)
	cols := len(grid[0])
	current := graphNode.pos

	if len(visited) > 0 && graphNode.start {
		return
	}

	i := current.y
	j := current.x

	if i < rows-1 {
		// down
		r := grid[i+1][j]
		downSli := downMap[graphNode.label]
		if slices.Contains(downSli, r) {
			fmt.Printf("Can connect down %c and %c\n", graphNode.label, r)
			graphNode.out = appendIfNot(graphNode.out, &GraphNode[rune]{label: r, pos: Point{x: j, y: i + 1}})
		}
	}
	if j < cols-1 {
		// right
		r := grid[i][j+1]
		rightSli := rightMap[graphNode.label]
		if slices.Contains(rightSli, r) {
			fmt.Printf("Can connect right %c and %c\n", graphNode.label, r)
			graphNode.out = appendIfNot(graphNode.out, &GraphNode[rune]{label: r, pos: Point{x: j + 1, y: i}})
		}
	}
	if i > 0 {
		// up
		r := grid[i-1][j]
		upSli := upMap[graphNode.label]
		if slices.Contains(upSli, r) {
			fmt.Printf("Can connect up %c and %c\n", graphNode.label, r)
			graphNode.out = appendIfNot(graphNode.out, &GraphNode[rune]{label: r, pos: Point{x: j, y: i - 1}})
		}
	}

	if j > 0 {
		// left
		r := grid[i][j-1]
		leftSli := leftMap[graphNode.label]
		if slices.Contains(leftSli, r) {
			fmt.Printf("Can connect left %c and %c\n", graphNode.label, r)
			graphNode.out = appendIfNot(graphNode.out, &GraphNode[rune]{label: r, pos: Point{x: j - 1, y: i}})
		}
	}

	visited[graphNode] = true

	for _, n := range graphNode.out {
		explore(n, grid, visited)
	}

}

func appendIfNot[L string | rune](nodes []*GraphNode[L], node *GraphNode[L]) []*GraphNode[L] {

	if !slices.Contains(nodes, node) {
		nodes = append(nodes, node)
	}

	return nodes

}
