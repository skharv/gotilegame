package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skharv/tilegame/globals"
)

var (
	transitionFrom = ebiten.NewImage(globals.ScreenWidth, globals.ScreenHeight)
	transitionTo   = ebiten.NewImage(globals.ScreenWidth, globals.ScreenHeight)
)

type Scene interface {
	Init()
	ReadInput()
	Update(state *GameState, deltaTime float64) error
	Draw(screen *ebiten.Image)
}

type Manager struct {
	current            Scene
	next               Scene
	transitionCount    float64
	transitionMaxCount float64
}

type GameState struct {
	SceneManager *Manager
}

func (m *Manager) ReadInput() {
	if m.transitionCount <= 0 {
		m.current.ReadInput()
		return
	}
}

func (m *Manager) Update(deltaTime float64) error {
	if m.transitionCount <= 0 {
		return m.current.Update(&GameState{
			SceneManager: m,
		}, deltaTime)
	}

	m.transitionCount -= 1 * deltaTime
	if m.transitionCount > 0 {
		return nil
	}

	m.current = m.next
	m.next = nil
	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {
	if m.transitionCount <= 0 {
		m.current.Draw(screen)
		return
	}

	transitionFrom.Clear()
	m.current.Draw(transitionFrom)

	transitionTo.Clear()
	m.next.Draw(transitionTo)

	screen.DrawImage(transitionFrom, nil)

	alpha := 1 - float64(m.transitionCount)/float64(m.transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawImage(transitionTo, op)
}

func (m *Manager) GoTo(scene Scene, fadeTime float64) {
	scene.Init()

	if m.current == nil {
		m.current = scene
	} else {
		m.next = scene
		m.transitionCount = fadeTime
		m.transitionMaxCount = fadeTime
	}
}
