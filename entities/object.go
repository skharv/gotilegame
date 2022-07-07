package entities

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/geom"
	"github.com/skharv/tilegame/resources"
)

type Object struct {
	worldPos  geom.Vector2[float64]
	originPos geom.Vector2[float64]
	sprite    *ebiten.Image
	color     *ebiten.ColorM
	polarity  int
	shake     bool
}

const (
	shakeMagnitude = 3
)

func (o *Object) Init() {
	o.sprite = resources.LoadFileAsImage("images/redunit.png")
	o.originPos.X = 32
	o.originPos.Y = 96
	o.color = &ebiten.ColorM{}
	o.polarity = 0
	o.shake = true
}

func (o *Object) GetDrawLayer() int {
	layer := int(o.worldPos.Y)
	return layer
}

func (o *Object) Update() error {
	return nil
}

func (o *Object) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(o.worldPos.X-o.originPos.X, o.worldPos.Y-o.originPos.Y)
	if o.shake {
		x := float64(rand.Intn(shakeMagnitude*2) - shakeMagnitude)
		y := float64(rand.Intn(shakeMagnitude*2) - shakeMagnitude)
		options.GeoM.Translate(x, y)
	}
	options.ColorM = *o.color
	screen.DrawImage(o.sprite, options)
}

func (o *Object) SetPosition(X, Y float64) {
	o.worldPos.X = X
	o.worldPos.Y = Y
}

func (o *Object) SetSprite(imageFilepath string) {
	o.sprite = resources.LoadFileAsImage(imageFilepath)
}

func (o *Object) SetPolarity(i int) {
	o.polarity = i
}

func (o *Object) GetPolarity() int {
	return o.polarity
}
