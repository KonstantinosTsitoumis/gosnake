package utils

import (
	"gosnake/game/constants"
)

func PositionToGrid(x int, y int) (int, int) {
	return x * constants.GridSize, y * constants.GridSize
}
