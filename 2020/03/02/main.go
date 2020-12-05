package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	col int
	row int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := make([]string, 0)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	points := []Point{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	trees := 1

	for _, point := range points {
		trees *= getTrees(grid, point.row, point.col)
	}
	fmt.Println(trees)
}

func getTrees(grid []string, startRow int, startCol int) int {
	trees := 0

	col := startCol

	for row := startRow; row < len(grid); row += startRow {
		area := grid[row]
		obstacle := area[col]

		if obstacle == '#' {
			trees++
		}

		col += startCol

		if col >= len(area) {
			col -= len(area)
		}
	}

	return trees
}
