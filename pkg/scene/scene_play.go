package scene

import (
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/pkg/sound"
	"github.com/masa213f/stg/pkg/stage"
)

type playSceneHandler struct {
	stgHandler *stage.Handler
}

func newPlayScene() handler {
	h := &playSceneHandler{}
	h.init()
	return h
}

func (h *playSceneHandler) init() {
	h.stgHandler = stage.NewHandler()
}

func (h *playSceneHandler) reset() {
	h.stgHandler.Init()
}

func (h *playSceneHandler) update() event {
	if input.Pause() {
		sound.BGM.Pause()
		return gameEventPause
	}

	sound.BGM.Play()
	result := h.stgHandler.Update()
	switch result {
	case stage.GameOver:
		return gameEventGameOver
	case stage.StageClear:
		return gameEventStageClear
	case stage.Playing:
		return eventNone
	default:
		panic("invalid result")
	}
}

func (h *playSceneHandler) draw() {
	h.stgHandler.Draw()
}
