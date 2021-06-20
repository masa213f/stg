package shooting

import (
	"image/color"

	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/shape"
)

const BombDuration = 60 // 60 frame = 1 sec

type PlayerBomb interface {
	Update()
	Draw()
	NewBomb(x, y int)
	IsActive() bool
	GetHitRect() *shape.Rect
}

type playerBombImpl struct {
	duration int
	x        int
	y        int
	size     int
	hitRect  *shape.Rect
}

func newPlayerBomb() PlayerBomb {
	return &playerBombImpl{
		hitRect: &shape.Rect{},
	}
}

func (bomb *playerBombImpl) Update() {
	if bomb.IsActive() {
		bomb.duration--
		bomb.size += 3
		bomb.hitRect.Reset(bomb.x-32-bomb.size, bomb.y-32-bomb.size, 64+bomb.size*2, 64+bomb.size*2)
	}
}

func (bomb *playerBombImpl) Draw() {
	if bomb.IsActive() {
		// FIXME
		draw.Rect(bomb.hitRect, color.Black)
		draw.LineX(bomb.hitRect, color.Black)
	}
}

func (bomb *playerBombImpl) NewBomb(x, y int) {
	bomb.duration = BombDuration
	bomb.x = x
	bomb.y = y
	bomb.size = 0
	bomb.hitRect.Reset(x-32, y-32, 64, 64) // FIXME: remove magic number
}

func (bomb *playerBombImpl) IsActive() bool {
	return bomb.duration > 0
}

func (bomb *playerBombImpl) GetHitRect() *shape.Rect {
	return bomb.hitRect
}
