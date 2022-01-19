package camera

import (
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
	"testing"
)

func TestViewTransformation(t *testing.T) {

	t.Run("correctly transform default orientation", func(t *testing.T) {
		from := primitives.Point(0, 0, 0)
		to := primitives.Point(0, 0, -1)
		up := primitives.Vector(0, 1, 0)
		I := matrix.Identity(4)
		got := View_transformation(from, to, up)
		want := I

		if got.Equal(&want) == false {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("it correctly transforms whil looking in the  +Z direction", func(t *testing.T) {
		from := primitives.Point(0, 0, 0)
		to := primitives.Point(0, 0, 1)
		up := primitives.Vector(0, 1, 0)

		got := View_transformation(from, to, up)
		want := matrix.Scaling(-1, 1, -1)

		if got.Equal(&want) == false {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("it correctly transforms the default orientation", func(t *testing.T) {
		from := primitives.Point(0, 0, 0)
		to := primitives.Point(0, 0, -1)
		up := primitives.Vector(0, 1, 0)
		I := matrix.Identity(4)
		got := View_transformation(from, to, up)
		want := I

		if got.Equal(&want) == false {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("the view transformation moves the world", func(t *testing.T) {
		from := primitives.Point(0, 0, 8)
		to := primitives.Point(0, 0, 0)
		up := primitives.Vector(0, 1, 0)

		got := View_transformation(from, to, up)
		want := matrix.Translation(0, 0, -8)

		if got.Equal(&want) == false {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("an arbitrary transformation", func(t *testing.T) {
		from := primitives.Point(1, 3, 2)
		to := primitives.Point(4, -2, 8)
		up := primitives.Vector(1, 1, 0)

		W := matrix.Identity(4)
		W.Set(0, 0, -0.5070925528371099)
		W.Set(0, 1, 0.5070925528371099)
		W.Set(0, 2, 0.6761234037828132)
		W.Set(0, 3, -2.366431913239846)
		W.Set(1, 0, 0.7677159338596801)
		W.Set(1, 1, 0.6060915267313263)
		W.Set(1, 2, 0.12121830534626524)
		W.Set(1, 3, -2.8284271247461894)
		W.Set(2, 0, -0.35856858280031806)
		W.Set(2, 1, 0.5976143046671968)
		W.Set(2, 2, -0.7171371656006361)

		got := View_transformation(from, to, up)
		want := W

		if got.Equal(&want) == false {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
