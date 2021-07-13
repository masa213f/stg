package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/resource"
)

type gameOverSceneHandler struct {
}

func newGameOverScene() handler {
	return &gameOverSceneHandler{}
}

func (h *gameOverSceneHandler) reset() {
	// nothing
}

func (h *gameOverSceneHandler) update() event {
	if input.OK() {
		return eventNext
	}
	return eventNone
}

func (h *gameOverSceneHandler) draw() {
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.Text(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle, "GameOver")
}
