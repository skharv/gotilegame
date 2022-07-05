package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/geom"
	"github.com/skharv/tilegame/resources"
	"github.com/skharv/tilegame/tilemap"
)

type Cursor struct {
	sprite     *ebiten.Image
	worldPos   geom.Vector2[float64]
	cursorPos  geom.Vector2[int]
	originPos  geom.Vector2[float64]
	visible    bool
	clicked    bool
	altClicked bool
}

func (c *Cursor) Init() {
	c.sprite = resources.LoadFileAsImage("images/cursor.png")
	c.originPos.X = 32
	c.originPos.Y = 32
}

func (c *Cursor) ReadInputs() {
	c.cursorPos.X, c.cursorPos.Y = ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		c.clicked = true
	} else {
		c.clicked = false
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		c.altClicked = true
	} else {
		c.altClicked = false
	}
}

func (c *Cursor) Update(tilemap *tilemap.TileMap) error {
	X, Y, found := tilemap.WorldToTileWorldPos(c.cursorPos.X, c.cursorPos.Y)

	c.visible = found

	if c.visible {
		c.worldPos.X = float64(X)
		c.worldPos.Y = float64(Y)
	}

	return nil
}

func (c *Cursor) Draw(screen *ebiten.Image) {
	if c.visible {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(c.worldPos.X-c.originPos.X, c.worldPos.Y-c.originPos.Y)
		screen.DrawImage(c.sprite, options)
	}
}

func (c *Cursor) IsClicked() bool {
	return c.clicked
}

func (c *Cursor) IsAltClicked() bool {
	return c.altClicked
}

func (c *Cursor) IsVisible() bool {
	return c.visible
}

func (c *Cursor) GetPosition() (float64, float64) {
	return c.worldPos.X, c.worldPos.Y
}
