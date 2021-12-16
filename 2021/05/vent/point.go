package vent

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPoint(s string) *Point {
	ss := strings.Split(s, ",")
	x, _ := strconv.Atoi(ss[0])
	y, _ := strconv.Atoi(ss[1])

	return &Point{x, y}
}

func (p *Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type Points []*Point

func NewPoints() *Points {
	points := make(Points, 0)
	return &points
}

func (pp *Points) AddPoint(point *Point) {
	*pp = append(*pp, point)
}

func (pp *Points) Append(points *Points) Points {
	return append(*pp, *points...)
}
