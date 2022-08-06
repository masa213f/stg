package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/resource"
)

type configSceneHandler struct {
}

func NewConfig() Handler {
	return &configSceneHandler{}
}

func (h *configSceneHandler) Reset() {
	// nothing
}

func (h *configSceneHandler) Update() Event {
	if input.Cancel() {
		return EventBack
	}
	return EventNone
}

func (h *configSceneHandler) Draw() {
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.Text(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle, "Config")
}
