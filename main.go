package main

import (
	"game/game"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// Initialize the game
	g := game.NewGame()

	ebiten.SetWindowSize(1280, 768)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
