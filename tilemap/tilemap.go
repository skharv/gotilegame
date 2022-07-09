package tilemap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/entities"
	"github.com/skharv/tilegame/geom"
	"github.com/skharv/tilegame/globals"
	"github.com/skharv/tilegame/resources"
)

type Direction int

type CombatPair struct {
	objectA *entities.Object
	objectB *entities.Object
}

type TileMap struct {
	position      geom.Vector2[float64]
	tiles         [mapSizeX][mapSizeY]*Tile
	pushActions   []*entities.Object
	combatActions []*CombatPair
	resolved      bool
}

const (
	mapSizeX = 10
	mapSizeY = 10
	tileOffset

	North Direction = 0
	South Direction = 1
	East  Direction = 2
	West  Direction = 3
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
				register:  t.AddToUpcoming,
			}
		}
	}
	t.resolved = true
}

func (t *TileMap) Update() {
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			t.tiles[i][j].Update()
		}
	}

	if !t.resolved {
		t.resolved = t.ResolveActions(t.pushActions...)
		t.resolved = t.ResolveCombat(t.combatActions...)
	}

	if t.AnyTileAwaiting() {
		t.resolved = false
	}
}

func (t *TileMap) Draw(screen *ebiten.Image) {
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			t.tiles[i][j].Draw(screen)
		}
	}
}

func (t *TileMap) AddToUpcoming(obj *entities.Object) {
	t.pushActions = append(t.pushActions, obj)
	t.resolved = false
}

func (t *TileMap) AnyTileAwaiting() bool {
	for i := 0; i < mapSizeX; i++ {
		for j := 0; j < mapSizeY; j++ {
			if t.tiles[i][j].GetState() == Awaiting {
				return true
			}
		}
	}
	return false
}

func (t *TileMap) SetAction(obj *entities.Object) {
	t.resolved = false
	t.pushActions = nil
	t.pushActions = append(t.pushActions, obj)
}

func (t *TileMap) RemoveAction(obj *entities.Object) {
	if len(t.pushActions) <= 1 {
		t.pushActions = t.pushActions[:len(t.pushActions)-1]
	} else {
		length := len(t.pushActions)
		lastIndex := length - 1

		for i, a := range t.pushActions {
			if a == obj {
				if i != lastIndex {
					t.pushActions[i] = t.pushActions[lastIndex]
				}
				t.pushActions = t.pushActions[:lastIndex]
				break
			}
		}
	}
}

func (t *TileMap) ResolveCombat(combat ...*CombatPair) bool {
	var nextCombat []*CombatPair

	for _, c := range combat {
		c.objectA.TakeDamage(c.objectB.GetAttackDamage())
		c.objectB.TakeDamage(c.objectA.GetAttackDamage())
		if c.objectA.GetState() != entities.Dying && c.objectB.GetState() != entities.Dying {
			nextCombat = append(nextCombat, c)
		}

		if c.objectA.GetState() == entities.Dying {
			t.ObjectToTile(c.objectA).SetObject(nil)
		}

		if c.objectB.GetState() == entities.Dying {
			t.ObjectToTile(c.objectB).SetObject(nil)
		}
	}

	t.combatActions = nextCombat

	return len(t.combatActions) == 0
}

func (t *TileMap) ResolveActions(actions ...*entities.Object) bool {
	var nextActions []*entities.Object

	for _, o := range actions {
		t.RemoveAction(o)

		var tile *Tile
		if o == nil {
			continue
		}

		if o.GetState() != entities.InTransit {
			tile = t.ObjectToTile(o)
		}

		if tile == nil {
			continue
		}

		nextActions = append(nextActions, t.ResolveTile(tile)...)
	}

	t.pushActions = append(t.pushActions, nextActions...)

	return len(t.pushActions) == 0
}

func (t *TileMap) ResolveTile(tile *Tile) []*entities.Object {
	var entities []*entities.Object
	north := t.GetNorthOf(tile)
	south := t.GetSouthOf(tile)
	east := t.GetEastOf(tile)
	west := t.GetWestOf(tile)

	if t.BothOccupied(tile, north) {
		if t.SamePolarity(tile, north) {
			entities = append(entities, t.MoveEntity(north, North))
		} else {
			t.CreateCombat(tile, north)
		}
	}
	if t.BothOccupied(tile, south) {
		if t.SamePolarity(tile, south) {
			entities = append(entities, t.MoveEntity(south, South))
		} else {
			t.CreateCombat(tile, south)
		}
	}
	if t.BothOccupied(tile, east) {
		if t.SamePolarity(tile, east) {
			entities = append(entities, t.MoveEntity(east, East))
		} else {
			t.CreateCombat(tile, east)
		}
	}
	if t.BothOccupied(tile, west) {
		if t.SamePolarity(tile, west) {
			entities = append(entities, t.MoveEntity(west, West))
		} else {
			t.CreateCombat(tile, west)
		}
	}

	return entities
}

func (t *TileMap) BothOccupied(tileA, tileB *Tile) bool {
	return tileA.IsOccupied() && tileB.IsOccupied()
}

func (t *TileMap) SamePolarity(tileA, tileB *Tile) bool {
	return tileA.GetObject().GetPolarity() == tileB.GetObject().GetPolarity()
}

func (t *TileMap) CreateCombat(tileA, tileB *Tile) {
	t.combatActions = append(t.combatActions, &CombatPair{objectA: tileA.GetObject(), objectB: tileB.GetObject()})
}

func (t *TileMap) MoveEntity(tile *Tile, direction Direction) *entities.Object {
	if tile.IsOccupied() {
		newTile := &Tile{}
		unit := tile.GetObject()

		switch direction {
		case North:
			newTile = t.GetNorthOf(tile)
		case South:
			newTile = t.GetSouthOf(tile)
		case East:
			newTile = t.GetEastOf(tile)
		case West:
			newTile = t.GetWestOf(tile)
		}

		if newTile.GetState() == Unoccupied {
			tile.SetObject(nil)
			newTile.SetObject(unit)
			unit.SetTargetPos(newTile.GetPosition())
			return unit
		}
	}
	return nil
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
	X := Clamp(tile.mapIndex.X, 0, mapSizeX-1)
	Y := Clamp(tile.mapIndex.Y-1, 0, mapSizeY-1)

	return t.tiles[X][Y]
}

func (t *TileMap) GetSouthOf(tile *Tile) *Tile {
	X := Clamp(tile.mapIndex.X, 0, mapSizeX-1)
	Y := Clamp(tile.mapIndex.Y+1, 0, mapSizeY-1)

	return t.tiles[X][Y]
}

func (t *TileMap) GetEastOf(tile *Tile) *Tile {
	X := Clamp(tile.mapIndex.X+1, 0, mapSizeX-1)
	Y := Clamp(tile.mapIndex.Y, 0, mapSizeY-1)

	return t.tiles[X][Y]
}

func (t *TileMap) GetWestOf(tile *Tile) *Tile {
	X := Clamp(tile.mapIndex.X-1, 0, mapSizeX-1)
	Y := Clamp(tile.mapIndex.Y, 0, mapSizeY-1)

	return t.tiles[X][Y]
}

func (t *TileMap) IsResolved() bool {
	return t.resolved
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
