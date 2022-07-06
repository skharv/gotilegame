package entities

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Init()
	Update() error
	Draw(screen *ebiten.Image)
	GetDrawLayer() int
}

type Manager struct {
	ents []Entity
}

func (m *Manager) Init() {
	for _, e := range m.ents {
		e.Init()
	}
}

func (m *Manager) Update() {
	sort.Slice(m.ents, func(i, j int) bool { return m.ents[i].GetDrawLayer() < m.ents[j].GetDrawLayer() })

	for _, e := range m.ents {
		e.Update()
	}
}

func (m *Manager) Draw(screen *ebiten.Image) {
	for _, e := range m.ents {
		e.Draw(screen)
	}
}

func (m *Manager) Register(e Entity) {
	if e != nil {
		m.ents = append(m.ents, e)
	}
}
