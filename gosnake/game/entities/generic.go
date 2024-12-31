package entities

type Position struct {
	X, Y int
}

type direction int

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
