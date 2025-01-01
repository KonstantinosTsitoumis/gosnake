package utils

import (
	"gosnake/game/constants"
	"math/rand/v2"
)

func GenerateRandomXY() (int, int) {
	x := rand.IntN(constants.ScreenWidth / constants.GridSize)
	y := rand.IntN(constants.ScreenHeight / constants.GridSize)

	return x, y
}
