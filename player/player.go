package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/entities"
	"github.com/skharv/tilegame/tilemap"
	"github.com/skharv/tilegame/ui"
)

type Player struct {
	units      []*entities.Object
	cursor     *ui.Cursor
	picker     *ui.Picker
	controller *Controller
}

func (p *Player) Init() {
	p.cursor = &ui.Cursor{}
	p.cursor.Init()

	p.picker = &ui.Picker{}
	p.picker.Init()

	p.controller = &Controller{}
	p.controller.Init(p)
}

func (p *Player) ReadInputs() {
	p.controller.ReadInputs()
	p.cursor.ReadInputs()
	p.picker.ReadInputs()
}

func (p *Player) Update(tileMap *tilemap.TileMap, entityManager *entities.Manager) {
	p.cursor.Update(tileMap)
	p.picker.Update()
	p.controller.Update(tileMap)

	if p.cursor.IsAltClicked() && p.cursor.IsVisible() {
		i, j := p.cursor.GetPosition()
		if p.picker.GetSelectedAction() == "RedUnit" {
			obj, register := p.CreateObject(0, tileMap.WorldToTile(int(i), int(j)))
			if register {
				entityManager.Register(obj)
			}
		}
		if p.picker.GetSelectedAction() == "BlueUnit" {
			obj, register := p.CreateObject(1, tileMap.WorldToTile(int(i), int(j)))
			if register {
				entityManager.Register(obj)
			}
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.cursor.Draw(screen)
	p.picker.Draw(screen)
}

func (p *Player) CreateObject(id int, tile *tilemap.Tile) (*entities.Object, bool) {
	if tile.IsOccupied() {
		return nil, false
	}

	obj := &entities.Object{}
	obj.Init()
	obj.SetPolarity(id)
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

		if !newTile.IsOccupied() {
			currentTile.SetObject(nil)
			newTile.SetObject(unit)

			unit.SetPosition(newTile.GetPosition())
		}
	}
}

func (p *Player) MoveEntitiesDown(tileMap *tilemap.TileMap) {
	for _, unit := range p.units {

		currentTile := tileMap.ObjectToTile(unit)
		newTile := tileMap.GetSouthOf(currentTile)

		if !newTile.IsOccupied() {
			currentTile.SetObject(nil)
			newTile.SetObject(unit)

			unit.SetPosition(newTile.GetPosition())
		}
	}
}

func (p *Player) MoveEntitiesLeft(tileMap *tilemap.TileMap) {
	for _, unit := range p.units {

		currentTile := tileMap.ObjectToTile(unit)
		newTile := tileMap.GetWestOf(currentTile)

		if !newTile.IsOccupied() {
			currentTile.SetObject(nil)
			newTile.SetObject(unit)

			unit.SetPosition(newTile.GetPosition())
		}
	}
}

func (p *Player) MoveEntitiesRight(tileMap *tilemap.TileMap) {
	for _, unit := range p.units {

		currentTile := tileMap.ObjectToTile(unit)
		newTile := tileMap.GetEastOf(currentTile)

		if !newTile.IsOccupied() {
			currentTile.SetObject(nil)
			newTile.SetObject(unit)

			unit.SetPosition(newTile.GetPosition())
		}
	}
}
