package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/entities"
	"github.com/skharv/tilegame/tilemap"
	"github.com/skharv/tilegame/ui"
)

type Game struct {
	tileMap       *tilemap.TileMap
	entityManager *entities.Manager
	cursor        *ui.Cursor
}

func (s *Game) Init() {
	s.entityManager = &entities.Manager{}
	s.entityManager.Init()

	s.tileMap = &tilemap.TileMap{}
	s.tileMap.Init()

	s.cursor = &ui.Cursor{}
	s.cursor.Init()
}

func (s *Game) ReadInput() {
	s.cursor.ReadInputs()
}

func (s *Game) Update(state *GameState, deltaTime float64) error {
	s.tileMap.Update()
	s.cursor.Update(s.tileMap)

	if s.cursor.IsClicked() && s.cursor.IsVisible() {
		obj := &entities.Object{}
		obj.Init()
		obj.SetSprite("images/redunit.png")
		obj.SetPosition(s.cursor.GetPosition())
		s.entityManager.Register(obj)
	}

	if s.cursor.IsAltClicked() && s.cursor.IsVisible() {
		obj := &entities.Object{}
		obj.Init()
		obj.SetSprite("images/blueunit.png")
		obj.SetPosition(s.cursor.GetPosition())
		s.entityManager.Register(obj)
	}

	s.entityManager.Update()

	return nil
}

func (s *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{126, 158, 153, 255})
	s.tileMap.Draw(screen)
	s.cursor.Draw(screen)
	s.entityManager.Draw(screen)
}
