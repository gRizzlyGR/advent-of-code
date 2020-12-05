package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Range struct {
	lowerHalf float64
	upperHalf float64
}

func main() {
	seats := make(map[float64]struct{})

	scanner := bufio.NewScanner(os.Stdin)

	maxSeat := -1.0
	minSeat := math.MaxFloat64

	for scanner.Scan() {
		positions := scanner.Text()

		row := Range{0, 127}
		for i := 0; i < 7; i++ {
			switch positions[i] {
			case 'F':
				row.upperHalf = math.Floor((row.lowerHalf + row.upperHalf) / 2)
			case 'B':
				row.lowerHalf = math.Ceil((row.lowerHalf + row.upperHalf) / 2)
			}
		}

		col := Range{0, 7}
		for i := 7; i < 10; i++ {
			switch positions[i] {
			case 'L':
				col.upperHalf = math.Floor((col.lowerHalf + col.upperHalf) / 2)
			case 'R':
				col.lowerHalf = math.Ceil((col.lowerHalf + col.upperHalf) / 2)
			}
		}

		// lowerHalf and upperHald will be the same at this point
		seat := row.lowerHalf*8 + col.lowerHalf

		if seat > maxSeat {
			maxSeat = seat
		}

		if seat < minSeat {
			minSeat = seat
		}

		seats[seat] = struct{}{}
	}

	beforeSeat := -1.0
	afterSeat := -1.0
	
	for seat := range seats {
		if seat == maxSeat || seat == minSeat {
			continue
		}
		_, before := seats[seat-1]
		_, after := seats[seat+1]
		
		if !after {
			beforeSeat = seat
		}
		
		if !before {
			afterSeat = seat
		}
	}
	
	personalSeat := (beforeSeat + afterSeat) / 2
	fmt.Println(personalSeat)

}
