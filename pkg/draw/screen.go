package draw

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/masa213f/stg/pkg/shape"
	"golang.org/x/image/font"
)

type Font struct {
	face font.Face
	size int
}

func NewFont(face font.Face, size int) *Font {
	return &Font{face: face, size: size}
}

type HorizontalAlign int

const (
	HorizontalAlignLeft   HorizontalAlign = iota // 左揃え
	HorizontalAlignCenter                        // 左右中央揃え
	HorizontalAlignRight                         // 右揃え
)

type VerticalAlign int

const (
	VerticalAlignTop    VerticalAlign = iota // 上揃え
	VerticalAlignMiddle                      // 上下中央揃え
	VerticalAlignBottom                      // 下揃え
)

var screen *ebiten.Image
var screenWidth int
var screenHeight int

// SetScreen は描画先イメージをセットするための関数。
func SetScreen(img *ebiten.Image) {
	screen = img
}

// SetScreenSize は描画先イメージのサイズをセットする関数。
func SetScreenSize(w, h int) {
	screenWidth = w
	screenHeight = h
}

// Fill は描画先イメージを指定された色で塗りつぶす。
func Fill(clr color.Color) {
	screen.Fill(clr)
}

// Line ...
func Line(x1, y1, x2, y2 int, clr color.Color) {
	ebitenutil.DrawLine(screen, float64(x1), float64(y1), float64(x2), float64(y2), clr)
}

// LineX
func LineX(r *shape.Rect, clr color.Color) {
	x0 := r.X0()
	x1 := r.X1()
	y0 := r.Y0()
	y1 := r.Y1()
	Line(x0, y0, x1, y1, clr)
	Line(x0, y1, x1, y0, clr)
}

// Rect
func Rect(r *shape.Rect, clr color.Color) {
	x0 := r.X0()
	x1 := r.X1()
	y0 := r.Y0()
	y1 := r.Y1()
	Line(x0, y0, x0, y1, clr)
	Line(x0, y1, x1, y1, clr)
	Line(x1, y1, x1, y0, clr)
	Line(x1, y0, x0, y0, clr)
}

// ImageAt ...
func ImageAt(img *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, op)
}

// Text ...
func Text(f *Font, clr color.Color, hAlign HorizontalAlign, vAlign VerticalAlign, line string) {
	var x int
	switch hAlign {
	case HorizontalAlignCenter:
		x = (screenWidth - len(line)*f.size) / 2
	case HorizontalAlignRight:
		x = screenWidth - len(line)*f.size
	}

	var y int
	switch vAlign {
	case VerticalAlignTop:
		y = f.size
	case VerticalAlignMiddle:
		y = ((screenHeight - f.size) / 2) + f.size
	case VerticalAlignBottom:
		y = screenHeight
	}

	text.Draw(screen, line, f.face, x, y, clr)
}

// MultiText ...
func MultiText(f *Font, clr color.Color, hAlign HorizontalAlign, vAlign VerticalAlign, texts []string) {
	var y0 int
	switch vAlign {
	case VerticalAlignTop:
		y0 = f.size
	case VerticalAlignMiddle:
		y0 = (screenHeight-len(texts)*f.size)/2 + f.size
	case VerticalAlignBottom:
		y0 = screenHeight - (len(texts)-1)*f.size
	}

	for i, line := range texts {
		var x int
		switch hAlign {
		case HorizontalAlignCenter:
			x = (screenWidth - len(line)*f.size) / 2
		case HorizontalAlignRight:
			x = screenWidth - len(line)*f.size
		}

		text.Draw(screen, line, f.face, x, y0+i*f.size, clr)
	}
}
