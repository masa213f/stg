package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

type titleSceneHandler struct {
	screen util.Screen
	ctrl   util.Control
}

func NewTitle(screen util.Screen, ctrl util.Control) Handler {
	return &titleSceneHandler{
		screen: screen,
		ctrl:   ctrl,
	}
}

func (h *titleSceneHandler) Reset() {
	// nothing
}

func (h *titleSceneHandler) Update() Event {
	if h.ctrl.Select() {
		return EventNext
	}
	return EventNone
}

func (h *titleSceneHandler) Draw() {
	h.screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	h.screen.MultiText(resource.FontArcade, color.White, util.HorizontalAlignCenter, util.VerticalAlignMiddle,
		[]string{"Shooting", "press z key"})
	h.screen.Text(resource.FontArcadeSmall, color.White, util.HorizontalAlignCenter, util.VerticalAlignBottom,
		"Copyright (c) 2021 masa213f")
}
