package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var last string
	var increased int

	for scanner.Scan() {
		curr := scanner.Text()

		if curr > last {
			increased++
		}

		last = curr
	}

	fmt.Println(increased)
}
