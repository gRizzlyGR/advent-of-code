package main

import (
	"bufio"
	"fmt"
	"os"
	"vent"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	c := make(vent.Covering)

	for scanner.Scan() {
		l := vent.NewLine(scanner.Text())
		c.Fill(l.RetrieveVerticalPoints())
		c.Fill(l.RetrieveHorizontalPoints())
		c.Fill(l.RetrieveDiagonalPoints())
	}

	fmt.Println(c.CountOverlaps())
}
