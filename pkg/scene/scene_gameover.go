package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

type gameOverSceneHandler struct {
	screen util.Screen
	ctrl   util.Control
}

func NewGameOver(screen util.Screen, ctrl util.Control) Handler {
	return &gameOverSceneHandler{ctrl: ctrl}
}

func (h *gameOverSceneHandler) Reset() {
	// nothing
}

func (h *gameOverSceneHandler) Update() Event {
	if h.ctrl.Select() {
		return EventNext
	}
	return EventNone
}

func (h *gameOverSceneHandler) Draw() {
	h.screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	h.screen.Text(resource.FontArcade, color.White, util.HorizontalAlignCenter, util.VerticalAlignMiddle, "GameOver")
}
