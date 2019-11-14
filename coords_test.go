package ray_tracer_challenge

import (
	"errors"
	"github.com/mathew/ray-tracer-challenge/asserts"
	"testing"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(1, 2, 3)

	asserts.Equals(t, 1.0, p.X)
	asserts.Equals(t, 2.0, p.Y)
	asserts.Equals(t, 3.0, p.Z)
	asserts.Equals(t, int8(1), p.W)
}

func TestNewVector(t *testing.T) {
	p := NewVector(1, 2, 3)

	asserts.Equals(t, 1.0, p.X)
	asserts.Equals(t, 2.0, p.Y)
	asserts.Equals(t, 3.0, p.Z)
	asserts.Equals(t, int8(0), p.W)
}

func TestCompareVectors(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(1, 2, 3)
	c := NewVector(1, 2, 4)

	asserts.Assert(t, a == b, "A does not equal B")
	asserts.Assert(t, a != c, "A does equal C")
}

func TestAddCoords(t *testing.T) {
	v := NewVector(1, 2, 3)
	p := NewPoint(1, 2, 3)

	add, _ := v.add(p)

	asserts.Equals(t, 2.0, add.X)
	asserts.Equals(t, 4.0, add.Y)
	asserts.Equals(t, 6.0, add.Z)

	_, err := p.add(p)
	asserts.Assert(t, errors.Is(err, TwoPointsError), "TwoPointsError not raised")
}
