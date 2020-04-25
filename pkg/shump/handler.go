package shooting

import (
	"errors"
	"fmt"
	"image/color"
	"math/rand"

	"github.com/masa213f/shootinggame/pkg/constant"
	"github.com/masa213f/shootinggame/pkg/draw"
	"github.com/masa213f/shootinggame/pkg/input"
	"github.com/masa213f/shootinggame/resource"
)

// 更新及び描画する何か
// - 背景(1)
// - ステータス(1)

// 更新及び描画するオブジェクト(括弧内はオブジェクト数)
// - 自機(1)
// - 自機ボム(0～1)
// - 自機ショット(0～多数)
// - 敵(0～多数)
// - 敵ショット(0～多数)
// - 障害物(0～多数)

// 当たり判定の種類、判定順序
// 1. 自機ボム     -> 敵、敵ショット ... 敵にダメージ。敵ショットは消滅。
// 2. 自機ショット -> 敵             ... 敵にダメージ。自機ショットは消滅。
// 3. 敵           -> 自機           ... 自機、敵共にダメージ。
// 4. 敵ショット   -> 自機           ... 自機にダメージ。敵ショットは消滅。
// 5. 障害物       -> 自機           ... ダメージなし。障害物の移動に合わせて、自機も移動。挟まれると自機にダメージ。

// オブジェクトの状態
// - 通常 ... 更新、描画及び当たり判定を実施。
// - 無効(disabled) ... 何もしない。
// - 無敵(untouchable) ... 当たり判定を実施しない。更新及び描画は実施する。(自機の復帰直後や、敵の消滅エフェクト中など)

type objectID uint64

const inactiveObjectID = ^objectID(0)

func debugLineX(r *Rect, clr color.Color) {
	x0 := r.X0()
	x1 := r.X1()
	y0 := r.Y0()
	y1 := r.Y1()
	draw.Line(x0, y0, x1, y1, clr)
	draw.Line(x0, y1, x1, y0, clr)
}

func debugLineV(r *Rect, clr color.Color) {
	x0 := r.X0()
	x1 := r.X1()
	y0 := r.Y0()
	y1 := r.Y1()
	draw.Line(x0, y0, x0, y1, clr)
	draw.Line(x1, y0, x1, y1, clr)
}

const (
	shotInterval = 3
	bombInterval = 60
)

// Handler はシューティングゲームをハンドリングするための構造体
type Handler struct {
	tick          int // 汎用的なカウンタ
	enemyInterval int
	shotWait      int // 次にショットが打てるようになるまでの待ちフレーム数
	bombWait      int // 次にボムが打てるようになるまでの待ちフレーム数

	score     int
	life      int
	shotSpeed int

	// game objects
	background  *background
	player      *player
	playerBomb  *playerBomb
	playerShots *playerShotList
	enemyList   []*enemy
}

// NewHandler ...
func NewHandler() *Handler {
	h := &Handler{}
	h.Init()
	return h
}

// Init ...
func (h *Handler) Init() {
	h.tick = 0
	h.enemyInterval = 0
	h.shotWait = 0
	h.bombWait = 0

	h.life = 3
	h.shotSpeed = 6

	h.background = newBackground()
	h.player = newPlayer(100, constant.ScreenHeight/2)
	h.playerBomb = newPlayerBomb()
	h.playerShots = newPlayerShotList()
}

// Update ...
func (h *Handler) Update() error {
	h.tick++
	h.background.update()

	{
		if h.tick == 1 {
			resource.BGM.Reset(resource.BGMPlay)
		}
		h.enemyInterval--
		if h.enemyInterval < 0 {
			h.enemyInterval = 5
			h.enemyList = append(h.enemyList, newEnemy(constant.ScreenWidth+16, rand.Intn(constant.ScreenHeight)))
		}
	}

	// var px, py int
	{
		h.bombWait--
		h.shotWait--
		h.player.update()
		px := h.player.centor.X()
		py := h.player.centor.Y()
		if h.bombWait < 0 && input.Bomb() {
			resource.SE.Play(resource.SEBomb)
			h.bombWait = bombInterval
			h.playerBomb.new(px, py)
		}
		if h.shotWait < 0 && input.Shot() {
			resource.SE.Play(resource.SEShot)
			h.shotWait = shotInterval

			h.playerShots.new(px, py, h.shotSpeed, -15)
			h.playerShots.new(px, py, h.shotSpeed, 0)
			h.playerShots.new(px, py, h.shotSpeed, 15)
		}
	}

	{
		if h.bombWait >= 0 {
			h.playerBomb.update()

			for _, e := range h.enemyList {
				if e.untouchable {
					continue
				}
				if Overlap(h.playerBomb.hitRect, e.hitRect) {
					resource.SE.Play(resource.SEHit)
					h.score += e.damage(1)
				}
			}
		}
	}

	{
	OUTER:
		for i, ps := range h.playerShots.list() {
			if ps.id == inactiveObjectID {
				break
			}
			ps.update()

			// 自機ショット <-> 敵 の当たり判定
			// 画面外(出現直前)の敵に当たるかもしれないので、自機ショットが画面外でも当たり判定を行う。
			for _, e := range h.enemyList {
				if e.untouchable {
					continue
				}
				if Overlap(ps.hitRect, e.hitRect) {
					resource.SE.Play(resource.SEHit)
					h.score += e.damage(1)

					// 自機ショットも消滅するので、次の自機ショットへ
					h.playerShots.inactive(i)
					continue OUTER
				}
			}
		}
		h.playerShots.gc()
	}

	{
		for i := 0; i < len(h.enemyList); i++ {
			e := h.enemyList[i]
			e.update()

			// 自機 <-> 敵 の当たり判定
			// プレイヤーが無敵 又は ボム実行中はスキップする
			if !e.disabled && !h.player.invincible && Overlap(e.hitRect, h.player.hitRect) {
				resource.SE.Play(resource.SEDamage)
				e.damage(1)
				h.player.damage()
				h.life--
			}

			if e.disabled {
				h.enemyList = append(h.enemyList[:i], h.enemyList[i+1:]...)
			}
		}
	}

	if h.life == 0 {
		resource.BGM.Pause()
		return errors.New("gameover")
	}

	return nil
}

// Draw ...
func (h *Handler) Draw() {
	h.background.draw()
	h.player.draw()
	if h.bombWait >= 0 {
		h.playerBomb.draw()
	}
	h.playerShots.drawAll()

	for _, e := range h.enemyList {
		e.draw()
	}
	draw.Text(resource.FontArcadeSmall, color.White, draw.HorizontalAlignRight, draw.VerticalAlignTop, fmt.Sprintf("Life: %2d, Score: %04d", h.life, h.score))
}
