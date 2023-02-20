package vector

import t "github.com/zhou-en/ray-tracing-by-go/pkg/tuple"

func New(x float64, y float64, z float64) t.Tuple {
	return t.Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 0,
	}
}
