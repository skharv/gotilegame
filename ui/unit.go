package ui

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/data"
	"github.com/skharv/tilegame/geom"
	"github.com/skharv/tilegame/resources"
	"github.com/tinne26/etxt"
)

type Unit struct {
	name, hitpoints, attack string
	data                    *data.Unit
	worldPos                *geom.Vector2[float64]
	txtRenderer             *etxt.Renderer
}

func (u *Unit) Init(data *data.Unit, position *geom.Vector2[float64]) {
	u.txtRenderer = &etxt.Renderer{}

	fontLib := resources.LoadFileAsFont("fonts/JosefinSlab-Bold.ttf")
	u.txtRenderer = etxt.NewStdRenderer()
	glyphsCache := etxt.NewDefaultCache(10 * 1024 * 1024)
	u.txtRenderer.SetCacheHandler(glyphsCache.NewHandler())
	u.txtRenderer.SetFont(fontLib.GetFont("Josefin Slab Bold"))
	u.txtRenderer.SetAlign(etxt.YCenter, etxt.XCenter)
	u.txtRenderer.SetSizePx(24)
	u.txtRenderer.SetColor(color.Black)
	u.data = data
	u.worldPos = position
}

func (u *Unit) Update() error {
	u.name = u.data.Name
	u.hitpoints = strconv.Itoa(u.data.HitPoints)
	u.attack = strconv.Itoa(u.data.Attack)
	return nil
}

func (u *Unit) Draw(screen *ebiten.Image) {
	u.txtRenderer.SetTarget(screen)
	u.txtRenderer.Draw(u.name, int(u.worldPos.X), int(u.worldPos.Y-80))
	u.txtRenderer.Draw(u.attack, int(u.worldPos.X-24), int(u.worldPos.Y+16))
	u.txtRenderer.Draw(u.hitpoints, int(u.worldPos.X+24), int(u.worldPos.Y+16))
}

func (u *Unit) GetHitPoints() int {
	return u.data.HitPoints
}
