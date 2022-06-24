package main

type Position struct {
	X int
	Y int
}

func main() {
	p := Position{X: 10, Y: 20}
	p.Move(5, -5)
}

func (m *Position) Move(dx, dy int) {
	m.X = m.X + dx
	m.Y = m.Y + dy
}
