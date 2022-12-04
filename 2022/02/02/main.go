package main

import (
	"bufio"
	"fmt"
	"os"
)

// R == Rock, P == Paper, S == Scissors
var letterToShape = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
}

var shapeToScore = map[string]int{
	"R": 1,
	"P": 2,
	"S": 3,
}

// The key is the combination between the desired outcome (X == Lose, Y == Draw,
// Z == Win) and the foe's shape, and the value is the shape I have to play to
// make the outcome happen
var roundToShape = map[string]string{
	"XR": "S",
	"YR": "R",
	"ZR": "P",
	"XP": "R",
	"YP": "P",
	"ZP": "S",
	"XS": "P",
	"YS": "S",
	"ZS": "R",
}

var roundToScore = map[string]int{
	"RS": 6,
	"SP": 6,
	"PR": 6,
	"RR": 3,
	"SS": 3,
	"PP": 3,
}

func play(foe string, outcome string) int {
	foeShape := letterToShape[foe]
	shapeToPlay := roundToShape[outcome+foeShape]

	myScore := shapeToScore[shapeToPlay]
	roundScore := roundToScore[shapeToPlay+foeShape]

	return myScore + roundScore
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		foe := string(line[0])
		outcome := string(line[2])
		totalScore += play(foe, outcome)
	}

	fmt.Println(totalScore)
}
