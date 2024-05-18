package color_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zhou-en/ray-tracing-by-go/pkg"
	"github.com/zhou-en/ray-tracing-by-go/pkg/color"
)

const tolerance = 0.0000001

var _ = Describe("Color Assignment", func() {
	It("Color are red green blue tuple", func() {
		c := color.New(-0.5, 0.4, 265)
		Expect(c.Red).To(Equal(-0.5))
		Expect(c.Green).To(Equal(0.4))
		Expect(c.Blue).To(Equal(float64(265)))
	})
})

var _ = Describe("Colors Operations", func() {
	It("Add another color to another", func() {
		c1 := color.New(0.9, 0.6, 0.75)
		c2 := color.New(0.7, 0.1, 0.25)
		c := c1.Add(c2)
		Expect(c.Red).To(Equal(1.6))
		Expect(c.Green).To(Equal(0.7))
		Expect(c.Blue).To(Equal(1.0))
	})

	It("Subtract a color from another", func() {
		c1 := color.New(0.9, 0.6, 0.75)
		c2 := color.New(0.7, 0.1, 0.25)
		c := c1.Minus(c2)
		Expect(c.Red).To(pkg.BeCloseTo(0.2, tolerance))
		Expect(c.Green).To(pkg.BeCloseTo(0.5, tolerance))
		Expect(c.Blue).To(pkg.BeCloseTo(0.5, tolerance))
	})

	It("Multiple by a scalar", func() {
		c := color.New(0.9, 0.6, 0.75)
		c = c.Times(1.6)
		Expect(c.Red).To(pkg.BeCloseTo(1.44, tolerance))
		Expect(c.Green).To(pkg.BeCloseTo(0.96, tolerance))
		Expect(c.Blue).To(pkg.BeCloseTo(1.2, tolerance))
	})

	It("Blend a color from another", func() {
		c1 := color.New(0.9, 0.6, 0.75)
		c2 := color.New(0.7, 0.1, 0.25)
		c := c1.Blend(c2)
		Expect(c.Red).To(pkg.BeCloseTo(0.63, tolerance))
		Expect(c.Green).To(pkg.BeCloseTo(0.06, tolerance))
		Expect(c.Blue).To(pkg.BeCloseTo(0.1875, tolerance))
	})

	// Scaling table tests
	var (
		testCases = []struct {
			Color         color.Color
			expectedColor color.Color
		}{
			{
				Color:         color.New(1.5, 0, 0),
				expectedColor: color.New(255, 0, 0),
			},
			{
				Color:         color.New(0, 0.5, 0),
				expectedColor: color.New(0, 128, 0),
			},
			{
				Color:         color.New(-0.5, 0, 1),
				expectedColor: color.New(0, 0, 255),
			},
		}
	)

	It("Scale a color", func() {
		for _, tc := range testCases {
			tc.Color.Scale()
			Expect(tc.Color.Red).To(Equal(tc.expectedColor.Red))
			Expect(tc.Color.Green).To(Equal(tc.expectedColor.Green))
			Expect(tc.Color.Blue).To(Equal(tc.expectedColor.Blue))
		}
	})

})
