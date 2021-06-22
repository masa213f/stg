package stage

import (
	"image/color"

	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/debug"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/pkg/shape"
	"github.com/masa213f/stg/resource"
)

type Player interface {
	Update()
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

func newPlayer(x, y int) Player {
	return &playerImpl{
		centor:   shape.NewPoint(x, y),
		hitRect:  shape.NewRect(x-4, y-4, 8, 8),
		drawRect: shape.NewRect(x-16, y-20, 32, 32),
		speedTable: [9]*shape.Vector{
			input.MoveNone:       shape.NewVectorP(0, 0),
			input.MoveRight:      shape.NewVectorP(speed, 0),
			input.MoveLowerRight: shape.NewVectorP(speed, 45),
			input.MoveDown:       shape.NewVectorP(speed, 90),
			input.MoveLowerLeft:  shape.NewVectorP(speed, 135),
			input.MoveLeft:       shape.NewVectorP(speed, 180),
			input.MoveUpperLeft:  shape.NewVectorP(speed, 225),
			input.MoveUp:         shape.NewVectorP(speed, 270),
			input.MoveUpperRight: shape.NewVectorP(speed, 315),
		},
	}
}

func (p *playerImpl) Update() {
	p.tick++

	if p.invincible {
		p.invincibleTime--
		if p.invincibleTime < 0 {
			p.invincible = false
		}
	}

	// Move
	p.centor.Move(p.speedTable[input.Move()])
	p.hitRect.Move(p.speedTable[input.Move()])
	p.drawRect.Move(p.speedTable[input.Move()])

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
