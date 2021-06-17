package shooting

import (
	"image/color"

	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/debug"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/pkg/shape"
	"github.com/masa213f/stg/resource"
)

type player struct {
	tick           int          // 汎用的なカウンタ
	invincible     bool         // 無敵状態かどうか
	invincibleTime int          // 無敵状態の持続時間
	centor         *shape.Point // 自機の中心(自機ショットの開始位置)
	hitRect        *shape.Rect  // 当たり範囲
	drawRect       *shape.Rect  // 描画範囲
	speedTable     [9]*shape.Vector
}

var speed = 4

func newPlayer(x, y int) *player {
	return &player{
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

func (p *player) damage() {
	p.invincible = true
	p.invincibleTime = 60
}

func (p *player) update() {
	p.tick++

	if p.invincible {
		p.invincibleTime--
		if p.invincibleTime < 0 {
			p.invincible = false
		}
	}

	// 移動
	p.centor.Move(p.speedTable[input.GameMove()])
	p.hitRect.Move(p.speedTable[input.GameMove()])
	p.drawRect.Move(p.speedTable[input.GameMove()])

	// 画面外に出た場合の調整量の計算
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

	// 画面外に出た場合の調整
	v := shape.NewVector(vx, vy)
	p.centor.Move(v)
	p.hitRect.Move(v)
	p.drawRect.Move(v)
}

func (p *player) draw() {
	if p.invincible && p.tick>>3%2 == 0 {
		// 無敵状態の場合は点滅する
		return
	}
	draw.ImageAt(resource.ImagePlayer[(p.tick>>5)%4], p.drawRect.X0(), p.drawRect.Y0())
	debug.DrawLineV(p.drawRect, color.White)
	debug.DrawLineX(p.hitRect, color.White)
}
