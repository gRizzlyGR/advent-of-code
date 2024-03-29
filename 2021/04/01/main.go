package main

import (
	"bingo"
	"fmt"
	"os"
)

func draw(draws []int, maps []bingo.Map) (int, bingo.Bingo, bingo.Map) {
	for _, draw := range draws {
		for _, m := range maps {
			if cell, ok := m[draw]; ok {
				cell.Mark()
				bingo := cell.FindBingo()

				if bingo != nil {
					return cell.Number(), bingo, m
				}
			}
		}
	}

	return -1, nil, nil
}

func computeScore(n int, b bingo.Map) int {
	var score int
	for _, cell := range b {
		if !cell.IsMarked() {
			score += cell.Number()
		}
	}

	return score * n
}

func main() {
	draws, maps := bingo.Build(os.Stdin)
	n, bingo, m := draw(draws, maps)

	fmt.Println(n)
	fmt.Println(bingo.String())
	fmt.Println(computeScore(n, m))
}
