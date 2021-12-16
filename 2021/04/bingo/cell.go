package bingo

import "fmt"

type Cell struct {
	number int
	up     *Cell
	down   *Cell
	left   *Cell
	right  *Cell
	marked bool
}

func (c *Cell) String() string {
	leftCell, rightCell, upCell, downCell := -1, -1, -1, -1
	if c.left != nil {
		leftCell = c.left.number
	}

	if c.right != nil {
		rightCell = c.right.number
	}

	if c.up != nil {
		upCell = c.up.number
	}

	if c.down != nil {
		downCell = c.down.number
	}

	s := fmt.Sprintf("\t%d\n", upCell) + fmt.Sprintf("%d\t%d\t%d\n", leftCell, c.number, rightCell) + fmt.Sprintf("\t%d\n", downCell)

	return s
}

func (c *Cell) Mark() {
	c.marked = true
}

func (c *Cell) IsMarked() bool {
	return c.marked
}

func (c *Cell) Number() int {
	return c.number
}

func (c *Cell) FindInColumn() Bingo {
	bingo := make(Bingo)
	next := c

	for next != nil && next.marked {
		bingo[next.number] = struct{}{}
		next = next.up
	}

	next = c

	for next != nil && next.marked {
		bingo[next.number] = struct{}{}
		next = next.down
	}

	return bingo
}

func (c *Cell) FindInRow() Bingo {
	bingo := make(Bingo)
	next := c

	for next != nil && next.marked {
		bingo[next.number] = struct{}{}
		next = next.left
	}

	next = c

	for next != nil && next.marked {
		bingo[next.number] = struct{}{}
		next = next.right
	}

	return bingo
}

func (c *Cell) FindBingo() Bingo {
	bingo := c.FindInColumn()

	if len(bingo) == 5 {
		return bingo
	}

	bingo = c.FindInRow()

	if len(bingo) == 5 {
		return bingo
	}

	return nil
}
