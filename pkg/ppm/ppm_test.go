package ppm_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zhou-en/ray-tracing-by-go/pkg/canvas"
	"github.com/zhou-en/ray-tracing-by-go/pkg/color"
	"github.com/zhou-en/ray-tracing-by-go/pkg/ppm"
)

var _ = Describe("PPM Initialization", func() {
	It("PPM is created with expected pixels and colors", func() {
		c := canvas.New(5, 3)
		c1 := color.New(1.5, 0, 0)
		c2 := color.New(0, 0.5, 0)
		c3 := color.New(-0.5, 0, 1)
		c.AddPixel(0, 0, c1)
		c.AddPixel(2, 1, c2)
		c.AddPixel(4, 2, c3)
		newPpm := ppm.New(c)

		c1.Scale()
		Expect(newPpm.Body[0].Red).To(Equal(c1.Red))
		Expect(newPpm.Body[0].Green).To(Equal(c1.Green))
		Expect(newPpm.Body[0].Blue).To(Equal(c1.Blue))

		Expect(newPpm.Body[1].Red).To(Equal(float64(0)))
		Expect(newPpm.Body[1].Green).To(Equal(float64(0)))
		Expect(newPpm.Body[1].Blue).To(Equal(float64(0)))

		i, _ := c.GetPixel(2, 1)
		c2.Scale()
		Expect(newPpm.Body[i].Red).To(Equal(c2.Red))
		Expect(newPpm.Body[i].Green).To(Equal(c2.Green))
		Expect(newPpm.Body[i].Blue).To(Equal(c2.Blue))
	})
})
