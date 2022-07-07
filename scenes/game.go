package scenes

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/entities"
	"github.com/skharv/tilegame/player"
	"github.com/skharv/tilegame/tilemap"
)

type Game struct {
	tileMap       *tilemap.TileMap
	entityManager *entities.Manager
	player        *player.Player
	random        *rand.Rand
}

func (s *Game) Init() {
	s.entityManager = &entities.Manager{}
	s.entityManager.Init()

	s.tileMap = &tilemap.TileMap{}
	s.tileMap.Init()

	s.player = &player.Player{}
	s.player.Init()

	seed := rand.NewSource(time.Now().UnixNano())
	s.random = rand.New(seed)
}

func (s *Game) ReadInput() {
	s.player.ReadInputs()
}

func (s *Game) Update(state *GameState, deltaTime float64) error {
	s.tileMap.Update()
	s.player.Update(s.tileMap, s.entityManager)
	s.entityManager.Update()

	return nil
}

func (s *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{126, 158, 153, 255})
	s.player.Draw(screen)
	s.tileMap.Draw(screen)
	s.entityManager.Draw(screen)
}
