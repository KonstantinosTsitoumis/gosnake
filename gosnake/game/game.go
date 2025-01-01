package game

import (
	"gosnake/game/constants"
	"gosnake/game/states"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	state states.State
}

func NewGame() Game {
	return Game{state: states.NewMenuState()}
}

func (g *Game) Update() error {
	changeState, err := g.state.Update()

	switch changeState {
	case "run":
		g.state = states.NewRunState()
		break
	case "terminate":
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}
