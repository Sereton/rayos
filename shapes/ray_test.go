package shapes

import (
	"bolijollo/rayos/primitives"

	"fmt"
	"testing"
)

func TestMovingAlongADirection(t *testing.T) {
	origin := primitives.Point(2, 3, 4)
	direction := primitives.Vector(1, 1, 1)
	r := Ray{origin, direction}
	if r.Position_at(0) != origin {
		t.Error("Ray position at 0 failed")
	}
	if r.Position_at(1) != primitives.Point(3, 4, 5) {
		t.Error("Ray position at 1 failed")
	}
	if r.Position_at(-1) != primitives.Point(1, 2, 3) {
		fmt.Println("Ray position at -1 failed", r.Position_at(-1))
		t.Error("Ray position at -1 failed")
	}

}

func TestRaySphereIntersection(t *testing.T) {
	r := Ray{primitives.Point(0, 0, -5), primitives.Vector(0, 0, 1)}
	s := UniqueSphere()
	xs := r.IntersectShape(s)
	if len(xs) != 2 {
		t.Error("Ray-sphere intersection failed")
	}
	if xs[0].T != 4 || xs[1].T != 6 {
		t.Error("Ray-sphere intersection failed")
	}
}
