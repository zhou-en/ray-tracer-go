package ppm

import (
	"fmt"
	"github.com/zhou-en/ray-tracing-by-go/pkg/canvas"
	"github.com/zhou-en/ray-tracing-by-go/pkg/color"
)

type PPM struct {
	Header []string
	Body   []color.Color
}

func New(canvas canvas.Canvas) PPM {
	header := []string{
		"P3",
		fmt.Sprintf("%d %d", canvas.Width, canvas.Height),
		"255",
	}

	var body []color.Color
	var lines []string
	var line string
	var nextLine string
	for _, p := range canvas.Pixels {
		p.Color.Scale()
		body = append(body, p.Color)
		if nextLine != "" {
			line = nextLine
			nextLine = ""
		}

		for c := range []float64{p.Color.Red, p.Color.Green, p.Color.Blue} {

			if len(line) < 69 {
				line = fmt.Sprintf("%s %d ", line, c)
			} else {
				line += "\n"
				lines = append(lines, line)
				nextLine = fmt.Sprintf("%s %d ", nextLine, c)
			}

		}
	}
	fmt.Println(lines)

	return PPM{
		Header: header,
		Body:   body,
	}
}

func (p *PPM) WriteToFile() {

}
