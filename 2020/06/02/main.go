package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	m := make(map[string]int)
	tot := 0
	nPeopleInGroup := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		if len(line) == 0 {
			for _, qty := range m {
				if qty == nPeopleInGroup {
					tot++
				}
			}

			m = make(map[string]int)
			nPeopleInGroup = 0
		} else {
			for _, letter := range line {
				m[letter]++
			}

			nPeopleInGroup++
		}
	}

	fmt.Println(tot)
}
