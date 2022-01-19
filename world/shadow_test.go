package world

import (
	"bolijollo/rayos/primitives"

	"testing"
)

func TestShadowingCalculation(t *testing.T) {
	t.Run("There is no shadow when light is not collinear to point", func(t *testing.T) {
		w := DefaultWorld()
		p := primitives.Point(0, 10, 0)
		got := Is_Shadowed(*w, p)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("There is shadow when object between point and light", func(t *testing.T) {
		w := DefaultWorld()
		p := primitives.Point(10, -10, 10)
		got := Is_Shadowed(*w, p)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("No shadow if object is behind the light", func(t *testing.T) {

		w := DefaultWorld()
		p := primitives.Point(-20, 20, -20)
		got := Is_Shadowed(*w, p)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("There is no shadow when object is behind the point", func(t *testing.T) {
		w := DefaultWorld()
		p := primitives.Point(-2, 2, -2)
		got := Is_Shadowed(*w, p)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})
}
