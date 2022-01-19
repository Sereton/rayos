package shapes

import (
	"bolijollo/rayos/materials"
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
)

type Shape interface {
	Intersect(ray *Ray) []Intersection
	GET_Matrix() matrix.Matrix
	GET_Material() materials.Material
	NormalAt(p primitives.Tuple) primitives.Tuple
}
