package shapes

import (
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"

	"math"
)

type Ray struct {
	Origin    primitives.Tuple
	Direction primitives.Tuple
}

func (r *Ray) Position_at(t float64) primitives.Tuple {
	vector_to_point := r.Direction.Scale(t)
	return r.Origin.Add(vector_to_point)
}

func (r *Ray) IntersectShape(shape Shape) []Intersection {

	inverse_transform := shape.GET_Matrix()
	inverse_transform = inverse_transform.Inverse()
	ray_transformed := r.Transform(&inverse_transform)
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
	intersections = append(intersections, Intersection{t1, shape})
	intersections = append(intersections, Intersection{t2, shape})

	return intersections

}

func (ray *Ray) Transform(m *matrix.Matrix) Ray {
	new_ray := Ray{}
	new_ray.Origin = m.VMulti(&ray.Origin)
	new_ray.Direction = m.VMulti(&ray.Direction)

	return new_ray
}

func (ray *Ray) PrepareComputations(s Shape, i Intersection) Computation {

	comps := Computation{}
	comps.T = i.T
	comps.Object = i.Object

	comps.Point = ray.Position_at(i.T)
	comps.EyeV = ray.Direction.Minus()
	v := comps.Point

	comps.NormalV = comps.Object.NormalAt(v)

	if comps.EyeV.Dot(comps.NormalV) < 0 {

		comps.Inside = true
		comps.NormalV = comps.NormalV.Scale(-1)

	} else {
		comps.Inside = false
	}

	return comps
}
