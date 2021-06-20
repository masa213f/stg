package shooting

import (
	"image/color"
	"sort"

	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/debug"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/shape"
	"github.com/masa213f/stg/resource"
)

const playerShotMaxNum = 100

type PlayerShots interface {
	Update()
	Draw()
	NewShot(x, y, r, t int)
	MakeInactive(index int)
	GetHitRects() []*shape.Rect
}

type playerShotsImpl struct {
	inactiveIndex int
	store         []*pshot
}

type pshot struct {
	active   bool
	speed    *shape.Vector
	hitRect  *shape.Rect
	drawRect *shape.Rect
}

func newPlayerShots() PlayerShots {
	shots := &playerShotsImpl{
		store: make([]*pshot, playerShotMaxNum),
	}
	for i := 0; i < playerShotMaxNum; i++ {
		shots.store[i] = &pshot{
			active:   false,
			speed:    &shape.Vector{},
			hitRect:  &shape.Rect{},
			drawRect: &shape.Rect{},
		}
	}
	return shots
}

func (ps *playerShotsImpl) Update() {
	for _, ps := range ps.store {
		if !ps.active {
			continue
		}
		ps.hitRect.Move(ps.speed)
		ps.drawRect.Move(ps.speed)

		if ps.drawRect.X0() >= constant.ScreenWidth {
			ps.active = false
		}
	}

	sort.SliceStable(ps.store, func(i, j int) bool { return ps.store[i].active })
	for i, shot := range ps.store {
		if !shot.active {
			ps.inactiveIndex = i
			break
		}
	}
}

func (ps *playerShotsImpl) Draw() {
	for _, shot := range ps.store {
		if !shot.active {
			continue
		}
		draw.ImageAt(resource.ImageShot, shot.drawRect.X0(), shot.drawRect.Y0())
		debug.DrawLineX(shot.hitRect, color.White)
	}
}

func (ps *playerShotsImpl) NewShot(x, y, r, t int) {
	if ps.inactiveIndex == playerShotMaxNum {
		return
	}
	shot := ps.store[ps.inactiveIndex]
	shot.active = true
	shot.speed.ResetP(r, t)
	shot.hitRect.Reset(x-16, y-6, 32, 12)   // FIXME: remove magic numbers
	shot.drawRect.Reset(x-16, y-16, 32, 32) // FIXME: remove magic numbers
	ps.inactiveIndex++
}

func (ps *playerShotsImpl) MakeInactive(i int) {
	ps.store[i].active = false
}

// GetHitRects returns a slice of shots hit rects. This function should be called after Update().
func (ps *playerShotsImpl) GetHitRects() []*shape.Rect {
	ret := make([]*shape.Rect, ps.inactiveIndex)
	for i, shot := range ps.store {
		if !shot.active {
			continue
		}
		ret[i] = shot.hitRect
	}
	return ret
}
