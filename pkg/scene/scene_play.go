package scene

import (
	"github.com/masa213f/stg/pkg/stage"
	"github.com/masa213f/stg/pkg/util"
)

type playSceneHandler struct {
	stgHandler *stage.Handler
	ctrl       util.Control
	audio      util.AudioPlayer
}

func NewPlay(ctrl util.Control, audio util.AudioPlayer) Handler {
	h := &playSceneHandler{
		ctrl:  ctrl,
		audio: audio,
	}
	h.init()
	return h
}

func (h *playSceneHandler) init() {
	h.stgHandler = stage.NewHandler(h.ctrl, h.audio)
}

func (h *playSceneHandler) Reset() {
	h.stgHandler.Init()
}

func (h *playSceneHandler) Update() Event {
	if h.ctrl.Pause() {
		h.audio.PauseBGM()
		return GameEventPause
	}

	h.audio.PlayBGM()
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
