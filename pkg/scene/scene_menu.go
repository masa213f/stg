package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/pkg/sound"
	"github.com/masa213f/stg/resource"
)

func newMenuScene() handler {
	h := &menuSceneHandler{
		items: newItemSelector([]item{
			{"Play", menuEventPlay},
			// {"Options", sceneConfig},
			{"Exit", eventExit},
		}),
	}
	return h
}

func (h *menuSceneHandler) reset() {
	sound.BGM.Reset(resource.BGMMenu)
	h.items.first()
}

func (h *menuSceneHandler) update() event {
	if input.OK() {
		return h.items.getValue()
	}
	if input.Cancel() {
		h.items.last()
		return eventNone
	}
	switch input.UpOrDown() {
	case input.MoveUp:
		h.items.priv()
	case input.MoveDown:
		h.items.next()
	}
	return eventNone
}

func (h *menuSceneHandler) draw() {
	idx := h.items.getIndex()
	disp := []string{}
	for i, t := range h.items.getTexts() {
		if i == idx {
			disp = append(disp, "["+t+"]")
		} else {
			disp = append(disp, t)
		}
	}
	draw.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	draw.MultiText(resource.FontArcade, color.White, draw.HorizontalAlignCenter, draw.VerticalAlignMiddle, disp)
}
