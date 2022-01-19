package shapes

import (
	"bolijollo/rayos/materials"
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
	"math"
	"time"
)

type Plane struct {
	Id       int64
	TMatrix  matrix.Matrix
	Material materials.Material
}

func (p Plane) GET_Material() materials.Material {
	return p.Material
}

func (p Plane) GET_Matrix() matrix.Matrix {
	return p.TMatrix
}

func UniquePlane() Plane {
	m := matrix.Identity(4)
	return Plane{Id: time.Now().UnixNano(), TMatrix: m, Material: materials.DefaultMaterial()}
}

func (p Plane) NormalAt(point primitives.Tuple) primitives.Tuple {
	plane_normal := primitives.Vector(0, 1, 0)
	world_normal := p.TMatrix.Inverse().Transpose().VMulti(&plane_normal)
	world_normal.W = 0
	return world_normal.Normalize()
}

func (p Plane) Intersect(ray *Ray) []Intersection {
	eps := 0.0000001
	inverse_transform := p.GET_Matrix()
	inverse_transform = inverse_transform.Inverse()
	ray_transformed := ray.Transform(&inverse_transform)
	if math.Abs(ray_transformed.Direction.Y) <= eps {
		return []Intersection{}
	}

	t := -ray_transformed.Origin.Y / ray_transformed.Direction.Y
	intersections := []Intersection{}
	intersections = append(intersections, Intersection{t, p})

	return intersections

}
