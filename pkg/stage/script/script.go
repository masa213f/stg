package script

import (
	"github.com/masa213f/stg/pkg/stage/enemy"
)

type Event uint

const (
	Nop        Event = 0
	End        Event = 1
	NewEnemies Event = 2
)

type Script interface {
	// When it returns false, no events are left. This means the end of the stage.
	NextEvent(playerX, playerY, enemyNum int) Event
	NewEnemies() []enemy.Enemy
}

const (
	condWait int = iota
	condSkip
	condDo
)

type waitFunc func(wait, playerX, playerY, enemyNum int) int

type step struct {
	waitFlame          int
	waitAllEnemiesGone bool
	enemies            []enemy.Enemy
}

func waitFrame(frame int) *step {
	return &step{
		waitFlame: frame,
	}
}

func waitAllEnemiesGone() *step {
	return &step{
		waitAllEnemiesGone: true,
	}
}

func (s *step) cond(wait, playerX, playerY, enemyNum int) int {
	if s.waitFlame != 0 && s.waitFlame == wait {
		return condDo
	}
	if s.waitAllEnemiesGone && enemyNum == 0 {
		return condDo
	}
	return condWait
}

func (s *step) showEnemies(e ...enemy.Enemy) *step {
	s.enemies = e
	return s
}

type stageScriptImpl struct {
	wait  int
	index int
	steps []*step
}

func newScript(steps []*step) Script {
	return &stageScriptImpl{
		steps: steps,
	}
}

func (s *stageScriptImpl) NextEvent(playerX, playerY, enemyNum int) Event {
	if s.index >= len(s.steps) {
		return End
	}

	s.wait++
	switch s.steps[s.index].cond(s.wait, playerX, playerY, enemyNum) {
	case condWait:
		return Nop
	case condSkip:
		s.wait = 0
		s.index++
		return Nop
	case condDo:
		s.wait = 0
		s.index++
		if len(s.steps[s.index-1].enemies) != 0 {
			return NewEnemies
		}
		return End
	default:
		panic("TODO")
	}
}

func (s *stageScriptImpl) NewEnemies() []enemy.Enemy {
	return s.steps[s.index-1].enemies
}
