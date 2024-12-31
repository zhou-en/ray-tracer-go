package canvas

import (
	"github.com/zhou-en/ray-tracing-by-go/pkg/color"
)

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

func (p *Pixel) UpdateColor(newColor color.Color) {
	p.Color = newColor
}

type Canvas struct {
	Width  int
	Height int
	Pixels []Pixel
}

func New(width, height int) Canvas {
	pixels := make([]Pixel, width*height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			index := i*height + j
			pixels[index] = Pixel{
				X:     i,
				Y:     j,
				Color: color.New(0, 0, 0),
			}
		}
	}
	return Canvas{
		Width:  width,
		Height: height,
		Pixels: pixels,
	}
}

func (c *Canvas) GetPixel(x, y int) (int, *Pixel) {
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height {
		return -1, nil
	}
	for i := range c.Pixels {
		if c.Pixels[i].X == x && c.Pixels[i].Y == y {
			return i, &c.Pixels[i]
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

// SetColor sets all pixels' color to the given one
func (c *Canvas) SetColor(newColor color.Color) {
	for i := range c.Pixels {
		c.Pixels[i].UpdateColor(newColor)
	}
}
