package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Particle struct {
	pos    *Vector
	rays   []*Ray
	radius int
}

func (particle *Particle) setPos(x float64, y float64) {
	particle.pos.X = x
	particle.pos.Y = y

	for _, r := range particle.rays {
		r.setPos(x, y)
	}
}

func (particle *Particle) Draw(screen *ebiten.Image) {
	white := color.RGBA{255, 255, 255, 255}

	ebitenutil.DrawCircle(screen, particle.pos.X, particle.pos.Y, float64(particle.radius), white)

	for _, r := range particle.rays {
		r.Draw(screen)
	}
}

func NewParticle() *Particle {
	pos := NewVector(0, 0)
	particle := &Particle{
		pos:    pos,
		radius: 5,
	}

	for a := 0.0; a <= 360.0; a += 10.0 {
		rayPos := NewVector(pos.X, pos.Y)

		angle := a * (math.Pi / 180)

		ray := NewRay(rayPos, angle)
		particle.rays = append(particle.rays, ray)

	}

	return particle
}
