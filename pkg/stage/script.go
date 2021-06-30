package stage

import (
	"github.com/masa213f/stg/pkg/constant"
)

type stageScript interface {
	// When it returns false, no events are left. This means the end of the stage.
	NextEvent(cond *condition) bool

	ShowEnemy() []*showEnemy
}

type condition struct {
	enemyCount int
}

type showEnemy struct {
	x int
	y int
}

func newStageScript() stageScript {
	return &stageScriptImpl{
		step: 0,
		last: 360,
		enemy: map[int][]*showEnemy{
			180: {
				{constant.ScreenWidth + 16, constant.ScreenHeight / 2},
				{constant.ScreenWidth + 16, (constant.ScreenHeight / 2) + 80},
				{constant.ScreenWidth + 16, (constant.ScreenHeight / 2) - 80},
			},
			360: {
				{constant.ScreenWidth + 16, constant.ScreenHeight / 2},
				{constant.ScreenWidth + 16, (constant.ScreenHeight / 2) + 80},
				{constant.ScreenWidth + 16, (constant.ScreenHeight / 2) - 80},
			},
		},
	}
}

type stageScriptImpl struct {
	step  int
	last  int
	enemy map[int][]*showEnemy
}

func (s *stageScriptImpl) NextEvent(cond *condition) bool {
	s.step++
	if s.step >= s.last && cond.enemyCount == 0 {
		return false
	}
	return true
}

func (s *stageScriptImpl) ShowEnemy() []*showEnemy {
	return s.enemy[s.step]
}
