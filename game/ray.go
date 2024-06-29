package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ray struct {
	X   *Vector
	dir *Vector
}

func (r *Ray) setPos(x float64, y float64) {
	r.X.X = x
	r.X.Y = y
}

func (r *Ray) Draw(screen *ebiten.Image) {
	white := color.RGBA{255, 255, 255, 255}

	ebitenutil.DrawLine(screen, r.X.X, r.X.Y, r.X.X+r.dir.X*10, r.X.Y+r.dir.Y*10, white)
}

func (r *Ray) Cast(boundary *Boundary) *Vector {
	x1 := boundary.A.X
	y1 := boundary.A.Y
	x2 := boundary.B.X
	y2 := boundary.B.Y

	x3 := r.X.X
	y3 := r.X.Y
	x4 := x3 + r.dir.X
	y4 := y3 + r.dir.Y

	den := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)

	if den == 0 {
		return nil
	}

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / den
	u := -((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / den

	if t > 0 && t < 1 && u > 0 {
		pt := NewVector(
			x1+t*(x2-x1),
			y1+t*(y2-y1),
		)
		return pt
	}

	return nil

}

func NewRay(X *Vector, angle float64) *Ray {
	dir := &Vector{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}

	r := &Ray{
		X:   X,
		dir: dir,
	}
	return r
}
