package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	boundary []*Boundary
	ray      []*Ray
}

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()

	for _, r := range g.ray {
		r.lookAt(float64(mouseX), float64(mouseY))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, b := range g.boundary {
		b.Draw(screen)
	}
	for _, r := range g.ray {
		r.Draw(screen)
	}

	for _, r := range g.ray {
		for _, b := range g.boundary {
			pt := r.Cast(b)
			if pt != nil {
				white := color.RGBA{255, 255, 255, 255}
				ebitenutil.DrawCircle(screen, pt.X, pt.Y, 1, white)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	A := NewVector(300, 100)

	B := NewVector(300, 200)

	b := NewBoundary(A, B)

	X := NewVector(100, 150)

	dir := NewVector(1, 0)

	ray := NewRay(
		X, dir,
	)

	g := &Game{}

	g.boundary = append(g.boundary, b)
	g.ray = append(g.ray, ray)

	return g
}
