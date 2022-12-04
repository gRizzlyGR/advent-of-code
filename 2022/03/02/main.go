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

func findBadge(group []string) (rune, error) {

	for _, a := range group[0] {
		for _, b := range group[1] {
			for _, c := range group[2] {
				if a == b && b == c {
					return a, nil
				}
			}
		}
	}

	return -1, errors.New("no badge found")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalPriorities := 0

	group := make([]string, 0, 3)

	for scanner.Scan() {
		items := scanner.Text()
		group = append(group, items)

		if len(group) == 3 {
			badge, err := findBadge(group)
			if err != nil {
				panic(err)
			}

			group = make([]string, 0)

			totalPriorities += computePriority(badge)
		}
	}

	fmt.Println(totalPriorities)
}
