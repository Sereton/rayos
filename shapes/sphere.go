package shapes

import (
	"bolijollo/rayos/matrix"
	"time"
)

type Sphere struct {
	Id       int64
	T_Matrix matrix.Matrix
}

func (s Sphere) Intersect(ray *Ray) []float64 {
	return []float64{}
}
func UniqueSphere() Sphere {
	m := matrix.Identity(4)
	return Sphere{Id: time.Now().UnixNano(), T_Matrix: m}
}

func (s Sphere) ChangeTransformation(m matrix.Matrix) Sphere {
	s.T_Matrix = m.MMulti(&s.T_Matrix)
	return s
}
