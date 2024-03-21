package vector

import (
	"math"

	t "github.com/zhou-en/ray-tracing-by-go/pkg/tuple"
)

type Vector struct {
	t.Tuple
}

func New(x float64, y float64, z float64) Vector {
	return Vector{
		t.Tuple{
			X: x,
			Y: y,
			Z: z,
			W: 0,
		},
	}
}

func (v *Vector) AddVector(v2 Vector) Vector {
	return New(
		v.X+v2.X,
		v.Y+v2.Y,
		v.Z+v2.Z,
	)
}

func AddVectors(v ...Vector) Vector {
	var resultVector Vector
	for _, vector := range v {
		resultVector = resultVector.AddVector(vector)
	}
	return resultVector
}

func (v *Vector) SubVector(v2 Vector) Vector {
	return New(
		v.X-v2.X,
		v.Y-v2.Y,
		v.Z-v2.Z,
	)
}

func (v *Vector) NegVector() {
	v.X *= -1
	v.Y *= -1
	v.Z *= -1
}

func (v *Vector) DotVector(b Vector) float64 {
	var s float64 = 0
	s += v.X * b.X
	s += v.Y * b.Y
	s += v.Z * b.Z
	return s
}

func (v *Vector) MagVector() float64 {
	return math.Sqrt(v.DotVector(*v))
}

func (v *Vector) NormVector() Vector {
	var m = v.MagVector()
	return New(v.X/m, v.Y/m, v.Z/m)
}

func (v *Vector) CrossVector(b Vector) Vector {
	return New(
		v.Y*b.Z-v.Z*b.Y,
		v.Z*b.X-v.X*b.Z,
		v.X*b.Y-v.Y*b.X,
	)
}
