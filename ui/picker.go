package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/resources"
)

type Icon struct {
	sprite *ebiten.Image
	title  string
}

type Picker struct {
	icons           []*Icon
	selectionSprite *ebiten.Image
	visible         bool
}

func (p *Picker) Init() {
	p.icons = append(p.icons, &Icon{sprite: resources.LoadFileAsImage("images/blueuniticon.png"), title: "BlueUnit"})
	p.icons = append(p.icons, &Icon{sprite: resources.LoadFileAsImage("images/reduniticon.png"), title: "RedUnit"})
}

func (p *Picker) Update() error {
	return nil
}

func (p *Picker) Draw(screen *ebiten.Image) {
	for _, i := range p.icons {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(i.sprite, op)
	}

}
