package scene

type Event int

// Handler is a interface to define update and draw functions for each scene.
type Handler interface {
	Reset()
	Update() Event
	Draw()
}

type dummySceneHandler struct{}

func NewDummy() Handler {
	return &dummySceneHandler{}
}

func (h *dummySceneHandler) Reset() {
	panic("dummy reset")
}

func (h *dummySceneHandler) Update() Event {
	panic("dummy update")
}

func (h *dummySceneHandler) Draw() {
	panic("dummy draw")
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
