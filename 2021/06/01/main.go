package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fish int
type School []Fish

func Plant(s string) School {
	parts := strings.Split(s, ",")

	school := make(School, len(parts))

	for i, part := range parts {
		fish, _ := strconv.Atoi(part)
		school[i] = Fish(fish)
	}

	return school
}

func Grow(school School, limit int) School {
	var growing School

	for i := 1; i <= limit; i++ {
		growing = make(School, 0)

		for _, fish := range school {
			if fish == 0 {
				growing = append(growing, 6, 8)
				continue
			}

			fish--
			growing = append(growing, fish)
		}

		school = growing
	}

	return growing
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	school := Plant(scanner.Text())

	grown := Grow(school, 18)
	fmt.Println(len(grown))
}
