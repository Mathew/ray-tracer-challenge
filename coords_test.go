package ray_tracer_challenge

import (
	"github.com/mathew/ray-tracer-challenge/asserts"
	"testing"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(1, 2, 3)

	asserts.Equals(t, 1.0, p.coords.X)
	asserts.Equals(t, 2.0, p.coords.Y)
	asserts.Equals(t, 3.0, p.coords.Z)
	asserts.Equals(t, int8(1), p.w)
}

func TestNewVector(t *testing.T) {
	p := NewVector(1, 2, 3)

	asserts.Equals(t, 1.0, p.coords.X)
	asserts.Equals(t, 2.0, p.coords.Y)
	asserts.Equals(t, 3.0, p.coords.Z)
	asserts.Equals(t, int8(0), p.w)
}
