package manager

import (
	"errors"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/masa213f/stg/pkg/scene"
	"github.com/masa213f/stg/pkg/util"
)

// Manager controls scene handlers.
type Manager struct {
	currentScene    int
	event           scene.Event
	screen          util.Screen
	handlers        []scene.Handler
	transitionTable map[int]map[scene.Event]next
}

type next struct {
	scene     int
	keepState bool
}

var ErrExit = errors.New("exit")

// NewManager creates a new Manager instance.
func NewManager(width, height int, debugMode bool) *Manager {
	screen := util.NewScreen(width, height, debugMode)
	mgr := &Manager{
		event:           scene.EventNone,
		screen:          screen,
		transitionTable: map[int]map[scene.Event]next{},
	}
	input := util.NewCombinedInput(util.NewKeyboardInput(), util.NewGamepadInput())
	ctrl := util.NewControl(input)

	sceneTitle := mgr.AddScene(scene.NewTitle(screen, ctrl))
	sceneMenu := mgr.AddScene(scene.NewMenu(screen, ctrl, audio))
	sceneConfig := mgr.AddScene(scene.NewConfig(screen, ctrl))
	scenePlay := mgr.AddScene(scene.NewPlay(screen, ctrl, audio))
	sceneGameOver := mgr.AddScene(scene.NewGameOver(screen, ctrl))
	sceneStageClear := mgr.AddScene(scene.NewStageClear(screen, ctrl))
	scenePause := mgr.AddScene(scene.NewPause(screen, ctrl))

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
func (m *Manager) Draw(screen *ebiten.Image) {
	m.screen.SetImage(screen)
	m.handlers[m.currentScene-1].Draw()
}
