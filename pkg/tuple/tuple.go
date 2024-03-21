package tuple

import "math"

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func (a Tuple) Add(b Tuple) Tuple {
	return Tuple{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

func (a Tuple) Sub(b Tuple) Tuple {
	return Tuple{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

func (a *Tuple) Neg() {
	a.X *= -1
	a.Y *= -1
	a.Z *= -1
}

func (a *Tuple) Prod(b float64) {
	a.X *= b
	a.Y *= b
	a.Z *= b
	a.W *= b
}

func (a *Tuple) Div(b float64) {
	a.X /= b
	a.Y /= b
	a.Z /= b
	a.W /= b
}

func (a Tuple) Dot(b Tuple) float64 {
	var s float64 = 0
	s += a.X * b.X
	s += a.Y * b.Y
	s += a.Z * b.Z
	return s
}

func (a Tuple) Mag() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Tuple) Norm() Tuple {
	var m float64 = a.Mag()
	return Tuple{
		X: a.X / m,
		Y: a.Y / m,
		Z: a.Z / m,
		W: a.W,
	}
}

func (a *Tuple) Cross(b Tuple) Tuple {
	return Tuple{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
		W: 0,
	}
}

func NewPoint(x float64, y float64, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 1.0,
	}
}
