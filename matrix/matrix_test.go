package matrix

import (
	pp "bolijollo/rayos/primitives"
	"math"
	"testing"
)

func TestDeterminantMatrix(t *testing.T) {
	m := CreateMatrix(4)
	m.Elements[0][0] = 5
	m.Elements[0][1] = 2
	m.Elements[0][2] = 6
	m.Elements[0][3] = -1
	m.Elements[1][0] = 1
	m.Elements[1][1] = -5
	m.Elements[1][2] = -2
	m.Elements[1][3] = 1
	m.Elements[2][0] = -2
	m.Elements[2][1] = 3
	m.Elements[2][2] = 4
	m.Elements[2][3] = -3
	m.Elements[3][0] = 0
	m.Elements[3][1] = 2
	m.Elements[3][2] = -1
	m.Elements[3][3] = 6
	if m.Determinant() != -196 {
		t.Error("Determinant failed")
	}

}

func TestInverseMatrix(t *testing.T) {
	m := CreateMatrix(4)
	m.Elements[0][0] = 5
	m.Elements[0][1] = 2
	m.Elements[0][2] = 6
	m.Elements[0][3] = -1
	m.Elements[1][0] = 1
	m.Elements[1][1] = -5
	m.Elements[1][2] = -2
	m.Elements[1][3] = 1
	m.Elements[2][0] = -2
	m.Elements[2][1] = 3
	m.Elements[2][2] = 4
	m.Elements[2][3] = -3
	m.Elements[3][0] = 0
	m.Elements[3][1] = 2
	m.Elements[3][2] = -1
	m.Elements[3][3] = 6
	l := m.Inverse()
	s := m.MMulti(l)
	eps := 0.000001
	if math.Abs(s.Elements[0][0]-1) > eps || math.Abs(s.Elements[0][1]) > eps ||
		math.Abs(s.Elements[0][2]) > eps || math.Abs(s.Elements[0][3]) > eps ||
		math.Abs(s.Elements[1][0]) > eps || math.Abs(s.Elements[1][1]-1) > eps ||
		math.Abs(s.Elements[1][2]) > eps || math.Abs(s.Elements[1][3]) > eps ||
		math.Abs(s.Elements[2][0]) > eps || math.Abs(s.Elements[2][1]) > eps ||
		math.Abs(s.Elements[2][2]-1) > eps || math.Abs(s.Elements[2][3]) > eps ||
		math.Abs(s.Elements[3][0]) > eps || math.Abs(s.Elements[3][1]) > eps ||
		math.Abs(s.Elements[3][2]) > eps || math.Abs(s.Elements[3][3]-1) > eps {
		t.Error("Inverse failed")
	}

}

func TestMatrixMultiplyTuple(t *testing.T) {
	m := CreateMatrix(4)
	m.Elements[0][0] = 1
	m.Elements[0][1] = 2
	m.Elements[0][2] = 3
	m.Elements[0][3] = 4
	m.Elements[1][0] = 2
	m.Elements[1][1] = 4
	m.Elements[1][2] = 4
	m.Elements[1][3] = 2
	m.Elements[2][0] = 8
	m.Elements[2][1] = 6
	m.Elements[2][2] = 4
	m.Elements[2][3] = 1
	m.Elements[3][0] = 0
	m.Elements[3][1] = 0
	m.Elements[3][2] = 0
	m.Elements[3][3] = 1
	t1 := pp.Point(1, 2, 3)
	t2 := m.VMulti(&t1)
	if t2.X != 18 || t2.Y != 24 || t2.Z != 33 || t2.W != 1 {
		t.Error("Matrix multiply tuple failed")
	}

}
