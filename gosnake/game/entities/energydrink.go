package entities

import (
	"gosnake/game/constants"
	"gosnake/game/utils"
	"time"
)

type EnergyDrink struct {
	Position Position
	Cooldown *time.Time
	isEaten  bool
}

func NewEnergyDrink() EnergyDrink {
	y, x := utils.GenerateRandomXY()

	return EnergyDrink{
		Position: Position{
			X: x,
			Y: y,
		},
		Cooldown: nil,
		isEaten:  false,
	}
}

func (energyDrink *EnergyDrink) Update(lastUpdate time.Time) {
	if !energyDrink.isEaten {
		return
	}

	if energyDrink.Position != OutOfGrid {
		energyDrink.Position = OutOfGrid
	}
	if energyDrink.Cooldown == nil {
		newCooldown := time.Now().Add(constants.EnergyDrinkCooldown)
		energyDrink.Cooldown = &newCooldown
	}

	if !energyDrink.Cooldown.Before(time.Now()) {
		return
	}

	x, y := utils.GenerateRandomXY()

	energyDrink.Position = Position{
		X: x,
		Y: y,
	}
	energyDrink.isEaten = false
	energyDrink.Cooldown = nil
}
