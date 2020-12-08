package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bagPattern := "bag(color(%s))"
	containsPattern := "contains(%s, %s, %s)."

	for scanner.Scan() {
		line := scanner.Text()

		// pale olive bags contain 1 bright yellow bag, 1 mirrored salmon bag.
		parts := strings.Split(line, " contain ")

		container := strings.ReplaceAll(parts[0], "bags", "")

		color := strings.TrimSpace(container)
		color = strings.ReplaceAll(color, " ", "_")
		container = fmt.Sprintf(bagPattern, color)

		containedBags := strings.Split(parts[1], ", ")

		for _, bag := range containedBags {
			bag = strings.ReplaceAll(bag, "bags", "")
			bag = strings.ReplaceAll(bag, "bag", "")
			bag = strings.ReplaceAll(bag, ".", "")
			bag = strings.TrimSpace(bag)

			qty := string(bag[0])
			color := strings.ReplaceAll(bag[2:], " ", "_")

			// This was originally "no other bag", so we reached the last level
			// and we do not store it in the kb
			if strings.Contains(color, "other") {
				continue
			}

			contained := fmt.Sprintf(bagPattern, color)
			contains := fmt.Sprintf(containsPattern, container, contained, qty)

			fmt.Println(contains)
		}
	}
}
