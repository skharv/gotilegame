package entities

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/data"
	"github.com/skharv/tilegame/geom"
	"github.com/skharv/tilegame/resources"
)

type ObjectState int

type Object struct {
	targetPos geom.Vector2[float64]
	worldPos  geom.Vector2[float64]
	originPos geom.Vector2[float64]
	sprite    *ebiten.Image
	color     *ebiten.ColorM
	data      *data.Unit
	state     ObjectState
	polarity  int
	shake     bool
}

const (
	shakeMagnitude = 3
	step           = 0.1
	snapDistance   = 10

	Idle      ObjectState = 0
	InTransit ObjectState = 1
	Dying     ObjectState = 2
)

func (o *Object) Init() {
	o.sprite = resources.LoadFileAsImage("images/redunit.png")
	o.originPos.X = 32
	o.originPos.Y = 96
	o.color = &ebiten.ColorM{}
	o.polarity = 0
	o.shake = false
	o.state = Idle
}

func (o *Object) GetDrawLayer() int {
	layer := int(o.worldPos.Y)
	return layer
}

func (o *Object) Update() error {
	if o.state != Dying {
		if o.worldPos.DistanceTo(o.targetPos) < snapDistance {
			o.worldPos = o.targetPos
			o.state = Idle
			o.shake = false
		} else {
			o.worldPos.Lerp(o.targetPos, step)
			o.state = InTransit
			o.shake = true
		}
	}

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

	o.targetPos.X = X
	o.targetPos.Y = Y
}

func (o *Object) SetTargetPos(X, Y float64) {
	o.targetPos.X = X
	o.targetPos.Y = Y
	o.state = InTransit
}

func (o *Object) SetSprite(imageFilepath string) {
	o.sprite = resources.LoadFileAsImage(imageFilepath)
}

func (o *Object) SetPolarity(i int) {
	o.data.Polarity = i
}

func (o *Object) SetState(state ObjectState) {
	o.state = state
}

func (o *Object) SetData(name string, hitpoints, attack, polarity int) {
	o.data = &data.Unit{Name: name, HitPoints: hitpoints, Attack: attack, Polarity: polarity}
}

func (o *Object) GetAttackDamage() int {
	return o.data.Attack
}

func (o *Object) TakeDamage(attackDamage int) {
	o.data.HitPoints -= attackDamage
	if o.data.HitPoints <= 0 {
		o.state = Dying
	}
}

func (o *Object) GetPolarity() int {
	return o.data.Polarity
}

func (o *Object) GetState() ObjectState {
	return o.state
}

func (o *Object) GetWorldPos() geom.Vector2[float64] {
	return o.worldPos
}

func (o *Object) IsAlive() bool {
	return o.state != Dying
}
