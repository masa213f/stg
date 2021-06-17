package debug

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/shape"
)

var debugMode bool

var (
	Print     func(screen *ebiten.Image, line string)
	DrawLineX func(r *shape.Rect, clr color.Color)
	DrawLineV func(r *shape.Rect, clr color.Color)
)

func SetMode(flag bool) {
	debugMode = flag
	if debugMode {
		Print = debugPrint
		DrawLineV = drawLineV
		DrawLineX = drawLineX
	} else {
		Print = debugPrintDummy
		DrawLineV = drawLineDummy
		DrawLineX = drawLineDummy
	}
}

func GetMode() bool {
	return debugMode
}

func debugPrint(screen *ebiten.Image, line string) {
	ebitenutil.DebugPrint(screen, line)
}

func debugPrintDummy(_ *ebiten.Image, _ string) {
}

func drawLineV(r *shape.Rect, clr color.Color) {
	x0 := r.X0()
	x1 := r.X1()
	y0 := r.Y0()
	y1 := r.Y1()
	draw.Line(x0, y0, x0, y1, clr)
	draw.Line(x1, y0, x1, y1, clr)
}

func drawLineX(r *shape.Rect, clr color.Color) {
	x0 := r.X0()
	x1 := r.X1()
	y0 := r.Y0()
	y1 := r.Y1()
	draw.Line(x0, y0, x1, y1, clr)
	draw.Line(x0, y1, x1, y0, clr)
}

func drawLineDummy(_ *shape.Rect, _ color.Color) {
}
