package scene

import "errors"

// id is the identifier for scenes.
type id int

const (
	sceneNone id = iota
	sceneExit
	sceneTitle
	sceneMenu
	sceneConfig
	scenePlay
	scenePause
	sceneGameOver
	numOfScene
)

// handler is a interface to define update and draw functions for each scene.
type handler interface {
	update(prev id) id
	draw()
}

// Manager controls scene handlers.
type Manager struct {
	currentScene id
	nextScene    id
	handlers     [numOfScene]handler
}

var ErrNormalTermination = errors.New("exit")

// NewManager creates a new Manager instance.
func NewManager() *Manager {
	return &Manager{
		currentScene: sceneTitle,
		nextScene:    sceneTitle,
		handlers: [numOfScene]handler{
			sceneTitle:    newTitleScene(),
			sceneMenu:     newMenuScene(),
			sceneConfig:   newConfigScene(),
			scenePlay:     newPlayScene(),
			sceneGameOver: newGameOverScene(),
			scenePause:    newPauseScene(),
		},
	}
}

// Update executes a Update function of the current scene.
func (s *Manager) Update() error {
	priv := s.currentScene
	s.currentScene = s.nextScene
	next := s.handlers[s.currentScene].update(priv)
	if next == sceneExit {
		return ErrNormalTermination
	}
	s.nextScene = next
	return nil
}

// Draw executes a Draw function of the current scene.
func (s *Manager) Draw() {
	s.handlers[s.currentScene].draw()
}
