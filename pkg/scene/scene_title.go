package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/resource"
)

type titleSceneHandler struct {
}

func NewTitle() Handler {
	return &titleSceneHandler{}
}

func (h *titleSceneHandler) Reset() {
	// nothing
}

func (h *titleSceneHandler) Update() Event {
	if input.OK() {
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
