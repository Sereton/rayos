package shapes

import (
	"bolijollo/rayos/materials"
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
	"time"
)

type Sphere struct {
	Id       int64
	T_Matrix matrix.Matrix
	Material materials.Material
}

func (s Sphere) Intersect(ray *Ray) []float64 {
	return []float64{}
}
func UniqueSphere() Sphere {
	m := matrix.Identity(4)
	return Sphere{Id: time.Now().UnixNano(), T_Matrix: m, Material: materials.DefaultMaterial()}
}

func (s Sphere) ChangeTransformation(m matrix.Matrix) Sphere {
	s.T_Matrix = m.MMulti(s.T_Matrix)
	return s
}

func (s Sphere) NormalAt(p primitives.Tuple) primitives.Tuple {
	object_point := s.T_Matrix.Inverse().VMulti(&p)
	origin := primitives.Point(0, 0, 0)
	object_normal := object_point.Sub(origin)
	world_normal := s.T_Matrix.Inverse().Transpose().VMulti(&object_normal)
	world_normal.W = 0
	return world_normal.Normalize()

}

func (s Sphere) GET_Matrix() matrix.Matrix {
	return s.T_Matrix
}

func (s Sphere) GET_Material() materials.Material {
	return s.Material
}
