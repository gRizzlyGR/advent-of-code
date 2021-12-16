package main

import (
	"bingo"
	"fmt"
	"os"
)

func draw(draws []int, maps []bingo.Map) (int, bingo.Bingo, bingo.Map) {
	wins := make(map[int]struct{})

	for _, draw := range draws {
		for i, m := range maps {
			if cell, ok := m[draw]; ok {
				cell.Mark()
				bingo := cell.FindBingo()

				if bingo != nil {
					// Record if a board has won, just once
					if _, ok := wins[i]; !ok {
						wins[i] = struct{}{}
					}

					// If all boards have won at least once, return the current
					// one
					if len(wins) == len(maps) {
						return cell.Number(), bingo, m
					}
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
