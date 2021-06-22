package stage

import (
	"errors"
	"fmt"
	"image/color"
	"math/rand"

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
	tick          int // General purpose counter.
	enemyInterval int
	shotWait      int
	bombWait      int

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
	h.enemyInterval = 0
	h.shotWait = 0
	h.bombWait = 0

	h.life = 3
	h.shotSpeed = 6

	h.background = newBackground()
	h.player = newPlayer(100, constant.ScreenHeight/2)
	h.playerBomb = newPlayerBomb()
	h.playerShots = newPlayerShots()
	h.enemyList = newEnemyList()
}

// Update updates game objects. This function is called every frame.
func (h *Handler) Update() error {
	if h.tick == 0 {
		sound.BGM.Reset(resource.BGMPlay)
	}
	if h.life == 0 {
		sound.BGM.Pause()
		return errors.New("gameover")
	}

	h.tick++

	h.background.Update()
	h.enemyList.Update()

	// Create enemies.
	h.enemyInterval--
	if h.enemyInterval < 0 {
		h.enemyInterval = 5
		h.enemyList.Add([]Enemy{newEnemy(constant.ScreenWidth+16, rand.Intn(constant.ScreenHeight))})
	}

	// Move player.
	h.player.Update()
	px := h.player.GetCentorPoint().X()
	py := h.player.GetCentorPoint().Y()

	// Fire a player bomb.
	h.bombWait--
	if h.bombWait < 0 && input.Bomb() {
		sound.SE.Play(resource.SEBomb)
		h.bombWait = BombDuration
		h.playerBomb.NewBomb(px, py)
	}
	h.playerBomb.Update()

	// Fire playser shots.
	h.shotWait--
	if h.shotWait < 0 && input.Shot() {
		sound.SE.Play(resource.SEShot)
		h.shotWait = shotInterval
		h.playerShots.NewShot(px, py, h.shotSpeed, -15)
		h.playerShots.NewShot(px, py, h.shotSpeed, 0)
		h.playerShots.NewShot(px, py, h.shotSpeed, 15)
	}
	h.playerShots.Update()

	// Collision detection: player bomb -> enemy
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

	// Collision detection: player shot -> enemy
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

	// Collision detection: enemy -> player
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
