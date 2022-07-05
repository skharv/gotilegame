package tilemap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/entities"
	"github.com/skharv/tilegame/geom"
)

type Tile struct {
	mapIndex  geom.Vector2[int]
	worldPos  geom.Vector2[float64]
	originPos geom.Vector2[float64]
	object    *entities.Object
	sprite    *ebiten.Image
	color     *ebiten.ColorM
}

const (
	tileSizeX = 64
	tileSizeY = 64
)

func (t *Tile) Init() {

}

func (t *Tile) GetDrawLayer() int {
	layer := int(t.worldPos.Y)
	return layer
}

func (t *Tile) Update() error {
	return nil
}

func (t *Tile) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(t.worldPos.X-t.originPos.X, t.worldPos.Y-t.originPos.Y)
	options.ColorM = *t.color
	screen.DrawImage(t.sprite, options)
}

func (t *Tile) GetRect() (X, Y, W, H float64) {
	return t.worldPos.X - t.originPos.X, t.worldPos.Y - t.originPos.Y, tileSizeX, tileSizeY
}

func (t *Tile) IsVecInRect(PosX, PosY int) bool {
	X, Y, W, H := t.GetRect()
	if PosX > int(X) && PosX < int(X+W) {
		if PosY > int(Y) && PosY < int(Y+H) {
			return true
		}
	}
	return false
}

func (t *Tile) GetPosition() (float64, float64) {
	return t.worldPos.X, t.worldPos.Y
}

func (t *Tile) IsOccupied() bool {
	return t.object != nil
}

func (t *Tile) SetObject(obj *entities.Object) {
	t.object = obj
}

func (t *Tile) GetObject() *entities.Object {
	return t.object
}
