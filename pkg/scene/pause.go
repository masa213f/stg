package scene

import (
	"image/color"

	"github.com/masa213f/shootinggame/pkg/draw"
	"github.com/masa213f/shootinggame/pkg/input"
	"github.com/masa213f/shootinggame/resource"
)

type pauseSceneHandler struct {
	items *itemSelector
}

func newPauseScene() handler {
	h := &pauseSceneHandler{
		items: newItemSelector([]item{
			{"continue", scenePlay},
			{"menu", sceneMenu},
		}),
	}
	return h
}

func (h *pauseSceneHandler) update(priv id) id {
	if priv != scenePause {
		h.items.first()
	}

	if input.Pause() {
		return scenePlay
	}

	if input.OK() {
		return h.items.getValue()
	}
	if input.Cancel() {
		h.items.last()
		return sceneMenu
	}
	switch input.MenuUpOrDown() {
	case input.MoveUp:
		h.items.priv()
	case input.MoveDown:
		h.items.next()
	}
	return scenePause
}

func (h *pauseSceneHandler) draw() {
	idx := h.items.getIndex()
	disp := []string{"Pause", ""}
	for i, t := range h.items.getTexts() {
		if i == idx {
			disp = append(disp, "["+t+"]")
		} else {
			disp = append(disp, t)
		}
	}
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.MultiText(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle, disp)
}
