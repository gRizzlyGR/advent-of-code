package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Range struct {
	lowerHalf float64
	upperHalf float64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxId := -1.0

	for scanner.Scan() {
		positions := scanner.Text()

		row := Range{0, 127}
		for i := 0; i < 7; i++ {
			switch positions[i] {
			case 'F':
				row.upperHalf = math.Floor((row.lowerHalf + row.upperHalf) / 2)
			case 'B':
				row.lowerHalf = math.Ceil((row.lowerHalf + row.upperHalf) / 2)
			}
		}

		col := Range{0, 7}
		for i := 7; i < 10; i++ {
			switch positions[i] {
			case 'L':
				col.upperHalf = math.Floor((col.lowerHalf + col.upperHalf) / 2)
			case 'R':
				col.lowerHalf = math.Ceil((col.lowerHalf + col.upperHalf) / 2)
			}
		}

		// lowerHalf and upperHald will be the same at this point
		id := row.lowerHalf*8 + col.lowerHalf

		if id > maxId {
			maxId = id
		}
	}

	fmt.Println(maxId)
}
