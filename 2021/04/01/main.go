package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseNumbers(strNumbers []string) []int {
	numbers := make([]int, 0, len(strNumbers))

	for _, sn := range strNumbers {
		number, _ := strconv.Atoi(sn)
		numbers = append(numbers, number)
	}

	return numbers
}

func draw(draws []int, maps []Map) (int, Bingo, Map) {
	for _, draw := range draws {
		for _, m := range maps {
			if cell, ok := m[draw]; ok {
				cell.Mark()
				bingo := cell.findBingo()

				if bingo != nil {
					return cell.number, bingo, m
				}
			}
		}
	}

	return -1, nil, nil
}

func computeScore(n int, b Map) int {
	var score int
	for _, cell := range b {
		if !cell.marked {
			score += cell.number
		}
	}

	return score * n
}

func main() {
	reg := regexp.MustCompile(" +")

	// file, _ := ioutil.ReadFile("../test_input")
	// scanner := bufio.NewScanner(bytes.NewReader(file))
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	draws := parseNumbers(strings.Split(scanner.Text(), ","))

	// Skip the unique empty line between numbers and boards
	scanner.Scan()

	var boards []Board
	var board Board
	var upRow Row

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			boards = append(boards, board)
			board = make(Board, 0)
			upRow = make(Row, 0)
			continue
		}

		strNumbers := reg.Split(strings.TrimSpace(line), -1)
		numbers := parseNumbers(strNumbers)
		row := BuildRow(upRow, numbers)
		board = append(board, row)
		upRow = row
	}

	// Append last board built
	boards = append(boards, board)
	maps := ToMaps(boards)
	n, bingo, m := draw(draws, maps)

	fmt.Println(bingo.String())
	fmt.Println(computeScore(n, m))
}
