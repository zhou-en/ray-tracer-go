package vector_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/zhou-en/ray-tracing-by-go/pkg/vector"
)

var _ = Describe("New Vector", func() {
	It("NewVector() should create a vector with w=0", func() {
		a := vector.New(1.0, 2.0, 3.0)
		Expect(a.X).To(Equal(1.0))
		Expect(a.Y).To(Equal(2.0))
		Expect(a.Z).To(Equal(3.0))
		Expect(a.W).To(Equal(0.0))
	})
})

var _ = Describe("Add Vectors", func() {
	It("Add() 2 vectors should create a new tuple with w=0", func() {
		a := vector.New(1.0, 2.0, 3.0)
		b := vector.New(4.0, 5.0, 6.0)
		c := a.Add(b)
		Expect(c.X).To(Equal(5.0))
		Expect(c.Y).To(Equal(7.0))
		Expect(c.Z).To(Equal(9.0))
		Expect(c.W).To(Equal(0.0))
	})
	It("Add() 3 vectors should create a new tuple with w=0", func() {
		a := vector.New(1.0, 1.0, 1.0)
		b := vector.New(1.0, 1.0, 1.0)
		c := vector.New(1.0, 1.0, 1.0)
		d := a.Add(b).Add(c)
		Expect(d.X).To(Equal(3.0))
		Expect(d.Y).To(Equal(3.0))
		Expect(d.Z).To(Equal(3.0))
		Expect(d.W).To(Equal(0.0))
	})

})

var _ = Describe("Vector Subtraction", func() {
	It("Subtract() 2 vectors should create a new tuple with w=0", func() {
		a := vector.New(1.0, 2.0, 3.0)
		b := vector.New(4.0, 5.0, 6.0)
		c := a.Sub(b)
		Expect(c.X).To(Equal(-3.0))
		Expect(c.Y).To(Equal(-3.0))
		Expect(c.Z).To(Equal(-3.0))
		Expect(c.W).To(Equal(0.0))
	})
})

var _ = Describe("Negating a vector", func() {
	It("Negating a vector should create a new vector with components that have opositin signs", func() {
		v := vector.New(1.0, 2.0, 3.0)
		v.Neg()
		Expect(v.X).To(Equal(-1.0))
		Expect(v.Y).To(Equal(-2.0))
		Expect(v.Z).To(Equal(-3.0))
		Expect(v.W).To(Equal(0.0))
	})
})

var _ = Describe("Computing the magnitude of a vector", func() {
	It("Computing the magnitude of a vector(1,0,0) should return 1", func() {
		v := vector.New(1.0, 0.0, 0.0)
		var m float64 = v.Mag()
		Expect(m).To(Equal(1.0))
	})
	It("Computing the magnitude of a vector(0,1,0) should return 1", func() {
		v := vector.New(0.0, 1.0, 0.0)
		var m float64 = v.Mag()
		Expect(m).To(Equal(1.0))
	})
	It("Computing the magnitude of a vector(0,0,1) should return 1", func() {
		v := vector.New(0.0, 0.0, 1.0)
		var m float64 = v.Mag()
		Expect(m).To(Equal(1.0))
	})
	It("Computing the magnitude of a vector(2,3,4) should return Sqrt(6.0)", func() {
		v := vector.New(2.0, 1.0, 1.0)
		var m float64 = v.Mag()
		Expect(m).To(Equal(math.Sqrt(6.0)))
	})

})

var _ = Describe("Normalize a vector", func() {
	It("Given a vector (4, 0, 0), Norm() should return a vector (1, 0, 0)", func() {
		v := vector.New(4.0, 0, 0)
		nv := v.Norm()
		Expect(nv.X).To(Equal(1.0))
		Expect(nv.Y).To(Equal(0.0))
		Expect(nv.Z).To(Equal(0.0))
		Expect(nv.W).To(Equal(v.W))
	})
	It("Given a vector (1, 2, 3), Norm() should return a vector (1/√14, 2/√14, 3/√14)", func() {
		v := vector.New(1.0, 2.0, 3.0)
		nv := v.Norm()
		Expect(nv.X).To(Equal(1.0 / math.Sqrt(14)))
		Expect(nv.Y).To(Equal(2.0 / math.Sqrt(14)))
		Expect(nv.Z).To(Equal(3.0 / math.Sqrt(14)))
		Expect(nv.W).To(Equal(v.W))
	})
	It("Given a vector (1, 2, 3), Norm() should return a vector (1/√14, 2/√14, 3/√14)", func() {
		v := vector.New(1.0, 2.0, 3.0)
		nv := v.Norm()
		Expect(nv.Mag()).To(Equal(1.0))
	})
})

var _ = Describe("Dot Product", func() {
	It("Given vector a (1, 2, 3) and b (2, 3, 4), dot product of a and b gives 20", func() {
		a := vector.New(1.0, 2.0, 3.0)
		b := vector.New(2.0, 3.0, 4.0)
		p := a.Dot(b)
		Expect(p).To(Equal(20.0))
	})

})

var _ = Describe("Cross Product", func() {
	It("Given vector a (1, 2, 3) and b (2, 3, 4), cross product of a X b gives a vector (-1, 2, -1)", func() {
		a := vector.New(1.0, 2.0, 3.0)
		b := vector.New(2.0, 3.0, 4.0)
		p := a.Cross(b)
		Expect(p.X).To(Equal(-1.0))
		Expect(p.Y).To(Equal(2.0))
		Expect(p.Z).To(Equal(-1.0))
		Expect(p.W).To(Equal(a.W))
		Expect(p.W).To(Equal(b.W))
	})
	It("Given vector a (1, 2, 3) and b (2, 3, 4), cross product of b X a gives a vector (1, -2, 1)", func() {
		a := vector.New(1.0, 2.0, 3.0)
		b := vector.New(2.0, 3.0, 4.0)
		p := b.Cross(a)
		Expect(p.X).To(Equal(1.0))
		Expect(p.Y).To(Equal(-2.0))
		Expect(p.Z).To(Equal(1.0))
		Expect(p.W).To(Equal(a.W))
		Expect(p.W).To(Equal(b.W))
	})
})
