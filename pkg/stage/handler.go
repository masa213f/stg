package stage

import (
	"fmt"
	"image/color"

	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/shape"
	"github.com/masa213f/stg/pkg/stage/background"
	"github.com/masa213f/stg/pkg/stage/enemy"
	"github.com/masa213f/stg/pkg/stage/player"
	"github.com/masa213f/stg/pkg/stage/script"
	"github.com/masa213f/stg/pkg/util"
	"github.com/masa213f/stg/resource"
)

// Kind of Object (Number of objects)
// - player (1)
// - player bomb (0～1)
// - player shot (0～n)
// - enemy (0～n)

// Order of collision detection
// 1. player bomb -> enemy  ... Damage the enemy.
// 2. player shot -> enemy  ... Damage the enemy. The player shot disappears.
// 3. enemy       -> player ... Damage both the player and the enemy.

// Object status
// - normal ... Do update, drawing and collision detection.
// - disabled ... Do nothing.
// - untouchable ... Do update and drawing. No collision detection is performed.
//                  (Immediately after the return of the player, during the enemy's disappearance effect, etc.)

const (
	shotInterval = 3
)

type Result int

const (
	Playing Result = iota
	GameOver
	StageClear
)

// Handler is a object for managing a game.
type Handler struct {
	ctrl     util.Control
	bgm      util.BGMPlayer
	se       util.SEPlayer
	tick     int // General purpose counter.
	shotWait int
	bombWait int

	script    script.Script
	score     int
	life      int
	shotSpeed int

	// game objects
	background     background.Background
	player         player.Player
	playerBomb     player.PlayerBomb
	playerShots    player.PlayerShots
	enemyContainer *enemy.Container
}

// NewHandler returns a new Hander struct.
func NewHandler(ctrl util.Control, bgm util.BGMPlayer, se util.SEPlayer) *Handler {
	h := &Handler{
		ctrl: ctrl,
		bgm:  bgm,
		se:   se,
	}
	h.Init()
	return h
}

// Init initializes the Handler struct.
func (h *Handler) Init() {
	h.tick = 0
	h.shotWait = 0
	h.bombWait = 0
	h.script = script.NewStage1()

	h.life = 3
	h.shotSpeed = 6

	h.background = background.NewCloudBackground()
	h.player = player.NewPlayer(100, constant.ScreenHeight/2)
	h.playerBomb = player.NewPlayerBomb()
	h.playerShots = player.NewPlayerShots()
	h.enemyContainer = enemy.NewContainer()
}

type InputAction uint

const (
	InputActionNone InputAction = 0 // 0000
	InputActionBomb InputAction = 1 // 0001
	InputActionShot InputAction = 2 // 0010
)

func (h *Handler) Input() InputAction {
	if h.bombWait > 0 {
		h.bombWait--
	}
	if h.shotWait > 0 {
		h.shotWait--
	}
	if h.bombWait == 0 && h.ctrl.Bomb() {
		h.bombWait = player.BombDuration
		return InputActionBomb
	}
	if h.shotWait == 0 && h.ctrl.Shot() {
		h.shotWait = shotInterval
		return InputActionShot
	}
	return InputActionNone
}

// HitTest: player bomb -> enemy
func (h *Handler) hitTestPlayerBombToEnemy() {
	if h.playerBomb.IsActive() {
		for _, e := range h.enemyContainer.GetList() {
			if e.IsInvincible() {
				continue
			}
			if shape.Overlap(h.playerBomb.GetHitRect(), e.GetHitRect()) {
				h.se.Play(resource.SEHit)
				h.score += e.Damage(1)
			}
		}
	}
}

// HitTest: player shot -> enemy
func (h *Handler) hitTestPlayerShotToEnemy() {
	shotsHitRects := h.playerShots.GetHitRects()
OUTER:
	for i, shot := range shotsHitRects {
		for _, e := range h.enemyContainer.GetList() {
			if e.IsInvincible() {
				continue
			}
			if shape.Overlap(shot, e.GetHitRect()) {
				h.se.Play(resource.SEHit)
				h.score += e.Damage(1)

				// the shot disappears.
				h.playerShots.MakeInactive(i)
				continue OUTER
			}
		}
	}
}

// HitTest: player <-> enemy
func (h *Handler) hitTestPlayerToEnemy() {
	// Skip while the player is invincible or a player bomb running.
	for _, e := range h.enemyContainer.GetList() {
		if h.player.IsInvincible() {
			break
		}
		if e.IsDisabled() || e.IsInvincible() {
			continue
		}
		if shape.Overlap(e.GetHitRect(), h.player.GetHitRect()) {
			h.se.Play(resource.SEDamage)
			e.Damage(1)
			h.player.Damage()
			h.life--
		}
	}
}

// Update updates game objects. This function is called every frame.
func (h *Handler) Update() Result {
	h.tick++

	if h.tick == 1 {
		h.bgm.Reset(resource.BGMPlay)
	}
	if h.life == 0 {
		h.bgm.Pause()
		return GameOver
	}

	px := h.player.GetCentorPoint().X()
	py := h.player.GetCentorPoint().Y()

	switch h.script.NextEvent(px, py, h.enemyContainer.Count()) {
	case script.End:
		h.bgm.Pause()
		return StageClear
	case script.NewEnemies:
		h.enemyContainer.Add(h.script.NewEnemies()...)
	}

	h.background.Update()
	h.player.Update(h.ctrl.Move())
	h.playerBomb.Update()
	h.playerShots.Update()
	h.enemyContainer.UpdateAll()

	switch h.Input() {
	case InputActionBomb:
		h.se.Play(resource.SEBomb)
		h.playerBomb.NewBomb(px, py)
	case InputActionShot:
		h.se.Play(resource.SEShot)
		h.playerShots.NewShot(px, py, h.shotSpeed, -15)
		h.playerShots.NewShot(px, py, h.shotSpeed, 0)
		h.playerShots.NewShot(px, py, h.shotSpeed, 15)
	}

	h.hitTestPlayerBombToEnemy()
	h.hitTestPlayerShotToEnemy()
	h.hitTestPlayerToEnemy()

	return Playing
}

// Draw draws game objects.
func (h *Handler) Draw() {
	h.background.Draw()
	h.player.Draw()
	h.playerBomb.Draw()
	h.playerShots.Draw()
	h.enemyContainer.DrawAll()
	draw.Text(resource.FontArcadeSmall, color.White, draw.HorizontalAlignRight, draw.VerticalAlignTop, fmt.Sprintf("Life: %2d, Score: %04d", h.life, h.score))
}
