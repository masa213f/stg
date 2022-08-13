package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

type pauseSceneHandler struct {
	items *itemSelector
	ctrl  util.Control
}

func NewPause(ctrl util.Control) Handler {
	h := &pauseSceneHandler{
		items: newItemSelector([]item{
			{"continue", EventBack},
			{"retire", GameEventRetire},
		}),
		ctrl: ctrl,
	}
	return h
}

func (h *pauseSceneHandler) Reset() {
	h.items.first()
}

func (h *pauseSceneHandler) Update() Event {
	if h.ctrl.Pause() {
		return EventBack
	}
	if h.ctrl.Select() {
		return h.items.getValue()
	}
	if h.ctrl.Cancel() {
		h.items.last()
		return EventNone
	}
	switch h.ctrl.UpOrDown() {
	case util.MoveUp:
		h.items.priv()
	case util.MoveDown:
		h.items.next()
	}
	return EventNone
}

func (h *pauseSceneHandler) Draw() {
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
