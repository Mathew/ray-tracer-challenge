package ray_tracer_challenge

import "errors"

type Coord struct {
	X float64
	Y float64
	Z float64
	W int8
}


var TwoPointsError = errors.New("cannot add two points")

func (c Coord) add (other Coord) (Coord, error) {
	if c.W + other.W == 2 {
		return Coord{}, TwoPointsError
	}

	return NewCoord(c.X + other.X, c.Y + other.Y, c.Z + other.Z, c.W + other.W), nil
}

func NewCoord(x, y, z float64, w int8) Coord {
	return Coord{x, y, z, w}
}

func NewVector(x, y, z float64) Coord {
	return NewCoord(x, y, z, 0)
}

func NewPoint(x, y, z float64) Coord {
	return NewCoord(x, y, z, 1)
}
