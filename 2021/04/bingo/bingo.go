package bingo

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Bingo map[int]struct{}

func (b *Bingo) String() string {
	var s string
	for n := range *b {
		s += fmt.Sprintf("%d ", n)
	}

	return s
}

func parseNumbers(strNumbers []string) []int {
	numbers := make([]int, 0, len(strNumbers))

	for _, sn := range strNumbers {
		number, _ := strconv.Atoi(sn)
		numbers = append(numbers, number)
	}

	return numbers
}

func Build(r io.Reader) ([]int, []Map) {
	reg := regexp.MustCompile(" +")

	scanner := bufio.NewScanner(r)

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

	return draws, maps
}
