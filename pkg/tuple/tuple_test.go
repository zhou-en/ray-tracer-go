package tuple_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/zhou-en/ray-tracing-by-go/pkg/tuple"
	"github.com/zhou-en/ray-tracing-by-go/pkg/vector"
)

var _ = Describe("Tuple", func() {
	It("NewPoint() should create a point with w=1", func() {
		a := tuple.NewPoint(1.0, 2.0, 3.0)
		Expect(a.X).To(Equal(1.0))
		Expect(a.Y).To(Equal(2.0))
		Expect(a.Z).To(Equal(3.0))
		Expect(a.W).To(Equal(1.0))
	})
})

var _ = Describe("Add 2 Points", func() {
	It("Add() 2 points should create a new tuple with w=2", func() {
		a := tuple.NewPoint(1.0, 2.0, 3.0)
		b := tuple.NewPoint(4.0, 5.0, 6.0)
		c := a.Add(b)
		Expect(c.X).To(Equal(5.0))
		Expect(c.Y).To(Equal(7.0))
		Expect(c.Z).To(Equal(9.0))
		Expect(c.W).To(Equal(2.0))
	})
})

var _ = Describe("Points Subtraction", func() {
	It("Subtract() 2 points should create a new tuple with w=0", func() {
		a := tuple.NewPoint(1.0, 2.0, 3.0)
		b := tuple.NewPoint(4.0, 5.0, 6.0)
		c := a.Sub(b)
		Expect(c.X).To(Equal(-3.0))
		Expect(c.Y).To(Equal(-3.0))
		Expect(c.Z).To(Equal(-3.0))
		Expect(c.W).To(Equal(0.0))
	})
})

var _ = Describe("Subtracting a vector from a point", func() {
	It("Subtracting a vector from a point should create a new tuple with w=1", func() {
		p := tuple.NewPoint(1.0, 2.0, 3.0)
		v := vector.New(4.0, 5.0, 6.0)
		c := p.Sub(v)
		Expect(c.X).To(Equal(-3.0))
		Expect(c.Y).To(Equal(-3.0))
		Expect(c.Z).To(Equal(-3.0))
		Expect(c.W).To(Equal(1.0))
	})
})

var _ = Describe("Subtracting a vector from the zero vector (0, 0, 0)", func() {
	It("Subtracting a vector from the zero vector should create a new vector with w=0", func() {
		v := vector.New(1.0, 2.0, 3.0)
		z := vector.New(0, 0, 0)
		c := z.Sub(v)
		Expect(c.X).To(Equal(-1.0))
		Expect(c.Y).To(Equal(-2.0))
		Expect(c.Z).To(Equal(-3.0))
		Expect(c.W).To(Equal(0.0))
	})
})

var _ = Describe("Multiply a Tuple by a scalar", func() {
	It("Tuple multiplies a scalar should create a new tuple with components that are muliplied by the scalar", func() {
		var s float64 = 2.0
		v := tuple.Tuple{1.0, 2.0, 3.0, 4.0}
		v.Prod(s)
		Expect(v.X).To(Equal(2.0))
		Expect(v.Y).To(Equal(4.0))
		Expect(v.Z).To(Equal(6.0))
		Expect(v.W).To(Equal(8.0))
	})
})

var _ = Describe("Divide a Tuple by a scalar", func() {
	It("Divide a Tuple by a scalar should update the components of the Tuple by dividing the scalar", func() {
		var s float64 = 2.0
		v := tuple.Tuple{1.0, 2.0, 3.0, 4.0}
		v.Div(s)
		Expect(v.X).To(Equal(0.5))
		Expect(v.Y).To(Equal(1.0))
		Expect(v.Z).To(Equal(1.5))
		Expect(v.W).To(Equal(2.0))
	})
})

var _ = Describe("Colors", func() {
	It("Color are red green blue tuple", func() {
		v := tuple.Color{-0.5, 0.4, 1.7}
		Expect(v.Red).To(Equal(-0.5))
		Expect(v.Green).To(Equal(0.4))
		Expect(v.Blue).To(Equal(1.7))
	})
})
