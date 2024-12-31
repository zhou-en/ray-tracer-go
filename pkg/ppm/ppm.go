package ppm

import (
	"bufio"
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
	for i := range canvas.Pixels {
		canvas.Pixels[i].Color.Scale()
		body = append(body, canvas.Pixels[i].Color)
	}
	return PPM{
		Header: header,
		Body:   body,
	}
}

func (p *PPM) WritePPMToFile(ppm PPM, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Write the header
	for _, line := range ppm.Header {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return fmt.Errorf("failed to write header: %w", err)
		}
	}

	// Write the body
	for i, c := range ppm.Body {
		r, g, b := int32(c.Red), int32(c.Green), int32(c.Blue)
		if (i+1)%23 == 0 {
			_, err := fmt.Fprintf(writer, "%d %d %d\n", r, g, b)
			if err != nil {
				return fmt.Errorf("failed to write body: %w", err)
			}
		} else {
			_, err := fmt.Fprintf(writer, "%d %d %d ", r, g, b)
			if err != nil {
				return fmt.Errorf("failed to write body: %w", err)
			}
		}
	}

	return nil
}
