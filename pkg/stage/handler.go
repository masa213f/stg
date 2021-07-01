package stage

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/input"
	"github.com/masa213f/stg/pkg/shape"
	"github.com/masa213f/stg/pkg/sound"
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

type objectID uint64

const inactiveObjectID = ^objectID(0)

const (
	shotInterval = 3
)

// Handler is a object for managing a game.
type Handler struct {
	tick     int // General purpose counter.
	shotWait int
	bombWait int

	script    stageScript
	score     int
	life      int
	shotSpeed int

	// game objects
	background  Background
	player      Player
	playerBomb  PlayerBomb
	playerShots PlayerShots
	enemyList   *EnemyList
}

// NewHandler returns a new Hander struct.
func NewHandler() *Handler {
	h := &Handler{}
	h.Init()
	return h
}

// Init initializes the Handler struct.
func (h *Handler) Init() {
	h.tick = 0
	h.shotWait = 0
	h.bombWait = 0
	h.script = newStageScript()

	h.life = 3
	h.shotSpeed = 6

	h.background = newBackground()
	h.player = newPlayer(100, constant.ScreenHeight/2)
	h.playerBomb = newPlayerBomb()
	h.playerShots = newPlayerShots()
	h.enemyList = newEnemyList()
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
	if h.bombWait == 0 && input.Bomb() {
		h.bombWait = BombDuration
		return InputActionBomb
	}
	if h.shotWait == 0 && input.Shot() {
		h.shotWait = shotInterval
		return InputActionShot
	}
	return InputActionNone
}

// HitTest: player bomb -> enemy
func (h *Handler) hitTestPlayerBombToEnemy() {
	if h.playerBomb.IsActive() {
		for _, e := range h.enemyList.GetList() {
			if e.IsInvincible() {
				continue
			}
			if shape.Overlap(h.playerBomb.GetHitRect(), e.GetHitRect()) {
				sound.SE.Play(resource.SEHit)
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
		for _, e := range h.enemyList.GetList() {
			if e.IsInvincible() {
				continue
			}
			if shape.Overlap(shot, e.GetHitRect()) {
				sound.SE.Play(resource.SEHit)
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
	for _, e := range h.enemyList.GetList() {
		if h.player.IsInvincible() {
			break
		}
		if e.IsDisabled() || e.IsInvincible() {
			continue
		}
		if shape.Overlap(e.GetHitRect(), h.player.GetHitRect()) {
			sound.SE.Play(resource.SEDamage)
			e.Damage(1)
			h.player.Damage()
			h.life--
		}
	}
}

// Update updates game objects. This function is called every frame.
func (h *Handler) Update() error {
	h.tick++

	// stage
	if h.tick == 1 {
		sound.BGM.Reset(resource.BGMPlay)
	}
	h.background.Update()

	cond := &condition{enemyCount: h.enemyList.Count()}
	if h.script.NextEvent(cond) {
		for _, evt := range h.script.ShowEnemy() {
			h.enemyList.Add([]Enemy{newEnemy(evt.x, evt.y)})
		}
	} else {
		sound.BGM.Pause()
		return errors.New("gameover")
	}

	if h.life == 0 {
		sound.BGM.Pause()
		return errors.New("gameover")
	}
	h.enemyList.Update()
	h.player.Update()
	h.playerBomb.Update()
	h.playerShots.Update()

	px := h.player.GetCentorPoint().X()
	py := h.player.GetCentorPoint().Y()

	switch h.Input() {
	case InputActionBomb:
		sound.SE.Play(resource.SEBomb)
		h.playerBomb.NewBomb(px, py)
	case InputActionShot:
		sound.SE.Play(resource.SEShot)
		h.playerShots.NewShot(px, py, h.shotSpeed, -15)
		h.playerShots.NewShot(px, py, h.shotSpeed, 0)
		h.playerShots.NewShot(px, py, h.shotSpeed, 15)
	}

	h.hitTestPlayerBombToEnemy()
	h.hitTestPlayerShotToEnemy()
	h.hitTestPlayerToEnemy()

	return nil
}

// Draw draws game objects.
func (h *Handler) Draw() {
	h.background.Draw()
	h.player.Draw()
	h.playerBomb.Draw()
	h.playerShots.Draw()
	h.enemyList.Draw()
	draw.Text(resource.FontArcadeSmall, color.White, draw.HorizontalAlignRight, draw.VerticalAlignTop, fmt.Sprintf("Life: %2d, Score: %04d", h.life, h.score))
}