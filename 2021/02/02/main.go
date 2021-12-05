package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x   int
	y   int
	aim int
}

type command struct {
	direction string
	units     int
}

func (p *point) pilot(c *command) {
	switch c.direction {
	case "forward":
		p.x += c.units
		p.y += c.units * p.aim
	case "up":
		p.aim -= c.units
	case "down":
		p.aim += c.units
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	point := point{0, 0, 0}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		direction := parts[0]
		unit, _ := strconv.Atoi(parts[1])

		point.pilot(&command{direction, unit})
	}

	fmt.Println(point.x * point.y)
}
