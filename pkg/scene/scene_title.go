package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

type titleSceneHandler struct {
	ctrl util.Control
}

func NewTitle(ctrl util.Control) Handler {
	return &titleSceneHandler{ctrl: ctrl}
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
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.MultiText(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle,
		[]string{"Shooting", "press z key"})
	draw.Text(resource.FontArcadeSmall, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignBottom,
		"Copyright (c) 2021 masa213f")
}
