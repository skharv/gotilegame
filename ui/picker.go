package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/skharv/tilegame/geom"
	"github.com/skharv/tilegame/globals"
	"github.com/skharv/tilegame/resources"
)

type Icon struct {
	sprite   *ebiten.Image
	position geom.Vector2[float64]
	title    string
}

type Picker struct {
	icons           []*Icon
	position        geom.Vector2[float64]
	cursorPos       geom.Vector2[int]
	selecionPos     geom.Vector2[float64]
	selectionSprite *ebiten.Image
	visible         bool
	clicked         bool
	altClicked      bool
}

func (p *Picker) Init() {
	p.selectionSprite = resources.LoadFileAsImage("images/iconselection.png")
	p.icons = append(p.icons, &Icon{sprite: resources.LoadFileAsImage("images/blueuniticon.png"), title: "BlueUnit"})
	p.icons = append(p.icons, &Icon{sprite: resources.LoadFileAsImage("images/reduniticon.png"), title: "RedUnit"})

	p.position = geom.Vector2[float64]{X: globals.ScreenWidth / 2, Y: globals.ScreenHeight - (float64(p.icons[0].sprite.Bounds().Dy()) * 1.5)}

	for i, v := range p.icons {
		v.position.Y = p.position.Y
		v.position.X = p.position.X + float64(v.sprite.Bounds().Dx()*i)
	}
}

func (p *Picker) ReadInputs() {
	p.cursorPos.X, p.cursorPos.Y = ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		p.clicked = true
	} else {
		p.clicked = false
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		p.altClicked = true
	} else {
		p.altClicked = false
	}
}

func (p *Picker) Update() error {
	if p.clicked {
		noneClicked := true
		for _, i := range p.icons {
			if p.IsVecInRect(i, p.cursorPos.X, p.cursorPos.Y) {
				p.visible = true
				p.selecionPos.X = i.position.X
				p.selecionPos.Y = i.position.Y
				noneClicked = false
			}
		}
		if noneClicked {
			p.visible = false
		}

	}
	return nil
}

func (p *Picker) Draw(screen *ebiten.Image) {
	for _, i := range p.icons {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(i.position.X, i.position.Y)
		screen.DrawImage(i.sprite, op)
	}

	if p.visible {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.selecionPos.X, p.selecionPos.Y)
		screen.DrawImage(p.selectionSprite, op)
	}
}

func (p *Picker) GetRect(icon *Icon) (X, Y, W, H float64) {
	return icon.position.X, icon.position.Y, float64(icon.sprite.Bounds().Dx()), float64(icon.sprite.Bounds().Dy())
}

func (p *Picker) IsVecInRect(icon *Icon, PosX, PosY int) bool {
	X, Y, W, H := p.GetRect(icon)
	if PosX > int(X) && PosX < int(X+W) {
		if PosY > int(Y) && PosY < int(Y+H) {
			return true
		}
	}
	return false
}
