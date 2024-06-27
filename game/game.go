package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	boundary []*Boundary
}

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, b := range g.boundary {
		b.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	A := &Vector{
		X: 300,
		Y: 100,
	}
	B := &Vector{
		X: 300,
		Y: 200,
	}

	b := NewBoundary(A, B)

	g := &Game{}

	g.boundary = append(g.boundary, b)

	return g
}
