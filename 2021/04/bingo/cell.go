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

func (n *Cell) String() string {
	leftCell, rightCell, upCell, downCell := -1, -1, -1, -1
	if n.left != nil {
		leftCell = n.left.number
	}

	if n.right != nil {
		rightCell = n.right.number
	}

	if n.up != nil {
		upCell = n.up.number
	}

	if n.down != nil {
		downCell = n.down.number
	}

	s := fmt.Sprintf("\t%d\n", upCell) + fmt.Sprintf("%d\t%d\t%d\n", leftCell, n.number, rightCell) + fmt.Sprintf("\t%d\n", downCell)

	return s
}

func (b *Cell) Mark() {
	b.marked = true
}

func (b *Cell) findInColumn() Bingo {
	bingo := make(Bingo)
	next := b

	for next != nil && next.marked {
		bingo[next.number] = struct{}{}
		next = next.up
	}

	next = b

	for next != nil && next.marked {
		bingo[next.number] = struct{}{}
		next = next.down
	}

	return bingo
}

func (b *Cell) findInRow() Bingo {
	bingo := make(Bingo)
	next := b

	for next != nil && next.marked {
		bingo[next.number] = struct{}{}
		next = next.left
	}

	next = b

	for next != nil && next.marked {
		bingo[next.number] = struct{}{}
		next = next.right
	}

	return bingo
}

func (b *Cell) findBingo() Bingo {
	bingo := b.findInColumn()

	if len(bingo) == 5 {
		return bingo
	}

	bingo = b.findInRow()

	if len(bingo) == 5 {
		return bingo
	}

	return nil
}
