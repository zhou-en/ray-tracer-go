package tuple_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zhou-en/ray-tracing-by-go/pkg/tuple"
)

var _ = Describe("Subtracting a tuple from another tuple", func() {
	It("Subtracting a tuple from another tuple should create a new tuple", func() {
		a := tuple.New(1.0, 2.0, 3.0, 1.0)
		b := tuple.New(4.0, 5.0, 6.0, 0)
		c := a.Sub(b)
		Expect(c.X).To(Equal(-3.0))
		Expect(c.Y).To(Equal(-3.0))
		Expect(c.Z).To(Equal(-3.0))
		Expect(c.W).To(Equal(1.0))
	})
})

var _ = Describe("Subtracting a vector from the zero vector (0, 0, 0)", func() {
	It("Subtracting a vector from the zero vector should create a new vector with w=0", func() {
		v := tuple.NewVector(1.0, 2.0, 3.0)
		z := tuple.NewVector(0, 0, 0)
		c := z.SubVector(v)
		Expect(c.X).To(Equal(-1.0))
		Expect(c.Y).To(Equal(-2.0))
		Expect(c.Z).To(Equal(-3.0))
		Expect(c.W).To(Equal(0.0))
	})
})

var _ = Describe("Multiply a Tuple by a scalar", func() {
	It("Tuple multiplies a scalar should create a new tuple with components that are muliplied by the scalar", func() {
		var s = 2.0
		v := tuple.Tuple{X: 1.0, Y: 2.0, Z: 3.0, W: 4.0}
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
		v := tuple.Tuple{X: 1.0, Y: 2.0, Z: 3.0, W: 4.0}
		v.Div(s)
		Expect(v.X).To(Equal(0.5))
		Expect(v.Y).To(Equal(1.0))
		Expect(v.Z).To(Equal(1.5))
		Expect(v.W).To(Equal(2.0))
	})
})
