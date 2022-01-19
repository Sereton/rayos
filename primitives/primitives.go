package primitives

import "math"

type Tuple struct {
	X, Y, Z, W float64
}

func (t *Tuple) Point() bool {
	return t.W == 1.0
}

func (t *Tuple) Vector() bool {
	return t.W == 0.0

}

func (t Tuple) Add(o Tuple) Tuple {
	return Tuple{t.X + o.X, t.Y + o.Y, t.Z + o.Z, t.W + o.W}
}

func (t Tuple) Minus() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, t.W}
}

func (t Tuple) Sub(o Tuple) Tuple {
	return Tuple{t.X - o.X, t.Y - o.Y, t.Z - o.Z, t.W - o.W}
}

func (t Tuple) Scale(f float64) Tuple {
	return Tuple{t.X * f, t.Y * f, t.Z * f, t.W}
}

func (t *Tuple) Div(f float64) Tuple {
	return Tuple{t.X / f, t.Y / f, t.Z / f, t.W}
}

func (t Tuple) Dot(o Tuple) float64 {
	return t.X*o.X + t.Y*o.Y + t.Z*o.Z
}

func (t Tuple) Cross(o Tuple) Tuple {
	return Tuple{t.Y*o.Z - t.Z*o.Y, t.Z*o.X - t.X*o.Z, t.X*o.Y - t.Y*o.X, 0}
}

func (t *Tuple) Length() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func (t Tuple) Normalize() Tuple {
	return t.Div(t.Length())
}

func (t Tuple) Reflect(n Tuple) Tuple {
	return t.Sub(n.Scale(2 * t.Dot(n)))
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}
