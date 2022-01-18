package world

import (
	"bolijollo/rayos/color"
	"bolijollo/rayos/lights"
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
	"bolijollo/rayos/shapes"
	"fmt"
	"testing"
)

func TestWorldCreation(t *testing.T) {
	w := CreateWorld()
	if w.Objects == nil {
		t.Error("World.Objects should be initialized to an empty slice")
	}
	if w.Lights == nil {
		t.Error("World.Lights should be initialized to an empty slice")
	}
}

func TestWorldAssignment(t *testing.T) {
	w := CreateWorld()
	s1 := shapes.UniqueSphere()
	s2 := shapes.UniqueSphere()
	s1.Material.Color = color.NewColor(200, 255, 200)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2
	s2.Material.Color = color.NewColor(255, 200, 200)
	s2.T_Matrix = matrix.Identity(4)
	light := lights.CreatePointLight(primitives.Point(0, 0, -10), color.NewColor(255, 255, 255))
	w.AddObject(s1)
	w.AddObject(&s2)
	w.AddLight(&light)

	if len(w.Objects) != 2 {
		t.Error("World.Objects should have 2 elements")
	}
	if len(w.Lights) != 1 || w.Lights[0] != &light {
		t.Error("World.Lights should have 1 element and be equal to the light")
	}

}

func TestIntersectAWorldWithARay(t *testing.T) {
	world := DefaultWorld()
	ray := shapes.Ray{Origin: primitives.Point(0, 0, -5), Direction: primitives.Vector(0, 0, 1)}
	xs := world.Intersect(&ray)
	if len(xs) != 4 || xs[0] != 4 || xs[1] != 4.5 || xs[2] != 5.5 || xs[3] != 6 {
		t.Error("The world should have 4 intersections with the ray")
		fmt.Println(xs)
	}
}
