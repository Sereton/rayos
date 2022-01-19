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

func (w *World) Intersect(ray *shapes.Ray) []shapes.Intersection {
	intersections := make([]shapes.Intersection, 0)
	for _, obj := range w.Objects {
		intersections_shape := ray.IntersectShape(obj)
		if len(intersections_shape) > 0 {

			intersections = append(intersections, intersections_shape...)

		}

	}
	//sort intersections before returning
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].T < intersections[j].T
	})

	return intersections

}

func Shade_Hit(w *World, comps *shapes.Computation) color.Color {
	return lights.Lighting(comps.Object.GET_Material(), w.Lights[0], comps.Point, comps.EyeV, comps.NormalV)
}

func Color_At(ray *shapes.Ray, world *World) color.Color {
	intersections := world.Intersect(ray)

	if intersections == nil {
		return color.Black
	}
	hit := shapes.Hit(intersections)
	if hit.T == 0 {
		return color.Black
	}
	// fmt.Println("intersections", intersections[1].T, "hit", hit, intersections[0].T)
	comps := ray.PrepareComputations(hit.Object, hit)
	return Shade_Hit(world, &comps)
}
