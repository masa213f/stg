package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

type configSceneHandler struct {
	ctrl util.Control
}

func NewConfig(ctrl util.Control) Handler {
	return &configSceneHandler{ctrl: ctrl}
}

func (h *configSceneHandler) Reset() {
	// nothing
}

func (h *configSceneHandler) Update() Event {
	if h.ctrl.Cancel() {
		return EventBack
	}
	return EventNone
}

func (h *configSceneHandler) Draw() {
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.Text(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle, "Config")
}
