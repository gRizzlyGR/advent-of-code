package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	start int
	end   int
}

func parseSections(pair string) (Assignment, Assignment) {
	assignments := strings.Split(pair, ",")
	assignment1 := strings.Split(assignments[0], "-")
	assignment2 := strings.Split(assignments[1], "-")

	sec1, _ := strconv.Atoi(assignment1[0])
	sec2, _ := strconv.Atoi(assignment1[1])
	sec3, _ := strconv.Atoi(assignment2[0])
	sec4, _ := strconv.Atoi(assignment2[1])

	return Assignment{sec1, sec2}, Assignment{sec3, sec4}
}

func doesOneStrictlyIncludeTheOther(assignment1, assignment2 Assignment) bool {
	return (assignment1.start <= assignment2.start && assignment1.end >= assignment2.end) || (assignment1.start >= assignment2.start && assignment1.end <= assignment2.end)
}

func areFullyDisjoined(assignment1, assignment2 Assignment) bool {
	return (assignment1.end < assignment2.start) || (assignment2.end < assignment1.start)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	overlaps := 0
	for scanner.Scan() {
		assignment1, assignment2 := parseSections(scanner.Text())
		if doesOneStrictlyIncludeTheOther(assignment1, assignment2) || !areFullyDisjoined(assignment1, assignment2) {
			overlaps++
		}
	}

	fmt.Println(overlaps)
}
