package shooting

import "image/color"

type playerBomb struct {
	tick    int
	x       int
	y       int
	hitRect *Rect
}

func newPlayerBomb() *playerBomb {
	return &playerBomb{
		hitRect: &Rect{},
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
	debugLineV(pb.hitRect, color.Black)
	debugLineX(pb.hitRect, color.Black)
}
