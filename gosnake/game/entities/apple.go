package entities

import (
	"gosnake/game/utils"
)

type Apple struct {
	Position Position
	Points   int
	Eaten    bool
}

func NewApple() Apple {
	x, y := utils.GenerateRandomXY()

	return Apple{
		Position: Position{
			X: x,
			Y: y,
		},
		Points: 1,
		Eaten:  false,
	}
}

func (apple *Apple) Update() {
	if !apple.Eaten {
		return
	}

	x, y := utils.GenerateRandomXY()

	apple.Eaten = false
	apple.Position = Position{
		X: x,
		Y: y,
	}
}
