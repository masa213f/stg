package scene

import (
	"github.com/masa213f/shootinggame/pkg/input"
	shooting "github.com/masa213f/shootinggame/pkg/shump"
	"github.com/masa213f/shootinggame/resource"
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
		resource.BGM.Pause()
		return scenePause
	}

	resource.BGM.Play()
	ret := h.stgHandler.Update()
	if ret != nil {
		return sceneGameOver
	}
	return scenePlay
}

func (h *playSceneHandler) draw() {
	h.stgHandler.Draw()
}
