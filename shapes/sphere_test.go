package shapes

import (
	"fmt"
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
