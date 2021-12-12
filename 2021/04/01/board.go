package main

import "fmt"

type Board []Row

type Map map[int]*Cell

func ToMaps(boards []Board) []Map {
	maps := make([]Map, 0)
	for _, board := range boards {
		m := make(Map)
		for _, row := range board {
			for _, cell := range row {
				m[cell.number] = cell
			}
		}
		maps = append(maps, m)
	}

	return maps
}

func (b *Board) Print() {
	for _, row := range *b {
		for _, cell := range row {
			fmt.Println(cell.String())
		}
	}
}
