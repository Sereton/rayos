package shapes

import (
	"bolijollo/rayos/materials"
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
	"math"
	"time"
)

type Sphere struct {
	Id       int64
	T_Matrix matrix.Matrix
	Material materials.Material
}

func (s Sphere) Intersect(ray *Ray) []Intersection {
	inverse_transform := s.GET_Matrix()
	inverse_transform = inverse_transform.Inverse()
	ray_transformed := ray.Transform(&inverse_transform)
	center_sphere := primitives.Point(0, 0, 0)
	sphere_to_ray := ray_transformed.Origin.Sub(center_sphere)
	a := ray_transformed.Direction.Dot(ray_transformed.Direction)
	b := 2 * sphere_to_ray.Dot(ray_transformed.Direction)
	c := sphere_to_ray.Dot(sphere_to_ray) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return []Intersection{}
	}
	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	intersections := []Intersection{}
	intersections = append(intersections, Intersection{t1, s})
	intersections = append(intersections, Intersection{t2, s})

	return intersections

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
