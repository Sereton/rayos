package shapes

import (
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
	"fmt"
	"math"
	"testing"
)

func TestNotTwoSpheresHaveSameID(t *testing.T) {
	s1 := UniqueSphere()
	s2 := UniqueSphere()
	if s1.Id == s2.Id {
		t.Error("Two spheres have same ID")
	}
	fmt.Println("s1.Id:", s1.Id)
	fmt.Println("s2.Id:", s2.Id)
}

func TestNormalAtSpherePoint(t *testing.T) {
	eps := 0.00001
	s := UniqueSphere()
	p := primitives.Point(1, 0, 0)
	n := s.NormalAt(&p)
	if n.X != 1 || n.Y != 0 || n.Z != 0 {
		t.Error("Sphere normal at point is not (1, 0, 0)")
	}

	s.T_Matrix = matrix.Translation(0, 1, 0)
	p2 := primitives.Point(0, 1.70711, -0.70711)
	n2 := s.NormalAt(&p2)
	if math.Abs(n2.X-0) > eps || math.Abs(n2.Y-0.70711) > eps || math.Abs(n2.Z+0.70711) > eps || n2.W != 0 {
		fmt.Println("n2:", n2)
		t.Error("Sphere normal at point is not (0, 0.70711, -0.70711,0)")
	}

}

func TestSphereImplementInterfaceShape(t *testing.T) {
	s := UniqueSphere()
	ray := Ray{Origin: primitives.Point(0, 0, -5), Direction: primitives.Vector(0, 0, 1)}
	var i Shape = &s
	if i.Intersect(&ray) == nil {
		t.Error("Sphere does not implement Shape interface")
	}
}
