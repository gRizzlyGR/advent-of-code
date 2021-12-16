package vent

type Covering map[Point]int

func (c Covering) Fill(points Points) {
	for _, p := range points {
		c[*p]++
	}
}

func (c Covering) CountOverlaps() int {
	var overlaps int
	for _, v := range c {
		if v > 1 {
			overlaps++
		}
	}

	return overlaps
}
