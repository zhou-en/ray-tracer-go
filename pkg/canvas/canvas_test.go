package canvas_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zhou-en/ray-tracing-by-go/pkg/canvas"
	"github.com/zhou-en/ray-tracing-by-go/pkg/color"
)

const tolerance = 0.0000001

var _ = Describe("Canvas Creation", func() {
	It("Creating a canvas", func() {
		width := 10
		height := 20
		black := color.New(0, 0, 0)
		c := canvas.New(width, height)
		for _, p := range c.Pixels {
			Expect(p.Color).To(Equal(black))
			Expect(p.X).To(BeNumerically("<", width))
			Expect(p.Y).To(BeNumerically("<", height))
		}
		Expect(c.Width).To(Equal(width))
		Expect(c.Height).To(Equal(height))
	})

	It("Add a pixel to a canvas", func() {
		width := 10
		height := 20
		red := color.New(255, 0, 0)
		c := canvas.New(width, height)
		c.AddPixel(2, 3, red)
		_, newPixel := c.GetPixel(2, 3)
		Expect(newPixel.Color).To(Equal(red))
		Expect(newPixel.X).To(Equal(2))
		Expect(newPixel.Y).To(Equal(3))
	})

	It("Set color for all pixels", func() {
		width := 2
		height := 3
		red := color.New(255, 0, 0)
		c := canvas.New(width, height)
		c.SetColor(red)
		for _, p := range c.Pixels {
			Expect(p.Color).To(Equal(red))
		}
	})
})
