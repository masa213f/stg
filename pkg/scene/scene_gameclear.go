package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

type stageClearSceneHandler struct {
	screen util.Screen
	ctrl   util.Control
}

func NewStageClear(screen util.Screen, ctrl util.Control) Handler {
	return &stageClearSceneHandler{
		screen: screen,
		ctrl:   ctrl,
	}
}

func (h *stageClearSceneHandler) Reset() {
	// nothing
}

func (h *stageClearSceneHandler) Update() Event {
	if h.ctrl.Select() {
		return EventNext
	}
	return EventNone
}

func (h *stageClearSceneHandler) Draw() {
	h.screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	h.screen.Text(resource.FontArcade, color.White, util.HorizontalAlignCenter, util.VerticalAlignMiddle, "Clear!")
}
