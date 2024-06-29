package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Boundary struct {
	A *Vector
	B *Vector
}

func (b *Boundary) Draw(screen *ebiten.Image) {
	white := color.RGBA{255, 255, 255, 255}

	ebitenutil.DrawLine(screen, b.A.X, b.A.Y, b.B.X, b.B.Y, white)
}

func NewBoundary(A *Vector, B *Vector) *Boundary {
	b := &Boundary{
		A: A,
		B: B,
	}

	return b
}
