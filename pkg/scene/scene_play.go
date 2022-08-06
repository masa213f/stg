package scene

import (
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/pkg/sound"
	"github.com/masa213f/stg/pkg/stage"
)

type playSceneHandler struct {
	stgHandler *stage.Handler
}

func NewPlay() Handler {
	h := &playSceneHandler{}
	h.init()
	return h
}

func (h *playSceneHandler) init() {
	h.stgHandler = stage.NewHandler()
}

func (h *playSceneHandler) Reset() {
	h.stgHandler.Init()
}

func (h *playSceneHandler) Update() Event {
	if input.Pause() {
		sound.BGM.Pause()
		return GameEventPause
	}

	sound.BGM.Play()
	result := h.stgHandler.Update()
	switch result {
	case stage.GameOver:
		return GameEventGameOver
	case stage.StageClear:
		return GameEventStageClear
	case stage.Playing:
		return EventNone
	default:
		panic("invalid result")
	}
}

func (h *playSceneHandler) Draw() {
	h.stgHandler.Draw()
}
