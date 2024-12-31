package tuple

import (
	"math"
)

// Tuple represents a 4D tuple.
type Tuple struct {
	X, Y, Z, W float64
}

// NewTuple creates a new Tuple.
func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
}

// Add adds two tuples and returns the result.
func (a Tuple) Add(b Tuple) Tuple {
	return Tuple{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

// Sub subtracts two tuples and returns the result.
func (a Tuple) Sub(b Tuple) Tuple {
	return Tuple{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

// Negate returns the negated tuple.
func (a Tuple) Negate() Tuple {
	return Tuple{
		X: -a.X,
		Y: -a.Y,
		Z: -a.Z,
		W: -a.W,
	}
}

// Scale scales the tuple by a scalar and returns the result.
func (a Tuple) Scale(factor float64) Tuple {
	return Tuple{
		X: a.X * factor,
		Y: a.Y * factor,
		Z: a.Z * factor,
		W: a.W * factor,
	}
}

// Divide divides the tuple by a scalar and returns the result.
func (a Tuple) Divide(factor float64) Tuple {
	return Tuple{
		X: a.X / factor,
		Y: a.Y / factor,
		Z: a.Z / factor,
		W: a.W / factor,
	}
}

// Point represents a point in 3D space.
type Point Tuple

// NewPoint creates a new Point.
func NewPoint(x, y, z float64) Point {
	return Point(NewTuple(x, y, z, 1.0))
}

// AddVector adds a vector to a point and returns a new point.
func (p Point) AddVector(v Vector) Point {
	t := Tuple(p).Add(Tuple(v))
	return Point{X: t.X, Y: t.Y, Z: t.Z, W: t.W}
}

// Vector represents a vector in 3D space.
type Vector Tuple

// NewVector creates a new Vector.
func NewVector(x, y, z float64) Vector {
	return Vector(NewTuple(x, y, z, 0.0))
}

// Dot calculates the dot product of two vectors.
func (v Vector) Dot(b Vector) float64 {
	return v.X*b.X + v.Y*b.Y + v.Z*b.Z
}

// Magnitude calculates the magnitude of the vector.
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize returns the normalized vector.
func (v Vector) Normalize() Vector {
	mag := v.Magnitude()
	return v.Scale(1 / mag)
}

// Cross calculates the cross product of two vectors.
func (v Vector) Cross(b Vector) Vector {
	return NewVector(
		v.Y*b.Z-v.Z*b.Y,
		v.Z*b.X-v.X*b.Z,
		v.X*b.Y-v.Y*b.X,
	)
}

// Add adds two vectors and returns the result.
func (v Vector) Add(b Vector) Vector {
	return Vector(Tuple(v).Add(Tuple(b)))
}

// Sub subtracts two vectors and returns the result.
func (v Vector) Sub(b Vector) Vector {
	return Vector(Tuple(v).Sub(Tuple(b)))
}

// Scale scales the vector by a scalar and returns the result.
func (v Vector) Scale(factor float64) Vector {
	return Vector(Tuple(v).Scale(factor))
}
