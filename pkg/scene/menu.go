package scene

import (
	"image/color"

	"github.com/masa213f/shootinggame/pkg/draw"
	"github.com/masa213f/shootinggame/pkg/input"
	"github.com/masa213f/shootinggame/resource"
)

type item struct {
	text  string
	value id
}

type itemSelector struct {
	cursor int
	texts  []string
	values []id
}

func newItemSelector(items []item) *itemSelector {
	num := len(items)
	selector := &itemSelector{
		texts:  make([]string, num),
		values: make([]id, num),
	}
	for i, it := range items {
		selector.texts[i] = it.text
		selector.values[i] = it.value
	}
	return selector
}

func (i *itemSelector) getIndex() int {
	return i.cursor
}

func (i *itemSelector) getValue() id {
	return i.values[i.cursor]
}

func (i *itemSelector) getTexts() []string {
	return i.texts
}

func (i *itemSelector) next() {
	i.cursor = (i.cursor + 1) % len(i.values)
}

func (i *itemSelector) priv() {
	i.cursor = (i.cursor - 1 + len(i.values)) % len(i.values)
}

func (i *itemSelector) first() {
	i.cursor = 0
}

func (i *itemSelector) last() {
	i.cursor = len(i.values) - 1
}

type menuSceneHandler struct {
	items *itemSelector
}

func newMenuScene() handler {
	h := &menuSceneHandler{
		items: newItemSelector([]item{
			{"Play", scenePlay},
			// {"Options", sceneConfig},
			{"Exit", sceneExit},
		}),
	}
	return h
}

func (h *menuSceneHandler) update(priv id) id {
	if priv != sceneMenu {
		resource.BGM.Reset(resource.BGMMenu)
		h.items.first()
	}

	if input.OK() {
		return h.items.getValue()
	}
	if input.Cancel() {
		h.items.last()
		return sceneMenu
	}
	switch input.MenuUpOrDown() {
	case input.MoveUp:
		h.items.priv()
	case input.MoveDown:
		h.items.next()
	}
	return sceneMenu
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
