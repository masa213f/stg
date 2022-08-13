package player

import (
	"image/color"

	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/debug"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/shape"
	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

type Player interface {
	Update(util.MoveDirection)
	Draw()
	Damage()
	IsInvincible() bool
	GetHitRect() *shape.Rect
	GetCentorPoint() *shape.Point
}

type playerImpl struct {
	tick           int
	invincible     bool
	invincibleTime int          // Invincible duration.
	centor         *shape.Point // Center of the player (start position of a player shot)
	hitRect        *shape.Rect
	drawRect       *shape.Rect
	speedTable     [9]*shape.Vector
}

var speed = 4

func NewPlayer(x, y int) Player {
	return &playerImpl{
		centor:   shape.NewPoint(x, y),
		hitRect:  shape.NewRect(x-4, y-4, 8, 8),
		drawRect: shape.NewRect(x-16, y-20, 32, 32),
		speedTable: [9]*shape.Vector{
			util.MoveNone:       shape.NewVectorP(0, 0),
			util.MoveRight:      shape.NewVectorP(speed, 0),
			util.MoveLowerRight: shape.NewVectorP(speed, 45),
			util.MoveDown:       shape.NewVectorP(speed, 90),
			util.MoveLowerLeft:  shape.NewVectorP(speed, 135),
			util.MoveLeft:       shape.NewVectorP(speed, 180),
			util.MoveUpperLeft:  shape.NewVectorP(speed, 225),
			util.MoveUp:         shape.NewVectorP(speed, 270),
			util.MoveUpperRight: shape.NewVectorP(speed, 315),
		},
	}
}

func (p *playerImpl) Update(direction util.MoveDirection) {
	p.tick++

	if p.invincible {
		p.invincibleTime--
		if p.invincibleTime < 0 {
			p.invincible = false
		}
	}

	// Move
	p.centor.Move(p.speedTable[direction])
	p.hitRect.Move(p.speedTable[direction])
	p.drawRect.Move(p.speedTable[direction])

	// Calculation of adjustment amount when going out of the screen.
	var vx, vy int
	if p.hitRect.X0() < 0 {
		vx = -p.hitRect.X0()
	} else if p.hitRect.X1() > constant.ScreenWidth {
		vx = constant.ScreenWidth - p.hitRect.X1()
	}
	if p.hitRect.Y0() < 0 {
		vy = -p.hitRect.Y0()
	} else if p.hitRect.Y1() > constant.ScreenHeight {
		vy = constant.ScreenHeight - p.hitRect.Y1()
	}

	// Adjustment when going out of the screen.
	v := shape.NewVector(vx, vy)
	p.centor.Move(v)
	p.hitRect.Move(v)
	p.drawRect.Move(v)
}

func (p *playerImpl) Draw() {
	if p.invincible && p.tick>>3%2 == 0 {
		// Flashes when invincible.
		return
	}
	draw.ImageAt(resource.ImagePlayer[(p.tick>>5)%4], p.drawRect.X0(), p.drawRect.Y0())
	debug.DrawLineV(p.drawRect, color.White)
	debug.DrawLineX(p.hitRect, color.White)
}

func (p *playerImpl) Damage() {
	p.invincible = true
	p.invincibleTime = 60
}

func (p *playerImpl) IsInvincible() bool {
	return p.invincible
}

func (p *playerImpl) GetHitRect() *shape.Rect {
	return p.hitRect
}

func (p *playerImpl) GetCentorPoint() *shape.Point {
	return p.centor
}
