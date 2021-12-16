package vent

import (
	"fmt"
	"strings"
)

type Line struct {
	From *Point
	To   *Point
}

func NewLine(s string) *Line {
	parts := strings.Split(s, " ")
	// We ignore part[1] since is "->"
	from := NewPoint(parts[0])
	to := NewPoint(parts[2])

	return &Line{from, to}
}

func (l *Line) String() string {
	return fmt.Sprintf("%v -> %v", l.From, l.To)
}

func minMax(a int, b int) (int, int) {
	if a < b {
		return a, b
	}

	return b, a
}

func (l *Line) RetrieveVerticalPoints() Points {
	points := *NewPoints()

	if l.From.X == l.To.X {
		min, max := minMax(l.From.Y, l.To.Y)

		for i := min; i <= max; i++ {
			p := &Point{l.From.X, i}
			points.AddPoint(p)
		}
	}

	return points
}

func (l *Line) RetrieveHorizontalPoints() Points {
	points := *NewPoints()

	if l.From.Y == l.To.Y {
		min, max := minMax(l.From.X, l.To.X)

		for i := min; i <= max; i++ {
			p := &Point{i, l.From.Y}
			points.AddPoint(p)
		}
	}

	return points
}

func (l *Line) RetrieveDiagonalPoints() Points {
	points := *NewPoints()

	if l.From.X < l.To.X {
		y := l.From.Y
		if l.From.Y < l.To.Y {
			// to bottom right: inc x, inc Y
			for x := l.From.X; x <= l.To.X; x++ {
				p := &Point{x, y}
				points.AddPoint(p)
				y++
			}
		}

		if l.From.Y > l.To.Y {
			// to top right: inc x, dec y
			for x := l.From.X; x <= l.To.X; x++ {
				p := &Point{x, y}
				points.AddPoint(p)
				y--
			}
		}
	}

	if l.From.X > l.To.X {
		y := l.From.Y
		if l.From.Y < l.To.Y {
			// to bottom left: dec x, inc y
			for x := l.From.X; x >= l.To.X; x-- {
				p := &Point{x, y}
				points.AddPoint(p)
				y++
			}
		}

		if l.From.Y > l.To.Y {
			// to top left: dec x, dec y
			for x := l.From.X; x >= l.To.X; x-- {
				p := &Point{x, y}
				points.AddPoint(p)
				y--
			}
		}
	}

	return points
}
