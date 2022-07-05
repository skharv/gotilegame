package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/globals"
	"github.com/skharv/tilegame/scenes"
)

type Game struct {
	sceneManager *scenes.Manager
}

var (
	newTime, oldTime int64
)

func (g *Game) Init() {
	ebiten.SetWindowSize(globals.ScreenWidth, globals.ScreenHeight)
	ebiten.SetWindowTitle("TileMap")
}

func (g *Game) Update() error {
	oldTime = newTime
	newTime = time.Now().UnixNano()
	deltaTime := float64((newTime-oldTime)/1000000) * 0.001

	if g.sceneManager == nil {
		g.sceneManager = &scenes.Manager{}
		g.sceneManager.GoTo(&scenes.Game{}, 0)
	}

	g.sceneManager.ReadInput()
	if err := g.sceneManager.Update(deltaTime); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return globals.ScreenWidth, globals.ScreenHeight
}
