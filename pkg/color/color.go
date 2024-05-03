package color

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
