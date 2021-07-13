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

func (h *playSceneHandler) update(priv id) id {
	if priv == sceneMenu {
		h.stgHandler.Init()
	}

	if input.Pause() {
		sound.BGM.Pause()
		return scenePause
	}

	sound.BGM.Play()
	ret := h.stgHandler.Update()
	if ret != nil {
		return sceneGameOver
	}
	return scenePlay
}

func (h *playSceneHandler) draw() {
	h.stgHandler.Draw()
}
