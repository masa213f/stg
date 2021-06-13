package shooting

import (
	"image/color"

	"github.com/masa213f/shootinggame/pkg/debug"
	"github.com/masa213f/shootinggame/pkg/shape"
)

type playerBomb struct {
	tick    int
	x       int
	y       int
	hitRect *shape.Rect
}

func newPlayerBomb() *playerBomb {
	return &playerBomb{
		hitRect: &shape.Rect{},
	}
}

func (pb *playerBomb) new(x, y int) {
	pb.tick = 0
	pb.x = x
	pb.y = y
	pb.hitRect.Reset(x-32, y-32, 64, 64)
}

func (pb *playerBomb) update() {
	pb.tick++
	pb.hitRect.Reset(pb.x-32-pb.tick*3, pb.y-32-pb.tick*3, 64+pb.tick*6, 64+pb.tick*6)
}

func (pb *playerBomb) draw() {
	// TODO
	debug.DrawLineV(pb.hitRect, color.Black)
	debug.DrawLineX(pb.hitRect, color.Black)
}
