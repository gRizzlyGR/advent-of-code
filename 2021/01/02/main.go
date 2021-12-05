package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := make([]int, 0, 3)

	// To skip the first window
	last := -1
	var increased int

	for scanner.Scan() {
		measure, _ := strconv.Atoi(scanner.Text())

		if len(queue) < 3 {
			// Enqueue
			queue = append(queue, measure)
		}

		if len(queue) == 3 {
			var window int
			for _, val := range queue {
				window += val
			}

			if last != -1 && window > last {
				increased++
			}

			last = window

			// Dequeue
			queue = queue[1:]
		}
	}

	fmt.Println(increased)
}
