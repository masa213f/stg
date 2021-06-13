package shape

import (
	"math"
)

// 固定小数点
type fixed int

// 丸め誤差を少なくするために内部でビットシフトして計算する。
const shift = 8

func intToFixed(n int) fixed {
	return fixed(n << shift)
}

func floatToFixed(n float64) fixed {
	return fixed(n * (1 << shift))
}

func fixedToInt(n fixed) int {
	return int(n >> shift)
}

// Shape は図形を表すインターフェース
type Shape interface {
	MoveX(vx int)
	MoveY(vy int)
	Move(v *Vector)
}

var sin [360]fixed
var cos [360]fixed

func init() {
	for i := 0; i < 360; i++ {
		sin[i] = floatToFixed(math.Sin(math.Pi * 2 * float64(i) / float64(360)))
		cos[i] = floatToFixed(math.Cos(math.Pi * 2 * float64(i) / float64(360)))
	}
}

// Vector はベクトル
type Vector struct {
	x fixed
	y fixed
}

func NewVector(x, y int) *Vector {
	return (&Vector{}).Reset(x, y)
}

func (v *Vector) Reset(x, y int) *Vector {
	v.x = intToFixed(x)
	v.y = intToFixed(y)
	return v
}

func NewVectorP(r, s int) *Vector {
	return (&Vector{}).ResetP(r, s)
}

func (v *Vector) ResetP(r, t int) *Vector {
	t = (t + 360) % 360
	v.x = fixed(r) * cos[t]
	v.y = fixed(r) * sin[t]
	return v
}

// 単位ベクトルに変換
func (v *Vector) Unit() *Vector {
	sqrt := math.Sqrt(float64(v.x*v.x + v.y*v.y))
	v.x = floatToFixed(float64(v.x) / sqrt)
	v.y = floatToFixed(float64(v.y) / sqrt)
	return v
}

func (v *Vector) Scale(n int) *Vector {
	v.x *= fixed(n)
	v.y *= fixed(n)
	return v
}

// Point は点を表す構造体
type Point struct {
	x fixed
	y fixed
}

// NewPoint は新しい Point を生成する
func NewPoint(x, y int) *Point {
	return (&Point{}).Reset(x, y)
}

// Reset ...
func (p *Point) Reset(x, y int) *Point {
	p.x = intToFixed(x)
	p.y = intToFixed(y)
	return p
}

// MoveX ...
func (p *Point) MoveX(vx int) *Point {
	p.x += intToFixed(vx)
	return p
}

// MoveY ...
func (p *Point) MoveY(vy int) *Point {
	p.y += intToFixed(vy)
	return p
}

// Move ...
func (p *Point) Move(v *Vector) *Point {
	p.x += v.x
	p.y += v.y
	return p
}

// X ...
func (p *Point) X() int {
	return fixedToInt(p.x)
}

// Y ...
func (p *Point) Y() int {
	return fixedToInt(p.y)
}

// Rect は矩形を表す構造体
type Rect struct {
	// (x0, y0) - (x1, y0)
	//  |               |
	// (x0, y1) - (x1, y1)
	x0 fixed
	y0 fixed
	x1 fixed
	y1 fixed
}

// NewRect は新しい Rect を生成する
func NewRect(x, y, w, h int) *Rect {
	return (&Rect{}).Reset(x, y, w, h)
}

// Reset ...
func (r *Rect) Reset(x, y, w, h int) *Rect {
	r.x0 = intToFixed(x)
	r.y0 = intToFixed(y)
	r.x1 = intToFixed(x + w)
	r.y1 = intToFixed(y + h)
	return r
}

// MoveX ...
func (r *Rect) MoveX(vx int) *Rect {
	r.x0 += intToFixed(vx)
	r.x1 += intToFixed(vx)
	return r
}

// MoveY ...
func (r *Rect) MoveY(vy int) *Rect {
	r.y0 += intToFixed(vy)
	r.y1 += intToFixed(vy)
	return r
}

// Move ...
func (r *Rect) Move(v *Vector) *Rect {
	r.x0 += v.x
	r.y0 += v.y
	r.x1 += v.x
	r.y1 += v.y
	return r
}

// X0 ...
func (r *Rect) X0() int {
	return fixedToInt(r.x0)
}

// X1 ...
func (r *Rect) X1() int {
	return fixedToInt(r.x1)
}

// Y0 ...
func (r *Rect) Y0() int {
	return fixedToInt(r.y0)
}

// Y1 ...
func (r *Rect) Y1() int {
	return fixedToInt(r.y1)
}

// Overlap ...
func Overlap(a, b *Rect) bool {
	return !(a.x0 > b.x1 || a.x1 < b.x0 || a.y0 > b.y1 || a.y1 < b.y0)
}
