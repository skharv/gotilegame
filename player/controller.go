package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/skharv/tilegame/tilemap"
)

type Bindings struct {
	up, down, left, right ebiten.Key
}

type Controller struct {
	bindings              *Bindings
	player                *Player
	up, down, left, right bool
}

func (c *Controller) Init(player *Player) {
	c.bindings = &Bindings{up: ebiten.KeyW, down: ebiten.KeyS, left: ebiten.KeyA, right: ebiten.KeyD}
	c.player = player
}

func (c *Controller) ReadInputs() {
	c.up = inpututil.IsKeyJustPressed(c.bindings.up)
	c.down = inpututil.IsKeyJustPressed(c.bindings.down)
	c.left = inpututil.IsKeyJustPressed(c.bindings.left)
	c.right = inpututil.IsKeyJustPressed(c.bindings.right)
}

func (c *Controller) Update(tileMap *tilemap.TileMap) error {
	if c.up {
		c.player.MoveEntitiesUp(tileMap)
	}

	return nil
}
