package shapes

import (
	"testing"
)

func TestIntersectionCreation(t *testing.T) {
	ball := UniqueSphere()
	T := 3.5
	intersection := Intersection{T, ball}
	if intersection.T != T {
		t.Errorf("Expected %f, got %f", T, intersection.T)
	}

}

func TestIntersectionsFunction(t *testing.T) {
	s := UniqueSphere()
	intersections := intersections(Intersection{1.0, s}, Intersection{2.0, s})
	if len(intersections) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(intersections))
	}
	if intersections[0].T != 1 {
		t.Errorf("Expected %f, got %f", 1.0, intersections[0].T)
	}
	if intersections[1].T != 2 {
		t.Errorf("Expected %f, got %f", 2.0, intersections[1].T)
	}

}
