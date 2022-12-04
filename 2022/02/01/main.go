package main

import (
	"bufio"
	"fmt"
	"os"
)

type Shape struct {
	name  string
	score int
}

// R == Rock, P == Paper, S == Scissors
var letterToShape = map[string]Shape{
	"A": {"R", -1},
	"B": {"P", -1},
	"C": {"S", -1},
	"X": {"R", 1},
	"Y": {"P", 2},
	"Z": {"S", 3},
}

var roundToScore = map[string]int{
	"RS": 6,
	"SP": 6,
	"PR": 6,
	"RR": 3,
	"SS": 3,
	"PP": 3,
}

func play(foe string, me string) int {
	foeShape := letterToShape[foe]
	myShape := letterToShape[me]

	roundScore := roundToScore[myShape.name+foeShape.name]

	return roundScore + myShape.score
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		foe := string(line[0])
		me := string(line[2])
		totalScore += play(foe, me)
	}

	fmt.Println(totalScore)
}
