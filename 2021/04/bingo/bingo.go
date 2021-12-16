package bingo

import "fmt"

type Bingo map[int]struct{}

func (b *Bingo) String() string {
	var s string
	for n := range *b {
		s += fmt.Sprintf("%d ", n)
	}

	return s
}
