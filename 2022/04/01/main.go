package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	start int
	end   int
}

func parseSections(assignments string) (Section, Section) {
	pairs := strings.Split(assignments, ",")
	pair1 := strings.Split(pairs[0], "-")
	pair2 := strings.Split(pairs[1], "-")

	a, _ := strconv.Atoi(pair1[0])
	b, _ := strconv.Atoi(pair1[1])
	c, _ := strconv.Atoi(pair2[0])
	d, _ := strconv.Atoi(pair2[1])

	return Section{a, b}, Section{c, d}
}

func doesOneStrictlyIncludeTheOther(sec1, sec2 Section) bool {
	return (sec1.start <= sec2.start && sec1.end >= sec2.end) || (sec1.start >= sec2.start && sec1.end <= sec2.end)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	overlaps := 0
	for scanner.Scan() {
		sec1, sec2 := parseSections(scanner.Text())
		if doesOneStrictlyIncludeTheOther(sec1, sec2) {
			overlaps++
		}
	}

	fmt.Println(overlaps)
}
