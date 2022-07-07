package tilemap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/entities"
	"github.com/skharv/tilegame/geom"
	"github.com/skharv/tilegame/globals"
)

type TileState int

type Tile struct {
	mapIndex     geom.Vector2[int]
	worldPos     geom.Vector2[float64]
	originPos    geom.Vector2[float64]
	state        TileState
	object       *entities.Object
	sprite       *ebiten.Image
	color        *ebiten.ColorM
	justOccupied bool
}

const (
	tileSizeX    = 64
	tileSizeY    = 64
	snapDistance = 10

	Unoccupied TileState = 0
	Occupied   TileState = 1
	Awaiting   TileState = 2
)

func (t *Tile) Init() {

}

func (t *Tile) GetDrawLayer() int {
	layer := int(t.worldPos.Y)
	return layer
}

func (t *Tile) Update() error {
	t.justOccupied = false
	if t.object != nil {
		if t.worldPos.DistanceTo(t.object.GetWorldPos()) < snapDistance {
			if t.state == Awaiting {
				t.justOccupied = true
			}
			t.state = Occupied
		} else {
			t.state = Awaiting
		}
	} else {
		t.state = Unoccupied
	}

	return nil
}

func (t *Tile) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(t.worldPos.X-t.originPos.X, t.worldPos.Y-t.originPos.Y)
	options.ColorM = *t.color
	if globals.Debug {
		switch t.state {
		case Unoccupied:
			options.ColorM.Scale(1, 0, 0, 1)
		case Occupied:
			options.ColorM.Scale(0, 1, 0, 1)
		case Awaiting:
			options.ColorM.Scale(0, 0, 1, 1)
		}
	}
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
	return t.state == Occupied
}

func (t *Tile) JustOccupied() bool {
	return t.justOccupied
}

func (t *Tile) GetState() TileState {
	return t.state
}

func (t *Tile) SetObject(obj *entities.Object) {
	t.object = obj
	if t.object == nil {
		t.state = Unoccupied
	} else {
		t.state = Awaiting
	}
}

func (t *Tile) GetObject() *entities.Object {
	return t.object
}
