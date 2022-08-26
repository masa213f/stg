package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

type configSceneHandler struct {
	screen util.Screen
	ctrl   util.Control
}

func NewConfig(screen util.Screen, ctrl util.Control) Handler {
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
	h.screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	h.screen.Text(resource.FontArcade, color.White, util.HorizontalAlignCenter, util.VerticalAlignMiddle, "Config")
}
