package game

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Normalize() {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	v.X = v.X / magnitude
	v.Y = v.Y / magnitude
}

func NewVector(x float64, y float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}
