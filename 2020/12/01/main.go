package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ship := Point{0, 0, 'E'}

	for scanner.Scan() {
		line := scanner.Text()
		towards := line[0]
		units, _ := strconv.Atoi(line[1:])

		ship.Move(Direction{towards, units})
	}

	fmt.Println(abs(ship.x) + abs(ship.y))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
