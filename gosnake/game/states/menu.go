package states

import (
	"bytes"
	"errors"
	"gosnake/game/constants"
	"gosnake/game/entities"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var ErrGameExit = errors.New("game exit")

type MenuState struct {
	Context MenuContext
}

type MenuContext struct {
	selectedOption int
	options        []string
	lastUpdate     time.Time
}

func NewMenuState() *MenuState {
	return &MenuState{
		Context: MenuContext{
			selectedOption: 0,
			options: []string{
				"Start Game",
				"Exit",
			},
		},
	}
}

func (m *MenuState) Update() (string, error) {
	if time.Since(m.Context.lastUpdate) < constants.MenuSpeed {
		return "", nil
	}

	m.Context.lastUpdate = time.Now()

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		m.Context.selectedOption = (m.Context.selectedOption + 1) % len(m.Context.options)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		m.Context.selectedOption = (m.Context.selectedOption - 1 + len(m.Context.options)) % len(m.Context.options)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		if m.Context.selectedOption == 0 {
			return "run", nil
		} else if m.Context.selectedOption == 1 {
			return "terminate", ErrGameExit
		}
	}

	return "", nil
}

func (m *MenuState) Draw(screen *ebiten.Image) {
	// Load font face
	s, _ := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	face := &text.GoTextFace{
		Source: s,
		Size:   24,
	}

	// Draw each menu option
	center := entities.Position{
		X: constants.ScreenWidth/2 - 200,
		Y: constants.ScreenHeight / 2,
	}

	for i, option := range m.Context.options {
		// Measure text width using Draw method
		// Set color
		colorScale := color.Gray{Y: 128} // Default color
		if i == m.Context.selectedOption {
			colorScale = color.Gray{Y: 255} // Highlight selected option
		}

		// Draw the text
		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(center.X), float64(center.Y)) // Set position

		center.Y += 100

		op.ColorScale.ScaleWithColor(colorScale)
		text.Draw(screen, option, face, op)
	}
}
