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
	IsAlive() bool
}

type Manager struct {
	ents  []Entity
	dying []Entity
}

func (m *Manager) Init() {
	for _, e := range m.ents {
		e.Init()
	}
}

func (m *Manager) Update() {
	for _, e := range m.ents {
		e.Update()
		if !e.IsAlive() {
			m.dying = append(m.dying, e)
		}
	}

	for _, d := range m.dying {
		if len(m.ents) <= 1 {
			m.ents = m.ents[:len(m.ents)-1]
		} else {
			length := len(m.ents)
			lastIndex := length - 1
			for i, e := range m.ents {
				if e == d {
					if i != lastIndex {
						m.ents[i] = m.ents[lastIndex]
					}
					m.ents = m.ents[:lastIndex]
					break
				}
			}
		}
	}

	sort.Slice(m.ents, func(i, j int) bool { return m.ents[i].GetDrawLayer() < m.ents[j].GetDrawLayer() })
}

func (m *Manager) RemoveDead() {
	if len(m.dying) > 0 {
		newDeadList := &[]Entity{}
		m.dying = *newDeadList
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
