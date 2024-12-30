package ppm

import (
	"fmt"
	"github.com/zhou-en/ray-tracing-by-go/pkg/canvas"
	"github.com/zhou-en/ray-tracing-by-go/pkg/color"
	"os"
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
	//var lines []string
	//var line string
	//var nextLine string
	for _, p := range canvas.Pixels {
		p.Color.Scale()
		body = append(body, p.Color)
		//if nextLine != "" {
		//	line = nextLine
		//	nextLine = ""
		//}

		//for c := range []float64{p.Color.Red, p.Color.Green, p.Color.Blue} {
		//	line = fmt.Sprintf("%d ", c) //if len(line) < 69 {
		//	line = fmt.Sprintf("%s %d ", line, c)
		//} else {
		//	line += "\n"
		//	lines = append(lines, line)
		//	nextLine = fmt.Sprintf("%s %d ", nextLine, c)
		//}

		//}
	}
	//fmt.Println(lines)

	return PPM{
		Header: header,
		Body:   body,
	}
}

func (p *PPM) WritePPMToFile(ppm PPM, filename string) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write the header
	for _, line := range ppm.Header {
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			return fmt.Errorf("failed to write header: %w", err)
		}
	}

	// Write the body
	for i, c := range ppm.Body {
		r := int32(c.Red)
		g := int32(c.Green)
		b := int32(c.Blue)
		_, err := fmt.Fprintf(file, "%d %d %d", r, g, b)
		if err != nil {
			return fmt.Errorf("failed to write body: %w", err)
		}

		// Add a newline after every 23 pixels for better readability
		if (i+1)%23 == 0 {
			_, err := file.WriteString("\n")
			if err != nil {
				return fmt.Errorf("failed to write newline: %w", err)
			}
		} else {
			_, err := file.WriteString(" ")
			if err != nil {
				return fmt.Errorf("failed to write space: %w", err)
			}
		}
	}

	return nil
}
