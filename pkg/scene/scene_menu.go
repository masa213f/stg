package scene

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

type menuSceneHandler struct {
	items *itemSelector
	ctrl  util.Control
	bgm   util.BGMPlayer
}

func NewMenu(ctrl util.Control, bgm util.BGMPlayer) Handler {
	h := &menuSceneHandler{
		items: newItemSelector([]item{
			{"Play", MenuEventPlay},
			// {"Options", sceneConfig},
			{"Exit", EventExit},
		}),
		ctrl: ctrl,
		bgm:  bgm,
	}
	return h
}

func (h *menuSceneHandler) Reset() {
	h.bgm.Reset(resource.BGMMenu)
	h.items.first()
}

func (h *menuSceneHandler) Update() Event {
	if h.ctrl.Select() {
		return h.items.getValue()
	}
	if h.ctrl.Cancel() {
		h.items.last()
		return EventNone
	}
	switch h.ctrl.UpOrDown() {
	case util.MoveUp:
		h.items.priv()
	case util.MoveDown:
		h.items.next()
	}
	return EventNone
}

func (h *menuSceneHandler) Draw() {
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
