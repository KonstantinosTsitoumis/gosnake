package entities

import (
	"gosnake/game/constants"
	"math/rand/v2"
)

type Apple struct {
	Position Position
	Points   int
	Eaten    bool
}

func NewApple() Apple {
	x := rand.IntN(constants.ScreenWidth / constants.GridSize)
	y := rand.IntN(constants.ScreenHeight / constants.GridSize)

	return Apple{
		Position: Position{
			X: x,
			Y: y,
		},
		Points: 1,
		Eaten:  false,
	}
}
