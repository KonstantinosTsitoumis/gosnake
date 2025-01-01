package entities

type Position struct {
	X, Y int
}

type direction int

var OutOfGrid = Position{X: -10, Y: -10}

const (
	Up direction = iota
	Down
	Right
	Left
)

var Direction = map[direction]string{
	Up:    "up",
	Down:  "down",
	Right: "right",
	Left:  "left",
}
