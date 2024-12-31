package color

import (
	"math"
	"strings"
)

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func New(red, green, blue float64) Color {
	return Color{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func Scaler(num float64) float64 {
	scaledNum := math.Round(num * 255)
	return math.Min(math.Max(scaledNum, 0), 255)
}

func (c *Color) Scale() {
	c.Red = Scaler(c.Red)
	c.Green = Scaler(c.Green)
	c.Blue = Scaler(c.Blue)
}

func CreateRGBByName(name string) Color {
	switch strings.ToLower(name) {
	case "red":
		return Color{
			Red:   255,
			Green: 0,
			Blue:  0,
		}
	case "green":
		return Color{
			Red:   0,
			Green: 255,
			Blue:  0,
		}
	case "blue":
		return Color{
			Red:   0,
			Green: 0,
			Blue:  255,
		}
	default:
		return Color{
			Red:   0,
			Green: 0,
			Blue:  0,
		}
	}
}

func (c *Color) Add(color Color) Color {
	return Color{
		Red:   c.Red + color.Red,
		Green: c.Green + color.Green,
		Blue:  c.Blue + color.Blue,
	}
}

func (c *Color) Minus(color Color) Color {
	return Color{
		Red:   c.Red - color.Red,
		Green: c.Green - color.Green,
		Blue:  c.Blue - color.Blue,
	}
}

func (c *Color) Times(scalar float64) Color {
	return Color{
		Red:   c.Red * scalar,
		Green: c.Green * scalar,
		Blue:  c.Blue * scalar,
	}
}

func (c *Color) Blend(color Color) Color {
	return Color{
		Red:   c.Red * color.Red,
		Green: c.Green * color.Green,
		Blue:  c.Blue * color.Blue,
	}
}
