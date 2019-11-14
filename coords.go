package ray_tracer_challenge

type Coords struct {
	X float64
	Y float64
	Z float64
}

type Vector struct {
	coords Coords
	w      int8
}

type Point struct {
	coords Coords
	w      int8
}

func NewCoords(x, y, z float64) Coords {
	return Coords{x, y, z}
}

func NewVector(x, y, z float64) Vector {
	return Vector{
		coords: NewCoords(x, y, z),
		w:      0,
	}
}

func NewPoint(x, y, z float64) Point {
	return Point{
		coords: NewCoords(x, y, z),
		w:      1,
	}
}
