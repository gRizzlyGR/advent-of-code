package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	m := make(map[string]struct{})
	tot := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		if len(line) == 0 {
			tot += len(m)
			m = make(map[string]struct{})
		} else {
			for _, letter := range line {
				m[letter] = struct{}{}
			}
		}
	}

	fmt.Println(tot)
}
