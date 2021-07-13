package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/resource"
)

type configSceneHandler struct {
}

func newConfigScene() handler {
	return &configSceneHandler{}
}

func (h *configSceneHandler) reset() {
	// nothing
}

func (h *configSceneHandler) update() event {
	if input.Cancel() {
		return eventBack
	}
	return eventNone
}

func (h *configSceneHandler) draw() {
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.Text(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle, "Config")
}
