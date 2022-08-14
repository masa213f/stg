package manager

import (
	"errors"
	"fmt"

	"github.com/masa213f/stg/pkg/scene"
	"github.com/masa213f/stg/pkg/util"
)

// Manager controls scene handlers.
type Manager struct {
	currentScene    int
	event           scene.Event
	handlers        []scene.Handler
	transitionTable map[int]map[scene.Event]next
}

type next struct {
	scene     int
	keepState bool
}

var ErrExit = errors.New("exit")

// NewManager creates a new Manager instance.
func NewManager() *Manager {
	mgr := &Manager{
		event:           scene.EventNone,
		transitionTable: map[int]map[scene.Event]next{},
	}

	input := util.NewCombinedInput(util.NewKeyboardInput(), util.NewGamepadInput())
	ctrl := util.NewControl(input)

	sceneTitle := mgr.AddScene(scene.NewTitle(ctrl))
	sceneMenu := mgr.AddScene(scene.NewMenu(ctrl, audio))
	sceneConfig := mgr.AddScene(scene.NewConfig(ctrl))
	scenePlay := mgr.AddScene(scene.NewPlay(ctrl, audio))
	sceneGameOver := mgr.AddScene(scene.NewGameOver(ctrl))
	sceneStageClear := mgr.AddScene(scene.NewStageClear(ctrl))
	scenePause := mgr.AddScene(scene.NewPause(ctrl))

	transitions := []*struct {
		from      int
		to        int
		event     scene.Event
		keepState bool
	}{
		{from: sceneTitle, to: sceneMenu, event: scene.EventNext},
		{from: sceneMenu, to: scenePlay, event: scene.MenuEventPlay},
		{from: sceneMenu, to: sceneConfig, event: scene.MenuEventConfig},
		{from: sceneMenu, to: 0, event: scene.EventExit},
		{from: sceneConfig, to: sceneMenu, event: scene.EventBack},
		{from: scenePlay, to: scenePause, event: scene.GameEventPause},
		{from: scenePlay, to: sceneGameOver, event: scene.GameEventGameOver},
		{from: scenePlay, to: sceneStageClear, event: scene.GameEventStageClear},
		{from: scenePause, to: scenePlay, event: scene.EventBack, keepState: true},
		{from: scenePause, to: sceneGameOver, event: scene.GameEventRetire},
		{from: sceneGameOver, to: sceneMenu, event: scene.EventNext},
		{from: sceneStageClear, to: sceneMenu, event: scene.EventNext},
	}
	for _, tr := range transitions {
		mgr.AddTransition(tr.from, tr.to, tr.event, tr.keepState)
	}

	mgr.SetInitialScene(sceneTitle)
	return mgr
}

func (m *Manager) AddScene(h scene.Handler) int {
	m.handlers = append(m.handlers, h)
	id := len(m.handlers)
	m.transitionTable[id] = map[scene.Event]next{}
	return id
}

func (m *Manager) AddTransition(from, to int, e scene.Event, keepState bool) error {
	if m.transitionTable[from] == nil {
		return fmt.Errorf("scene %d is not registered", from)
	}
	m.transitionTable[from][e] = next{scene: to, keepState: keepState}
	return nil
}

func (m *Manager) SetInitialScene(s int) {
	m.currentScene = s
}

// Update executes a Update function of the current scene.
func (m *Manager) Update() error {
	if m.currentScene == 0 {
		return ErrExit
	}

	if m.event != scene.EventNone {
		tr, ok := m.transitionTable[m.currentScene]
		if !ok {
			panic("invalid transition")
		}
		m.currentScene = tr[m.event].scene
		if !tr[m.event].keepState {
			m.handlers[m.currentScene-1].Reset()
		}
	}

	m.event = m.handlers[m.currentScene-1].Update()
	return nil
}

// Draw executes a Draw function of the current scene.
func (m *Manager) Draw() {
	m.handlers[m.currentScene-1].Draw()
}
