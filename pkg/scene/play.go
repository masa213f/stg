package scene

import (
	"github.com/masa213f/stg/pkg/input"
	shooting "github.com/masa213f/stg/pkg/shump"
	"github.com/masa213f/stg/pkg/sound"
)

type playSceneHandler struct {
	stgHandler *shooting.Handler
}

func newPlayScene() handler {
	h := &playSceneHandler{}
	h.init()
	return h
}

func (h *playSceneHandler) init() {
	h.stgHandler = shooting.NewHandler()
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
