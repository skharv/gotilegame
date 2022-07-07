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

type PlayState int

type Game struct {
	tileMap       *tilemap.TileMap
	entityManager *entities.Manager
	player        *player.Player
	random        *rand.Rand
	currentState  PlayState
}

const (
	playerInput   PlayState = 0
	resolveAction PlayState = 1
)

func (s *Game) Init() {
	s.currentState = playerInput

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
	switch s.currentState {
	case playerInput:
		s.player.ReadInputs()
	case resolveAction:
	}
}

func (s *Game) Update(state *GameState, deltaTime float64) error {
	switch s.currentState {
	case playerInput:
		s.player.Update(s.tileMap, s.entityManager)
		if !s.player.IsReady() {
			s.tileMap.SetAction(s.player.GetLastPlayed())
			s.currentState = resolveAction
		}
	case resolveAction:
		s.tileMap.Update()
		if s.tileMap.IsResolved() {
			s.player.GetReady()
			s.currentState = playerInput
		}
	}
	s.entityManager.Update()

	return nil
}

func (s *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{126, 158, 153, 255})
	s.player.Draw(screen)
	s.tileMap.Draw(screen)
	s.entityManager.Draw(screen)
}
