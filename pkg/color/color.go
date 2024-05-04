package color

import "strings"

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func New(red, green, blue float64) Color {
	if red > 255 {
		red = float64(255)
	} else if red < 0 {
		red = float64(0)
	}

	if green > 255 {
		green = float64(255)
	} else if green < 0 {
		green = float64(0)
	}

	if blue > 255 {
		blue = float64(255)
	} else if blue < 0 {
		blue = float64(0)
	}

	return Color{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
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
