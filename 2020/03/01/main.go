package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := make([]string, 0)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	trees := 0
	col := 3

	for row := 1; row < len(grid); row++ {
		area := grid[row]
		obstacle := area[col]

		if obstacle == '#' {
			trees++
		}

		col += 3

		if col >= len(area) {
			col -= len(area)
		}
	}

	fmt.Println(trees)
}
