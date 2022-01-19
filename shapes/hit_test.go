package shapes

import (
	"fmt"
	"testing"
)

func TestHitInASphere(t *testing.T) {
	s := UniqueSphere()

	i1 := Intersection{T: 0.5, Object: s}
	i2 := Intersection{T: 0.7, Object: s}
	i3 := Intersection{T: -0.4, Object: s}
	i4 := Intersection{T: 0.2, Object: s}

	intersections := intersections(i1, i2, i3, i4)
	i := Hit(intersections)
	fmt.Println("AHHHHHHHHHHH", i,
		"OOOOOOOOOOOOOO", intersections)
}
