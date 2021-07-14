package script

import (
	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/stage/enemy"
)

func NewStage1() Script {
	return newScript([]*step{
		waitFrame(180).
			showEnemies(
				enemy.NewGhost(constant.ScreenWidth+16, constant.ScreenHeight/2),
				enemy.NewGhost(constant.ScreenWidth+16, (constant.ScreenHeight/2)+80),
				enemy.NewGhost(constant.ScreenWidth+16, (constant.ScreenHeight/2)-80),
			),
		waitFrame(180).
			showEnemies(
				enemy.NewGhost(constant.ScreenWidth+16, constant.ScreenHeight/2),
				enemy.NewGhost(constant.ScreenWidth+16, (constant.ScreenHeight/2)+80),
				enemy.NewGhost(constant.ScreenWidth+16, (constant.ScreenHeight/2)-80),
			),
		waitAllEnemiesGone().
			showEnemies(
				enemy.NewGhost(constant.ScreenWidth+16, constant.ScreenHeight/2),
				enemy.NewGhost(constant.ScreenWidth+16, (constant.ScreenHeight/2)+80),
				enemy.NewGhost(constant.ScreenWidth+16, (constant.ScreenHeight/2)-80),
			),
		waitAllEnemiesGone(),
	})
}
