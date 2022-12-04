package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func computePriority(item rune) int {
	if item <= 'Z' {
		return 52 - int(('Z' - item))
	}

	return 26 - int(('z' - item))
}

func buildCompartments(items string) (string, string) {
	size := len(items) / 2
	first := items[:size]
	second := items[size:]

	return first, second
}

func findCommonItem(first string, second string) {
	for _, a := range first {
		for _, b := range second {
			if a == b {
				return a, nil
			}
		}
	}

	return -1, errors.New("no common item found")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalPriorities := 0
	for scanner.Scan() {
		items := scanner.Text()
		first, second := buildCompartments(items)
		item, err := findCommonItem(first, second)

		if err != nil {
			panic(err)
		}

		totalPriorities += computePriority(item)
	}

	fmt.Println(totalPriorities)
}
