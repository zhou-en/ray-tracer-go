package tuple

import "math"

/*
	Tuple
*/

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func New(x float64, y float64, z float64, w float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func (a *Tuple) AddTuple(b Tuple) Tuple {
	return Tuple{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

func (a *Tuple) Add(v ...Tuple) Tuple {
	var result Tuple
	for _, t := range v {
		result = result.AddTuple(t)
	}
	return result
}

func (a *Tuple) SubTuple(b Tuple) Tuple {
	return Tuple{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

func (a *Tuple) Sub(v ...Tuple) Tuple {
	var result Tuple
	for _, t := range v {
		result = result.SubTuple(t)
	}
	return result
}

func (a *Tuple) Negate() {
	a.X *= -1
	a.Y *= -1
	a.Z *= -1
	a.W *= -1
}

func (a *Tuple) Prod(b float64) {
	a.X = a.X * b
	a.Y = a.Y * b
	a.Z = a.Z * b
	a.W = a.W * b
}

func (a *Tuple) Div(b float64) {
	a.X = a.X / b
	a.Y = a.Y / b
	a.Z = a.Z / b
	a.W = a.W / b
}

/*
	Point
*/

type Point struct {
	Tuple
}

func NewPoint(x float64, y float64, z float64) Point {
	return Point{
		Tuple{
			X: x,
			Y: y,
			Z: z,
			W: 1.0,
		},

		//New(x, y, z, 1.0),
	}
}

func (a *Point) AddVector(b Vector) Point {
	return NewPoint(
		a.X+b.X,
		a.Y+b.Y,
		a.Z+b.Z,
	)
}

/*
	Vector
*/

type Vector struct {
	Tuple
}

func NewVector(x float64, y float64, z float64) Vector {
	return Vector{
		New(
			x,
			y,
			z,
			0,
		),
	}
}
func (v *Vector) Dot(b Vector) float64 {
	var s float64 = 0
	s += v.X * b.X
	s += v.Y * b.Y
	s += v.Z * b.Z
	return s
}

func (v *Vector) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector) Norm() Vector {
	var m float64 = v.Mag()
	return NewVector(
		v.X/m,
		v.Y/m,
		v.Z/m,
	)
}

func (v *Vector) Cross(b Vector) Vector {
	return NewVector(
		v.Y*b.Z-v.Z*b.Y,
		v.Z*b.X-v.X*b.Z,
		v.X*b.Y-v.Y*b.X,
	)
}

func (a *Vector) SubVector(b Vector) Vector {
	return NewVector(
		a.X-b.X,
		a.Y-b.Y,
		a.Z-b.Z,
	)
}

func (a *Vector) AddVector(b Vector) Vector {
	return NewVector(
		a.X+b.X,
		a.Y+b.Y,
		a.Z+b.Z,
	)
}
