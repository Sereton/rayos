package primitives

import (
	"testing"
)

func TestTuplePointVector(t *testing.T) {
	a := Tuple{1, 2, 3, 1}
	b := Tuple{4, 5, 6, 0}
	if !a.Point() {
		t.Error("a is not a Point")
	}
	if !b.Vector() {
		t.Error("b is not a Vector")
	}
}

func TestVectorCreation(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(4, 5, 6)
	if a.X != 1 || a.Y != 2 || a.Z != 3 || a.W != 0 {
		t.Error("Vector creation failed")
	}
	if b.X != 4 || b.Y != 5 || b.Z != 6 || b.W != 0 {
		t.Error("Vector creation failed")
	}
}

func TestPointCreation(t *testing.T) {
	a := Point(1, 2, 3)
	b := Point(4, 5, 6)
	if a.X != 1 || a.Y != 2 || a.Z != 3 || a.W != 1 {
		t.Error("Point creation failed")
	}
	if b.X != 4 || b.Y != 5 || b.Z != 6 || b.W != 1 {
		t.Error("Point creation failed")
	}
}

func TestTupleAddition(t *testing.T) {
	a := Point(1, 2, 3)
	b := Vector(4, 5, 6)
	c := a.Add(b)
	if c.X != 5 || c.Y != 7 || c.Z != 9 || c.W != 1 {
		t.Error("Point Addition failed")
	}
	if !c.Point() {
		t.Error("Point Addition failed")
	}
}

func TestTupleNegation(t *testing.T) {
	a := Point(1, 2, 3)
	b := a.Minus()
	if b.X != -1 || b.Y != -2 || b.Z != -3 || b.W != 1 {
		t.Error("Point negation failed")
	}
	if !b.Point() {
		t.Error("Point negation failed")
	}
}

func TestTupleSubstraction(t *testing.T) {
	a := Point(1, 2, 3)
	b := Point(4, 5, 6)
	c := a.Sub(b)
	if c.X != -3 || c.Y != -3 || c.Z != -3 || c.W != 0 {
		t.Error("Point Substraction failed")
	}
	if !c.Vector() {
		t.Error("Point Substraction failed")
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	a := Point(1, 2, 3)
	b := Vector(4, 5, 6)
	c := a.Sub(b)
	if c.X != -3 || c.Y != -3 || c.Z != -3 || c.W != 1 {
		t.Error("Point Substraction failed")
	}
	if !c.Point() {
		t.Error("Result is not a Point")
	}
}

func TestVectorSubtraction(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(4, 5, 6)
	c := a.Sub(b)
	if c.X != -3 || c.Y != -3 || c.Z != -3 || c.W != 0 {
		t.Error("Vector Substraction failed")
	}
	if !c.Vector() {

		t.Error("Vector Substraction failed")
	}
}

func TestVectorMultiplicationByScalar(t *testing.T) {
	a := Vector(1, 2, 3)
	b := a.Scale(3)
	if b.X != 3 || b.Y != 6 || b.Z != 9 || b.W != 0 {
		t.Error("Vector multiplication by scalar failed")
	}
}

func TestVectorDivisionByScalar(t *testing.T) {
	a := Vector(1, 2, 3)
	b := a.Div(2)
	if b.X != 0.5 || b.Y != 1 || b.Z != 1.5 || b.W != 0 {
		t.Error("Vector division by scalar failed")
	}
}
