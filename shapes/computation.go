package shapes

import "bolijollo/rayos/primitives"

type Computation struct {
	T       float64
	Object  Shape
	Point   primitives.Tuple
	EyeV    primitives.Tuple
	NormalV primitives.Tuple
	Inside  bool
}
