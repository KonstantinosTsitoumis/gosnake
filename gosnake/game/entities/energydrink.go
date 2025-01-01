package entities

import (
	"gosnake/game/constants"
	"gosnake/game/utils"
	"time"
)

type EnergyDrink struct {
	Position Position
	Cooldown time.Duration
	Eaten    bool
}

func NewEnergyDrink() EnergyDrink {
	y, x := utils.GenerateRandomXY()

	return EnergyDrink{
		Position: Position{
			X: x,
			Y: y,
		},
		Cooldown: constants.EnergyDrinkCooldown,
		Eaten:    false,
	}
}

func (energyDrink *EnergyDrink) Update(lastUpdate time.Time) {
	if !energyDrink.Eaten {
		return
	}

	if time.Since(lastUpdate) > constants.EnergyDrinkCooldown {
		return
	}

	newEnergyDrink := NewEnergyDrink()
	energyDrink = &newEnergyDrink
}
