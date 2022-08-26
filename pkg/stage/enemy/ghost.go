package enemy

import (
	"image/color"
	"math/rand"

	"github.com/masa213f/stg/pkg/shape"
	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

const (
	enemyStateNormal int = iota
	enemyStateDisappearing
)

type ghost struct {
	tick       int
	state      int
	disabled   bool
	invincible bool
	life       int
	hitRect    *shape.Rect
	drawRect   *shape.Rect
}

func NewGhost(x, y int) Enemy {
	return &ghost{
		life:     rand.Intn(3) + 1,
		hitRect:  shape.NewRect(x-8, y-8, 16, 16),
		drawRect: shape.NewRect(x-16, y-16, 32, 32),
	}
}

func (g *ghost) Update() {
	g.tick++
	switch g.state {
	case enemyStateNormal:
		var v *shape.Vector
		if (g.tick>>6)%2 == 0 {
			v = shape.NewVector(-1, 1)
		} else {
			v = shape.NewVector(-1, -1)
		}
		g.hitRect.Move(v)
		g.drawRect.Move(v)
		if g.drawRect.X1() <= 0 {
			// When going off the screen, the enemy disappears.
			g.disabled = true
			g.invincible = true
		}

	case enemyStateDisappearing:
		if g.tick >= 16 {
			g.disabled = true
			g.invincible = true
		}
		v := shape.NewVector(0, -2)
		g.hitRect.Move(v)
		g.drawRect.Move(v)
	}
}

func (g *ghost) Draw(screen util.Screen) {
	switch g.state {
	case enemyStateNormal:
		screen.ImageAt(resource.ImageObake[(g.tick>>5)%4], g.drawRect.X0(), g.drawRect.Y0())
		screen.DebugLineX(g.hitRect, color.Black)
	case enemyStateDisappearing:
		screen.ImageAt(resource.ImageEffectIce[(g.tick/2)], g.drawRect.X0(), g.drawRect.Y0())
	}
}

func (g *ghost) Damage(d int) (score int) {
	g.life -= d
	if g.life > 0 {
		return 0
	}

	g.state = enemyStateDisappearing
	g.invincible = true
	g.tick = 0
	return 1
}

func (g *ghost) IsDisabled() bool {
	return g.disabled
}

func (g *ghost) IsInvincible() bool {
	return g.invincible
}

func (g *ghost) GetHitRect() *shape.Rect {
	return g.hitRect
}
