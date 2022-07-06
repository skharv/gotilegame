package tilemap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/entities"
	"github.com/skharv/tilegame/geom"
	"github.com/skharv/tilegame/globals"
	"github.com/skharv/tilegame/resources"
)

type TileMap struct {
	position geom.Vector2[float64]
	tiles    [mapSizeX][mapSizeY]*Tile
}

const (
	mapSizeX = 10
	mapSizeY = 10
	tileOffset
)

func (t *TileMap) Init() {
	t.position = geom.Vector2[float64]{X: globals.ScreenWidth/2 + tileSizeX/2, Y: globals.ScreenHeight/2 + tileSizeY/2}

	img := resources.LoadFileAsImage("images/tile.png")

	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			t.tiles[i][j] = &Tile{
				mapIndex:  geom.Vector2[int]{X: i, Y: j},
				worldPos:  geom.Vector2[float64]{X: float64(tileSizeX*(i-(mapSizeX/2))) + t.position.X, Y: float64(tileSizeY*(j-(mapSizeY/2))) + t.position.Y},
				originPos: geom.Vector2[float64]{X: tileSizeX / 2, Y: tileSizeY / 2},
				sprite:    img,
				color:     &ebiten.ColorM{},
			}
		}
	}
}

func (t *TileMap) Update() error {
	return nil
}

func (t *TileMap) Draw(screen *ebiten.Image) {
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			t.tiles[i][j].Draw(screen)
		}
	}
}

func (t *TileMap) WorldToTilePos(X, Y int) (int, int, bool) {
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			if t.tiles[i][j].IsVecInRect(X, Y) {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func (t *TileMap) WorldToTileWorldPos(X, Y int) (int, int, bool) {
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			if t.tiles[i][j].IsVecInRect(X, Y) {
				a, b := t.tiles[i][j].GetPosition()

				return int(a), int(b), true
			}
		}
	}
	return -1, -1, false
}

func (t *TileMap) WorldToTile(X, Y int) *Tile {
	i, j, _ := t.WorldToTilePos(X, Y)
	i = Clamp(i, 0, mapSizeX)
	j = Clamp(j, 0, mapSizeY)

	return t.tiles[i][j]
}

func (t *TileMap) MapIndexToTile(X, Y int) *Tile {
	X = Clamp(X, 0, mapSizeX)
	Y = Clamp(Y, 0, mapSizeY)

	return t.tiles[X][Y]
}

func (t *TileMap) ObjectToTile(obj *entities.Object) *Tile {
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			if t.tiles[i][j].GetObject() == obj {
				return t.tiles[i][j]
			}
		}
	}

	return nil
}

func (t *TileMap) IsTileOccupied(X, Y int) bool {
	X = Clamp(X, 0, mapSizeX)
	Y = Clamp(Y, 0, mapSizeY)

	return t.tiles[X][Y].IsOccupied()
}

func (t *TileMap) SetTileOccupant(obj *entities.Object) {
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			t.tiles[i][j].SetObject(obj)
		}
	}
}

func (t *TileMap) GetTileOccupant(X, Y int) (*entities.Object, bool) {
	X = Clamp(X, 0, mapSizeX)
	Y = Clamp(Y, 0, mapSizeY)

	return t.tiles[X][Y].GetObject(), true
}

func (t *TileMap) GetNorthOf(tile *Tile) *Tile {
	X := Clamp(tile.mapIndex.X, 0, mapSizeX)
	Y := Clamp(tile.mapIndex.Y-1, 0, mapSizeY)

	return t.tiles[X][Y]
}

func (t *TileMap) GetSouthOf(tile *Tile) *Tile {
	X := Clamp(tile.mapIndex.X, 0, mapSizeX)
	Y := Clamp(tile.mapIndex.Y+1, 0, mapSizeY)

	return t.tiles[X][Y]
}

func (t *TileMap) GetEastOf(tile *Tile) *Tile {
	X := Clamp(tile.mapIndex.X+1, 0, mapSizeX)
	Y := Clamp(tile.mapIndex.Y, 0, mapSizeY)

	return t.tiles[X][Y]
}

func (t *TileMap) GetWestOf(tile *Tile) *Tile {
	X := Clamp(tile.mapIndex.X-1, 0, mapSizeX)
	Y := Clamp(tile.mapIndex.Y, 0, mapSizeY)

	return t.tiles[X][Y]
}

//Universal Functions
func Clamp[T int | float64](Value, Min, Max T) T {
	value := Value
	if value < Min {
		value = Min
	} else if value > Max {
		value = Max
	}

	return value
}
