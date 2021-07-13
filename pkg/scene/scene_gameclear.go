package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/resource"
)

type stageClearSceneHandler struct {
}

func newStageClearScene() handler {
	return &stageClearSceneHandler{}
}

func (h *stageClearSceneHandler) update(priv id) id {
	if input.OK() {
		return sceneMenu
	}
	return sceneStageClear
}

func (h *stageClearSceneHandler) draw() {
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.Text(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle, "Clear!")
}
