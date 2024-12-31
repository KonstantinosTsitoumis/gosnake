package game

import (
	"bytes"
	"gosnake/game/constants"
	"gosnake/game/entities"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	snake      entities.Snake
	apple      entities.Apple
	lastUpdate time.Time
	gameover   bool
}

func NewGame() Game {
	snake := entities.NewSnake(
		constants.StartingPointX,
		constants.StartingPointY,
	)

	apple := entities.NewApple()

	return Game{
		snake:      snake,
		apple:      apple,
		lastUpdate: time.Now(),
	}
}

func (g *Game) Update() error {
	if g.gameover {
		return nil
	}

	if time.Since(g.lastUpdate) < constants.GameSpeed {
		return nil
	}

	g.lastUpdate = time.Now()

	direction := g.snake.GetDirection()

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

	g.snake.Update(direction, &g.apple, &g.gameover)

	if g.apple.Eaten {
		g.apple = entities.NewApple()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameover {
		s, _ := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))

		op := &text.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, "Game Over", &text.GoTextFace{
			Source: s,
			Size:   24,
		}, op)

		return
	}

	snakePositions := g.snake.GetSnakeBodyPositions()

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

	applePosition := g.apple.Position

	x := applePosition.X * constants.GridSize
	y := applePosition.Y * constants.GridSize

	vector.DrawFilledRect(
		screen,
		float32(x),
		float32(y),
		constants.GridSize,
		constants.GridSize,
		color.Opaque,
		true,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}
