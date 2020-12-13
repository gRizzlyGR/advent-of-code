package main

type Mover interface {
	Move(dir Direction)
	turn(dir Direction)
	goStraightRight()
	goStraightLeft()
	goOpposite()
}

type Point struct {
	x       int
	y       int
	towards byte
}

type Direction struct {
	towards byte
	units   int
}

func (p *Point) Move(dir Direction) {
	switch dir.towards {
	case 'N':
		p.y += dir.units
	case 'S':
		p.y -= dir.units
	case 'E':
		p.x += dir.units
	case 'W':
		p.x -= dir.units
	case 'L', 'R':
		p.turn(dir)
	case 'F':
		p.Move(Direction{p.towards, dir.units})
	default:
		panic(dir)
	}
}

func (p *Point) turn(dir Direction) {
	switch dir {
	case Direction{'R', 90}, Direction{'L', 270}:
		p.goStraightRight()
	case Direction{'R', 270}, Direction{'L', 90}:
		p.goStraightLeft()
	case Direction{'R', 180}, Direction{'L', 180}:
		p.goOpposite()
	default:
		panic("Unable to turn")
	}
}

func (p *Point) goStraightRight() {
	switch p.towards {
	case 'N':
		p.towards = 'E'
	case 'S':
		p.towards = 'W'
	case 'E':
		p.towards = 'S'
	case 'W':
		p.towards = 'N'
	default:
		panic(string(p.towards))
	}
}

func (p *Point) goStraightLeft() {
	switch p.towards {
	case 'N':
		p.towards = 'W'
	case 'S':
		p.towards = 'E'
	case 'E':
		p.towards = 'N'
	case 'W':
		p.towards = 'S'
	default:
		panic(string(p.towards))
	}
}

func (p *Point) goOpposite() {
	switch p.towards {
	case 'N':
		p.towards = 'S'
	case 'S':
		p.towards = 'N'
	case 'E':
		p.towards = 'W'
	case 'W':
		p.towards = 'E'
	default:
		panic(string(p.towards))
	}
}
