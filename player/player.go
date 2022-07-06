package player

import (
	"github.com/skharv/tilegame/entities"
	"github.com/skharv/tilegame/tilemap"
)

type Player struct {
	units []*entities.Object
}

func (p *Player) CreateObject(id int, tile *tilemap.Tile) (*entities.Object, bool) {
	if tile.IsOccupied() {
		return nil, false
	}

	obj := &entities.Object{}
	obj.Init()
	switch id {
	case 0:
		obj.SetSprite("images/redunit.png")
	case 1:
		obj.SetSprite("images/blueunit.png")
	}

	obj.SetPosition(tile.GetPosition())
	tile.SetObject(obj)

	p.units = append(p.units, obj)

	return obj, true
}

func (p *Player) MoveEntitiesUp(tileMap *tilemap.TileMap) {
	for _, unit := range p.units {
		currentTile := tileMap.ObjectToTile(unit)
		newTile := tileMap.GetNorthOf(currentTile)

		currentTile.SetObject(nil)
		newTile.SetObject(unit)

		unit.SetPosition(newTile.GetPosition())
	}
}
