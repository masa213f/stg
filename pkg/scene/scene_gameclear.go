package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/resource"
)

type stageClearSceneHandler struct {
}

func NewStageClear() Handler {
	return &stageClearSceneHandler{}
}

func (h *stageClearSceneHandler) Reset() {
	// nothing
}

func (h *stageClearSceneHandler) Update() Event {
	if input.OK() {
		return EventNext
	}
	return EventNone
}

func (h *stageClearSceneHandler) Draw() {
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.Text(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle, "Clear!")
}
