package shooting

import (
	"image/color"
	"sort"

	"github.com/masa213f/shootinggame/pkg/constant"
	"github.com/masa213f/shootinggame/pkg/debug"
	"github.com/masa213f/shootinggame/pkg/draw"
	"github.com/masa213f/shootinggame/pkg/shape"
	"github.com/masa213f/shootinggame/resource"
)

const playerShotMaxNum = 100

type playerShot struct {
	id       objectID
	speed    *shape.Vector
	hitRect  *shape.Rect
	drawRect *shape.Rect
}

func (ps *playerShot) update() {
	ps.hitRect.Move(ps.speed)
	ps.drawRect.Move(ps.speed)
}

type playerShotList struct {
	nextID    objectID
	activeNum int
	buffer    []*playerShot
}

func newPlayerShotList() *playerShotList {
	list := &playerShotList{
		buffer: make([]*playerShot, playerShotMaxNum),
	}
	for i := 0; i < playerShotMaxNum; i++ {
		list.buffer[i] = &playerShot{
			id:       inactiveObjectID,
			speed:    &shape.Vector{},
			hitRect:  &shape.Rect{},
			drawRect: &shape.Rect{},
		}
	}
	return list
}

func (list *playerShotList) new(x, y, r, t int) {
	if list.activeNum == playerShotMaxNum {
		return
	}
	ent := list.buffer[list.activeNum]
	ent.id = list.nextID
	ent.speed.ResetP(r, t)
	ent.hitRect.Reset(x-16, y-6, 32, 12)
	ent.drawRect.Reset(x-16, y-16, 32, 32)
	list.nextID++
	list.activeNum++
}

func (list *playerShotList) list() []*playerShot {
	return list.buffer
}

func (list *playerShotList) inactive(i int) {
	if list.buffer[i].id != inactiveObjectID {
		list.buffer[i].id = inactiveObjectID
		list.activeNum--
	}
}

func (list *playerShotList) gc() {
	for _, ent := range list.buffer {
		if ent.id == inactiveObjectID {
			break
		}
		if ent.drawRect.X0() >= constant.ScreenWidth {
			ent.id = inactiveObjectID
			list.activeNum--
		}
	}
	sort.Slice(list.buffer, func(i, j int) bool { return list.buffer[i].id < list.buffer[j].id })
}

func (list *playerShotList) drawAll() {
	for _, ent := range list.buffer {
		if ent.id == inactiveObjectID {
			break
		}
		draw.ImageAt(resource.ImageShot, ent.drawRect.X0(), ent.drawRect.Y0())
		debug.DrawLineX(ent.hitRect, color.White)
	}

}
