package scene

import (
	"image/color"

	"github.com/masa213f/shootinggame/pkg/draw"
	"github.com/masa213f/shootinggame/pkg/input"
	"github.com/masa213f/shootinggame/resource"
)

type configSceneHandler struct {
}

func newConfigScene() handler {
	return &configSceneHandler{}
}

func (h *configSceneHandler) update(priv id) id {
	if input.Cancel() {
		return sceneMenu
	}
	return sceneConfig
}

func (h *configSceneHandler) draw() {
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.Text(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle, "Config")
}
