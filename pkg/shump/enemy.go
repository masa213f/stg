package shooting

import (
	"math/rand"

	"github.com/masa213f/shootinggame/pkg/draw"
	"github.com/masa213f/shootinggame/pkg/shape"
	"github.com/masa213f/shootinggame/resource"
)

type enemy struct {
	tick        int
	phase       int
	disabled    bool
	untouchable bool
	life        int
	hitRect     *shape.Rect
	drawRect    *shape.Rect
}

func newEnemy(x, y int) *enemy {
	return &enemy{
		life:     rand.Intn(3) + 1,
		hitRect:  shape.NewRect(x-8, y-8, 16, 16),
		drawRect: shape.NewRect(x-16, y-16, 32, 32),
	}
}

func (e *enemy) damage(d int) (score int) {
	e.life -= d
	if e.life > 0 {
		return 0
	}

	e.phase = 1
	e.untouchable = true
	e.tick = 0
	return 1
}

func (e *enemy) update() {
	e.tick++
	if e.phase == 0 {
		// 通常移動
		var v *shape.Vector
		if (e.tick>>6)%2 == 0 {
			v = shape.NewVector(-1, 1)
		} else {
			v = shape.NewVector(-1, -1)
		}
		e.hitRect.Move(v)
		e.drawRect.Move(v)
		if e.drawRect.X1() <= 0 {
			// 画面外に出た場合、敵は消滅
			e.disabled = true
			e.untouchable = true
		}
	} else if e.phase == 1 {
		// やられた場合
		if e.tick >= 16 {
			e.disabled = true
			e.untouchable = true
		}
		v := shape.NewVector(0, -2)
		e.hitRect.Move(v)
		e.drawRect.Move(v)
	}
}

func (e *enemy) draw() {
	if e.phase == 0 {
		draw.ImageAt(resource.ImageObake[(e.tick>>5)%4], e.drawRect.X0(), e.drawRect.Y0())
		// debugLineX(e.hitRect, color.Black)
	} else if e.phase == 1 {
		draw.ImageAt(resource.ImageEffectIce[(e.tick/2)], e.drawRect.X0(), e.drawRect.Y0())
	}
}
