package color

import (
	primitives "bolijollo/rayos/primitives"
)

type Color primitives.Tuple

func Hadamard(a, b *Color) Color {
	factor := 1.0 / 255.0
	return NewColor(a.X*b.X*factor, a.Y*b.Y*factor, a.Z*b.Z*factor)
}

func (c Color) Add(o Color) Color {
	return Color{c.X + o.X, c.Y + o.Y, c.Z + o.Z, 0}
}

func (c Color) Substract(o Color) Color {
	return Color{c.X - o.X, c.Y - o.Y, c.Z - o.Z, 0}
}

func (c Color) Scale(f float64) Color {
	return Color{c.X * f, c.Y * f, c.Z * f, 0}
}

func NewColor(a, b, c float64) Color {
	return Color{a, b, c, 0}
}

var White = NewColor(255, 255, 255)
var Black = NewColor(0, 0, 0)
