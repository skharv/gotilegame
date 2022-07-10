package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/entities"
	"github.com/skharv/tilegame/tilemap"
	"github.com/skharv/tilegame/ui"
)

type Player struct {
	units      []*entities.Object
	objectUi   []*ui.Unit
	cursor     *ui.Cursor
	picker     *ui.Picker
	controller *Controller
	ready      bool
}

func (p *Player) Init() {
	p.cursor = &ui.Cursor{}
	p.cursor.Init()

	p.picker = &ui.Picker{}
	p.picker.Init()

	p.controller = &Controller{}
	p.controller.Init(p)

	p.ready = true
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
	for i, u := range p.objectUi {
		u.Update()
		if u.GetHitPoints() <= 0 {
			if len(p.objectUi) <= 1 {
				p.objectUi = p.objectUi[:len(p.objectUi)-1]
			} else {
				length := len(p.objectUi)
				lastIndex := length - 1
				if i != lastIndex {
					p.objectUi[i] = p.objectUi[lastIndex]
				}
				p.objectUi = p.objectUi[:lastIndex]
			}
		}
	}

	if p.cursor.IsAltClicked() && p.cursor.IsVisible() {
		i, j := p.cursor.GetPosition()
		if p.picker.GetSelectedAction() == "RedUnit" {
			obj, register := p.CreateObject(0, tileMap.WorldToTile(int(i), int(j)))
			if register {
				entityManager.Register(obj)
				p.ready = false
			}
		}
		if p.picker.GetSelectedAction() == "BlueUnit" {
			obj, register := p.CreateObject(1, tileMap.WorldToTile(int(i), int(j)))
			if register {
				entityManager.Register(obj)
				p.ready = false
			}
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.cursor.Draw(screen)
	p.picker.Draw(screen)
	for _, u := range p.objectUi {
		u.Draw(screen)
	}
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
		obj.SetData("red", 6, 3, id)
	case 1:
		obj.SetSprite("images/blueunit.png")
		obj.SetData("blue", 7, 3, id)
	}

	obj.SetPosition(tile.GetPosition())
	tile.SetObject(obj)

	p.units = append(p.units, obj)

	ui := &ui.Unit{}
	ui.Init(obj.GetData(), obj.GetWorldPos())
	p.objectUi = append(p.objectUi, ui)

	return obj, true
}

func (p *Player) GetLastPlayed() *entities.Object {
	return p.units[len(p.units)-1]
}

func (p *Player) IsReady() bool {
	return p.ready
}

func (p *Player) GetReady() {
	p.ready = true
}
