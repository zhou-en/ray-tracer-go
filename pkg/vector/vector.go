package vector

import t "github.com/zhou-en/ray-tracing-by-go/pkg/tuple"

type Operation interface {
	Add(v Vector) Vector
}

type Vector struct {
	t.Tuple
	X float64
}

func New(x float64, y float64, z float64) t.Tuple {
	return t.Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 0,
	}
}

func (v *Vector) AddVector(v2 Vector) Vector {
	return Vector{
		Tuple: t.Tuple{
			X: v.X + v2.X,
			Y: v.Y + v2.Y,
			Z: v.Z + v2.Z,
			W: 0,
		},
	}
}
