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

func Add(t, o *Tuple) Tuple {
	return Tuple{t.X + o.X, t.Y + o.Y, t.Z + o.Z, t.W + o.W}
}

func Minus(t *Tuple) Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, t.W}
}

func Sub(t, o *Tuple) Tuple {
	return Tuple{t.X - o.X, t.Y - o.Y, t.Z - o.Z, t.W - o.W}
}

func Times(t *Tuple, f float64) Tuple {
	return Tuple{t.X * f, t.Y * f, t.Z * f, t.W}
}

func Div(t *Tuple, f float64) Tuple {
	return Tuple{t.X / f, t.Y / f, t.Z / f, t.W}
}

func Dot(t, o *Tuple) float64 {
	return t.X*o.X + t.Y*o.Y + t.Z*o.Z
}

func Cross(t, o *Tuple) Tuple {
	return Tuple{t.Y*o.Z - t.Z*o.Y, t.Z*o.X - t.X*o.Z, t.X*o.Y - t.Y*o.X, 0}
}

func Length(t *Tuple) float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func Normalize(t *Tuple) Tuple {
	l := Length(t)
	return Div(t, l)
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}
