package game

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	boundary []*Boundary
	particle *Particle
}

const (
	ScreenWidth  = 1280
	ScreenHeight = 768
)

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()

	g.particle.setPos(float64(mouseX), float64(mouseY))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, b := range g.boundary {
		b.Draw(screen)
	}
	g.particle.Draw(screen)

	for _, r := range g.particle.rays {
		minDis := math.Inf(0)
		var minpt *Vector
		minpt = nil
		for _, b := range g.boundary {
			insc := r.Cast(b)
			if insc != nil {
				dis := r.X.distanceFrom(insc)
				if dis <= minDis {
					minDis = dis
					minpt = insc
				}
			}
		}
		if minpt != nil {
			white := color.RGBA{255, 255, 255, 255}
			ebitenutil.DrawLine(screen, r.X.X, r.X.Y, minpt.X, minpt.Y, white)
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	g := &Game{}

	g.particle = NewParticle()

	b1 := NewBoundary(&Vector{X: 0, Y: 0}, &Vector{X: ScreenWidth, Y: 0})
	b2 := NewBoundary(&Vector{X: 0, Y: 0}, &Vector{X: 0, Y: ScreenHeight})
	b3 := NewBoundary(&Vector{X: ScreenWidth, Y: 0}, &Vector{X: ScreenWidth, Y: ScreenHeight})
	b4 := NewBoundary(&Vector{X: ScreenWidth, Y: ScreenHeight}, &Vector{X: 0, Y: ScreenHeight})

	g.boundary = append(g.boundary, b1)
	g.boundary = append(g.boundary, b2)
	g.boundary = append(g.boundary, b3)
	g.boundary = append(g.boundary, b4)

	for i := 0; i < 10; i++ {
		x1 := float64(rand.Intn(ScreenWidth + 1))
		y1 := float64(rand.Intn(ScreenHeight + 1))

		disx := float64(rand.Intn(ScreenWidth))
		disy := float64(rand.Intn(ScreenHeight))

		x2 := x1 + disx
		y2 := y1 + disy

		g.boundary = append(g.boundary, NewBoundary(NewVector(x1, y1), NewVector(x2, y2)))
	}

	return g
}
