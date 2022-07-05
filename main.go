package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/game"
)

func main() {
	game := &game.Game{}
	game.Init()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
