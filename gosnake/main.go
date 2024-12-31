package main

import (
	"gosnake/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game_ := game.NewGame()

	x, y := game_.Layout(0, 0)

	ebiten.SetWindowSize(x, y)

	if err := ebiten.RunGame(&game_); err != nil {
		log.Fatal(err)
	}
}
