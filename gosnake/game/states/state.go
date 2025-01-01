package states

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Context interface{}

type State interface {
	Update() (string, error)
	Draw(screen *ebiten.Image)
}
