package stage

import (
	"image/color"
	"math/rand"

	"github.com/masa213f/stg/pkg/debug"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/shape"
	"github.com/masa213f/stg/resource"
)

const (
	enemyStateNormal int = iota
	enemyStateDisappearing
)

type Enemy interface {
	Update()
	Draw()
	Damage(d int) (score int)
	IsDisabled() bool
	IsInvincible() bool
	GetHitRect() *shape.Rect
}

type enemyImpl struct {
	tick       int
	state      int
	disabled   bool
	invincible bool
	life       int
	hitRect    *shape.Rect
	drawRect   *shape.Rect
}

func newEnemy(x, y int) Enemy {
	return &enemyImpl{
		life:     rand.Intn(3) + 1,
		hitRect:  shape.NewRect(x-8, y-8, 16, 16),
		drawRect: shape.NewRect(x-16, y-16, 32, 32),
	}
}

func (e *enemyImpl) Update() {
	e.tick++
	switch e.state {
	case enemyStateNormal:
		var v *shape.Vector
		if (e.tick>>6)%2 == 0 {
			v = shape.NewVector(-1, 1)
		} else {
			v = shape.NewVector(-1, -1)
		}
		e.hitRect.Move(v)
		e.drawRect.Move(v)
		if e.drawRect.X1() <= 0 {
			// When going off the screen, the enemy disappears.
			e.disabled = true
			e.invincible = true
		}

	case enemyStateDisappearing:
		if e.tick >= 16 {
			e.disabled = true
			e.invincible = true
		}
		v := shape.NewVector(0, -2)
		e.hitRect.Move(v)
		e.drawRect.Move(v)
	}
}

func (e *enemyImpl) Draw() {
	switch e.state {
	case enemyStateNormal:
		draw.ImageAt(resource.ImageObake[(e.tick>>5)%4], e.drawRect.X0(), e.drawRect.Y0())
		debug.DrawLineX(e.hitRect, color.Black)
	case enemyStateDisappearing:
		draw.ImageAt(resource.ImageEffectIce[(e.tick/2)], e.drawRect.X0(), e.drawRect.Y0())
	}
}

func (e *enemyImpl) Damage(d int) (score int) {
	e.life -= d
	if e.life > 0 {
		return 0
	}

	e.state = enemyStateDisappearing
	e.invincible = true
	e.tick = 0
	return 1
}

func (e *enemyImpl) IsDisabled() bool {
	return e.disabled
}

func (e *enemyImpl) IsInvincible() bool {
	return e.invincible
}

func (e *enemyImpl) GetHitRect() *shape.Rect {
	return e.hitRect
}
