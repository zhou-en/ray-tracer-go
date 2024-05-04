package canvas

import "github.com/zhou-en/ray-tracing-by-go/pkg/color"

type Pixel struct {
	X     int
	Y     int
	Color color.Color
}

func NewPixel(x, y int, color color.Color) Pixel {
	return Pixel{
		X:     x,
		Y:     y,
		Color: color,
	}
}

type Canvas struct {
	Width  int
	Height int
	Pixels []Pixel
}

func New(width, height int) Canvas {
	var pixels []Pixel

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			pixels = append(pixels, Pixel{
				X:     i,
				Y:     j,
				Color: color.New(0, 0, 0),
			})
		}
	}

	return Canvas{
		Width:  width,
		Height: height,
		Pixels: pixels,
	}
}

func (c *Canvas) GetPixel(x, y int) (int, *Pixel) {
	for i, p := range c.Pixels {
		if p.X == x && p.Y == y {
			return i, &p
		}
	}
	return -1, nil
}

func (c *Canvas) AddPixel(x, y int, color color.Color) bool {
	i, p := c.GetPixel(x, y)
	if p != nil {
		c.Pixels[i] = NewPixel(x, y, color)
		return true
	}
	return false
}
