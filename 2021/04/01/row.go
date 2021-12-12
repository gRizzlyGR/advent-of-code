package main

import "fmt"

type Row []*Cell

func BuildRow(upRow Row, numbers []int) Row {
	row := make(Row, 0)

	var leftcell *Cell

	for i, number := range numbers {
		currcell := &Cell{
			number: number,
			up:     nil,
			down:   nil,
			left:   nil,
			right:  nil,
		}

		if len(upRow) > 0 {
			upcell := upRow[i]
			currcell.up = upcell
			upcell.down = currcell
		}

		if leftcell != nil {
			currcell.left = leftcell
			leftcell.right = currcell
			row = append(row, leftcell)
		}

		leftcell = currcell
	}

	// Append the last hanging one
	row = append(row, leftcell)

	return row
}

func (b *Row) String() string {
	var s string
	for _, cell := range *b {
		s += fmt.Sprintf("%d ", cell.number)
	}

	return s
}
