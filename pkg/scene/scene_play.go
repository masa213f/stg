package scene

import (
	"github.com/masa213f/stg/pkg/stage"
	"github.com/masa213f/stg/pkg/util"
)

type playSceneHandler struct {
	stgHandler *stage.Handler
	ctrl       util.Control
	bgm        util.BGMPlayer
	se         util.SEPlayer
}

func NewPlay(ctrl util.Control, bgm util.BGMPlayer, se util.SEPlayer) Handler {
	h := &playSceneHandler{
		ctrl: ctrl,
		bgm:  bgm,
		se:   se,
	}
	h.init()
	return h
}

func (h *playSceneHandler) init() {
	h.stgHandler = stage.NewHandler(h.ctrl, h.bgm, h.se)
}

func (h *playSceneHandler) Reset() {
	h.stgHandler.Init()
}

func (h *playSceneHandler) Update() Event {
	if h.ctrl.Pause() {
		h.bgm.Pause()
		return GameEventPause
	}

	h.bgm.Play()
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
