package shapes

import "bolijollo/rayos/matrix"

type Shape interface {
	Intersect(ray *Ray) []float64
	GET_Matrix() matrix.Matrix
}
