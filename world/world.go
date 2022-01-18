package world

import (
	"bolijollo/rayos/color"
	"bolijollo/rayos/lights"
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
	"bolijollo/rayos/shapes"
	"sort"
)

type World struct {
	Objects []shapes.Shape       // List of objects in the world
	Lights  []*lights.PointLight // List of lights in the world
}

func CreateWorld() *World {
	return &World{
		Objects: make([]shapes.Shape, 0),
		Lights:  make([]*lights.PointLight, 0),
	}
}
func DefaultWorld() *World {
	w := CreateWorld()
	s1 := shapes.UniqueSphere()
	s2 := shapes.UniqueSphere()
	s1.Material.Color = color.NewColor(200, 255, 150)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2
	s2.Material.Color = color.NewColor(255, 200, 200)
	s2.T_Matrix = matrix.Scaling(0.5, 0.5, 0.5)
	light := lights.CreatePointLight(primitives.Point(-10, 10, -10), color.NewColor(255, 255, 255))
	w.AddObject(s1)
	w.AddObject(&s2)
	w.AddLight(&light)
	return w
}
func (w *World) AddObject(obj shapes.Shape) {
	w.Objects = append(w.Objects, obj)
}

func (w *World) AddLight(light *lights.PointLight) {
	w.Lights = append(w.Lights, light)
}

func (w *World) Intersect(ray *shapes.Ray) []float64 {
	intersections := make([]float64, 0)
	for _, obj := range w.Objects {
		intersections_shape := ray.IntersectShape(obj)
		if len(intersections_shape) > 0 {
			for _, t := range intersections_shape {
				intersections = append(intersections, t.T)
			}
		}

	}
	//sort intersections before returning

	sort.Float64s(intersections)

	return intersections

}
