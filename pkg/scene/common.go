package scene

type Event int

// Handler is a interface to define update and draw functions for each scene.
type Handler interface {
	Reset()
	Update() Event
	Draw()
}

const (
	EventNone Event = iota
	EventExit

	// general
	EventNext
	EventBack

	// menu
	MenuEventPlay
	MenuEventConfig

	// game
	GameEventPause
	GameEventRetire
	GameEventGameOver
	GameEventStageClear
)
