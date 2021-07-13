package scene

import "errors"

// id is the identifier for scenes.
type id int

const (
	sceneTitle id = iota
	sceneMenu
	sceneConfig
	scenePlay
	scenePause
	sceneGameOver
	sceneStageClear
	sceneExit
	numOfScene
)

type event int

const (
	eventNone event = iota
	eventExit

	// general
	eventNext
	eventBack

	// menu
	menuEventPlay
	menuEventConfig

	// game
	gameEventPause
	gameEventRetire
	gameEventGameOver
	gameEventStageClear
)

type transition map[event]struct {
	Next  id
	Reset bool
}

var transitionTable map[id]transition = map[id]transition{
	sceneTitle: {
		eventNext: {Next: sceneMenu, Reset: true},
	},
	sceneMenu: {
		menuEventPlay:   {Next: scenePlay, Reset: true},
		menuEventConfig: {Next: sceneConfig, Reset: true},
		eventExit:       {Next: sceneExit},
	},
	sceneConfig: {
		eventBack: {Next: sceneMenu, Reset: true},
	},
	scenePlay: {
		gameEventPause:      {Next: scenePause, Reset: true},
		gameEventGameOver:   {Next: sceneGameOver, Reset: true},
		gameEventStageClear: {Next: sceneStageClear, Reset: true},
	},
	scenePause: {
		eventBack:       {Next: scenePlay},
		gameEventRetire: {Next: sceneGameOver, Reset: true},
	},
	sceneGameOver: {
		eventNext: {Next: sceneMenu, Reset: true},
	},
	sceneStageClear: {
		eventNext: {Next: sceneMenu, Reset: true},
	},
	sceneExit: {},
}

// handler is a interface to define update and draw functions for each scene.
type handler interface {
	reset()
	update() event
	draw()
}

// Manager controls scene handlers.
type Manager struct {
	currentScene id
	event        event
	handlers     [numOfScene]handler
}

var ErrNormalTermination = errors.New("exit")

// NewManager creates a new Manager instance.
func NewManager() *Manager {
	return &Manager{
		currentScene: sceneTitle,
		event:        eventNone,
		handlers: [numOfScene]handler{
			sceneTitle:      newTitleScene(),
			sceneMenu:       newMenuScene(),
			sceneConfig:     newConfigScene(),
			scenePlay:       newPlayScene(),
			sceneGameOver:   newGameOverScene(),
			sceneStageClear: newStageClearScene(),
			scenePause:      newPauseScene(),
		},
	}
}

// Update executes a Update function of the current scene.
func (s *Manager) Update() error {
	if s.event == eventExit {
		return ErrNormalTermination
	}

	if s.event != eventNone {
		tr, ok := transitionTable[s.currentScene]
		if !ok {
			panic("invalid transition")
		}
		s.currentScene = tr[s.event].Next
		if tr[s.event].Reset {
			s.handlers[s.currentScene].reset()
		}
	}

	s.event = s.handlers[s.currentScene].update()
	return nil
}

// Draw executes a Draw function of the current scene.
func (s *Manager) Draw() {
	s.handlers[s.currentScene].draw()
}
