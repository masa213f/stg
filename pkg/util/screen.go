package util

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
	HorizontalAlignLeft HorizontalAlign = iota
	HorizontalAlignCenter
	HorizontalAlignRight
)

type VerticalAlign int

const (
	VerticalAlignTop VerticalAlign = iota
	VerticalAlignMiddle
	VerticalAlignBottom
)

type Screen interface {
	SetImage(image *ebiten.Image)
	Fill(clr color.Color)
	Line(x1, y1, x2, y2 int, clr color.Color)
	LineV(r *shape.Rect, clr color.Color)
	LineX(r *shape.Rect, clr color.Color)
	Rect(r *shape.Rect, clr color.Color)
	ImageAt(img *ebiten.Image, x, y int)
	Text(f *Font, clr color.Color, hAlign HorizontalAlign, vAlign VerticalAlign, line string)
	MultiText(f *Font, clr color.Color, hAlign HorizontalAlign, vAlign VerticalAlign, texts []string)
	DebugPrint(line string)
	DebugLineV(r *shape.Rect, clr color.Color)
	DebugLineX(r *shape.Rect, clr color.Color)
}

type screen struct {
	width          int
	height         int
	debugMode      bool
	image          *ebiten.Image
	debugPrintFunc func(line string)
	debugLineVFunc func(r *shape.Rect, clr color.Color)
	debugLineXFunc func(r *shape.Rect, clr color.Color)
}

func NewScreen(width, height int, debugMode bool) Screen {
	s := &screen{
		width:     width,
		height:    height,
		debugMode: debugMode,
	}
	if debugMode {
		s.debugPrintFunc = s.debugPrint
		s.debugLineVFunc = s.LineV
		s.debugLineXFunc = s.LineX
	} else {
		debugPrintDummy := func(_ string) {}
		drawLineDummy := func(_ *shape.Rect, _ color.Color) {}
		s.debugPrintFunc = debugPrintDummy
		s.debugLineVFunc = drawLineDummy
		s.debugLineXFunc = drawLineDummy
	}
	return s
}

// SetImage set a image to draw on.
func (s *screen) SetImage(image *ebiten.Image) {
	s.image = image
}

// Fill fills the screen image with the specified color.
func (s *screen) Fill(clr color.Color) {
	s.image.Fill(clr)
}

// Line draws a line.
func (s *screen) Line(x1, y1, x2, y2 int, clr color.Color) {
	ebitenutil.DrawLine(s.image, float64(x1), float64(y1), float64(x2), float64(y2), clr)
}

// LineV draws the diagonal of the specified rect.
func (s *screen) LineV(r *shape.Rect, clr color.Color) {
	x0 := r.X0()
	x1 := r.X1()
	y0 := r.Y0()
	y1 := r.Y1()
	s.Line(x0, y0, x0, y1, clr)
	s.Line(x1, y0, x1, y1, clr)
}

// LineX draws the diagonal of the specified rect.
func (s *screen) LineX(r *shape.Rect, clr color.Color) {
	x0 := r.X0()
	x1 := r.X1()
	y0 := r.Y0()
	y1 := r.Y1()
	s.Line(x0, y0, x1, y1, clr)
	s.Line(x0, y1, x1, y0, clr)
}

// Rect draws the specified rect.
func (s *screen) Rect(r *shape.Rect, clr color.Color) {
	x0 := r.X0()
	x1 := r.X1()
	y0 := r.Y0()
	y1 := r.Y1()
	s.Line(x0, y0, x0, y1, clr)
	s.Line(x0, y1, x1, y1, clr)
	s.Line(x1, y1, x1, y0, clr)
	s.Line(x1, y0, x0, y0, clr)
}

// ImageAt draws an image at the specified coordinates.
func (s *screen) ImageAt(img *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	s.image.DrawImage(img, op)
}

// Text prints a single line string.
func (s *screen) Text(f *Font, clr color.Color, hAlign HorizontalAlign, vAlign VerticalAlign, line string) {
	var x int
	switch hAlign {
	case HorizontalAlignCenter:
		x = (s.width - len(line)*f.size) / 2
	case HorizontalAlignRight:
		x = s.width - len(line)*f.size
	}

	var y int
	switch vAlign {
	case VerticalAlignTop:
		y = f.size
	case VerticalAlignMiddle:
		y = ((s.height - f.size) / 2) + f.size
	case VerticalAlignBottom:
		y = s.height
	}

	text.Draw(s.image, line, f.face, x, y, clr)
}

// MultiText prints multi-line string.
func (s *screen) MultiText(f *Font, clr color.Color, hAlign HorizontalAlign, vAlign VerticalAlign, texts []string) {
	var y0 int
	switch vAlign {
	case VerticalAlignTop:
		y0 = f.size
	case VerticalAlignMiddle:
		y0 = (s.height-len(texts)*f.size)/2 + f.size
	case VerticalAlignBottom:
		y0 = s.height - (len(texts)-1)*f.size
	}

	for i, line := range texts {
		var x int
		switch hAlign {
		case HorizontalAlignCenter:
			x = (s.width - len(line)*f.size) / 2
		case HorizontalAlignRight:
			x = s.width - len(line)*f.size
		}

		text.Draw(s.image, line, f.face, x, y0+i*f.size, clr)
	}
}

func (s *screen) DebugPrint(line string) {
	s.debugPrintFunc(line)
}

func (s *screen) debugPrint(line string) {
	ebitenutil.DebugPrint(s.image, line)
}

func (s *screen) DebugLineV(r *shape.Rect, clr color.Color) {
	s.debugLineVFunc(r, clr)
}

func (s *screen) DebugLineX(r *shape.Rect, clr color.Color) {
	s.debugLineXFunc(r, clr)
}
