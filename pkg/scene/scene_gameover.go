package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

type gameOverSceneHandler struct {
	ctrl util.Control
}

func NewGameOver(ctrl util.Control) Handler {
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
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.Text(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle, "GameOver")
}
