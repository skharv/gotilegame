package scenes

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/entities"
	"github.com/skharv/tilegame/player"
	"github.com/skharv/tilegame/tilemap"
	"github.com/skharv/tilegame/ui"
)

type Game struct {
	tileMap       *tilemap.TileMap
	entityManager *entities.Manager
	cursor        *ui.Cursor
	player, comp  *player.Player
	controller    *player.Controller
	random        *rand.Rand
}

func (s *Game) Init() {
	s.entityManager = &entities.Manager{}
	s.entityManager.Init()

	s.tileMap = &tilemap.TileMap{}
	s.tileMap.Init()

	s.cursor = &ui.Cursor{}
	s.cursor.Init()

	s.player = &player.Player{}
	s.comp = &player.Player{}

	s.controller = &player.Controller{}
	s.controller.Init(s.player)

	seed := rand.NewSource(time.Now().UnixNano())
	s.random = rand.New(seed)
}

func (s *Game) ReadInput() {
	s.controller.ReadInputs()
	s.cursor.ReadInputs()
}

func (s *Game) Update(state *GameState, deltaTime float64) error {
	s.tileMap.Update()
	s.cursor.Update(s.tileMap)
	s.controller.Update(s.tileMap)

	if s.cursor.IsClicked() && s.cursor.IsVisible() {
		i, j := s.cursor.GetPosition()
		obj, register := s.player.CreateObject(0, s.tileMap.WorldToTile(int(i), int(j)))
		if register {
			s.entityManager.Register(obj)
		}
	}

	if s.cursor.IsAltClicked() && s.cursor.IsVisible() {
		i, j := s.cursor.GetPosition()
		obj, register := s.player.CreateObject(1, s.tileMap.WorldToTile(int(i), int(j)))
		if register {
			s.entityManager.Register(obj)
		}
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
