package states

import (
	"bytes"
	"fmt"
	"gosnake/game/constants"
	"gosnake/game/entities"
	"gosnake/game/utils"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type RunState struct {
	Context RunContext
}

type RunContext struct {
	snake       entities.Snake
	apple       entities.Apple
	energyDrink entities.EnergyDrink
	lastUpdate  time.Time
	gameover    bool
	score       int
}

func NewRunState() *RunState {
	snake := entities.NewSnake(
		constants.StartingPointX,
		constants.StartingPointY,
	)

	apple := entities.NewApple()
	energyDrink := entities.NewEnergyDrink()

	return &RunState{
		Context: RunContext{
			snake:       snake,
			apple:       apple,
			energyDrink: energyDrink,
			lastUpdate:  time.Now(),
			score:       0,
		},
	}
}

func (r *RunState) Update() (string, error) {
	if r.Context.gameover {
		return "gameover", nil
	}

	if time.Since(r.Context.lastUpdate) < constants.GameSpeed {
		return "", nil
	}

	r.Context.lastUpdate = time.Now()

	direction := r.Context.snake.GetDirection()

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		direction = entities.Direction[entities.Up]
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		direction = entities.Direction[entities.Down]
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		direction = entities.Direction[entities.Right]
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		direction = entities.Direction[entities.Left]
	}

	r.Context.snake.Update(direction, &r.Context.apple, &r.Context.energyDrink, &r.Context.gameover)
	r.Context.apple.Update(&r.Context.score)
	r.Context.energyDrink.Update(r.Context.lastUpdate)

	return "", nil
}

func (r *RunState) Draw(screen *ebiten.Image) {
	s, _ := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(0.0, 0.0)
	text.Draw(screen, fmt.Sprintf("score %d", r.Context.score), &text.GoTextFace{
		Source: s,
		Size:   24,
	}, op)

	if r.Context.gameover {
		op.GeoM.Translate(200.0, 200.0)
		text.Draw(screen, "Game Over", &text.GoTextFace{
			Source: s,
			Size:   24,
		}, op)

		return
	}

	snakePositions := r.Context.snake.GetSnakeBodyPositions()

	for i := range len(snakePositions) {
		x := snakePositions[i].X * constants.GridSize
		y := snakePositions[i].Y * constants.GridSize

		vector.DrawFilledRect(
			screen,
			float32(x),
			float32(y),
			constants.GridSize,
			constants.GridSize,
			color.White,
			true,
		)
	}
	{
		applePosition := r.Context.apple.Position

		x, y := utils.PositionToGrid(applePosition.X, applePosition.Y)

		red := color.RGBA{R: 255, G: 0, B: 0, A: 255}

		vector.DrawFilledRect(
			screen,
			float32(x),
			float32(y),
			constants.GridSize,
			constants.GridSize,
			red,
			true,
		)
	}
	{
		energyDrinkPosition := r.Context.energyDrink.Position

		x, y := utils.PositionToGrid(energyDrinkPosition.X, energyDrinkPosition.Y)

		green := color.RGBA{R: 0, G: 255, B: 0, A: 255}

		vector.DrawFilledRect(
			screen,
			float32(x),
			float32(y),
			constants.GridSize,
			constants.GridSize,
			green,
			true,
		)
	}
}
