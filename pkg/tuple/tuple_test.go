package tuple_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zhou-en/ray-tracing-by-go/pkg/tuple"
	"math"
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

var _ = Describe("Tuple", func() {
	It("NewVector() should create a vector with w=0", func() {
		a := tuple.NewVector(1.0, 2.0, 3.0)
		Expect(a.X).To(Equal(1.0))
		Expect(a.Y).To(Equal(2.0))
		Expect(a.Z).To(Equal(3.0))
		Expect(a.W).To(Equal(0.0))
	})
})

var _ = Describe("Add Vectors", func() {
	It("Add() 2 vectors should create a new tuple with w=0", func() {
		a := tuple.NewVector(1.0, 2.0, 3.0)
		b := tuple.NewVector(4.0, 5.0, 6.0)
		c := a.Add(b)
		Expect(c.X).To(Equal(5.0))
		Expect(c.Y).To(Equal(7.0))
		Expect(c.Z).To(Equal(9.0))
		Expect(c.W).To(Equal(0.0))
	})
	It("Add() 3 vectors should create a new tuple with w=0", func() {
		a := tuple.NewVector(1.0, 1.0, 1.0)
		b := tuple.NewVector(1.0, 1.0, 1.0)
		c := tuple.NewVector(1.0, 1.0, 1.0)
		d := a.Add(b).Add(c)
		Expect(d.X).To(Equal(3.0))
		Expect(d.Y).To(Equal(3.0))
		Expect(d.Z).To(Equal(3.0))
		Expect(d.W).To(Equal(0.0))
	})

})

var _ = Describe("Vector Subtraction", func() {
	It("Subtract() 2 vectors should create a new tuple with w=0", func() {
		a := tuple.NewVector(1.0, 2.0, 3.0)
		b := tuple.NewVector(4.0, 5.0, 6.0)
		c := a.Sub(b)
		Expect(c.X).To(Equal(-3.0))
		Expect(c.Y).To(Equal(-3.0))
		Expect(c.Z).To(Equal(-3.0))
		Expect(c.W).To(Equal(0.0))
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
		v := tuple.NewVector(4.0, 5.0, 6.0)
		c := p.Sub(v)
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
		c := z.Sub(v)
		Expect(c.X).To(Equal(-1.0))
		Expect(c.Y).To(Equal(-2.0))
		Expect(c.Z).To(Equal(-3.0))
		Expect(c.W).To(Equal(0.0))
	})
})

var _ = Describe("Negating a vector", func() {
	It("Negating a vector should create a new vector with components that have opositin signs", func() {
		v := tuple.NewVector(1.0, 2.0, 3.0)
		v.Neg()
		Expect(v.X).To(Equal(-1.0))
		Expect(v.Y).To(Equal(-2.0))
		Expect(v.Z).To(Equal(-3.0))
		Expect(v.W).To(Equal(0.0))
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

var _ = Describe("Computing the magnitude of a vector", func() {
	It("Computing the magnitude of a vector(1,0,0) should return 1", func() {
		v := tuple.NewVector(1.0, 0.0, 0.0)
		var m float64 = v.Mag()
		Expect(m).To(Equal(1.0))
	})
	It("Computing the magnitude of a vector(0,1,0) should return 1", func() {
		v := tuple.NewVector(0.0, 1.0, 0.0)
		var m float64 = v.Mag()
		Expect(m).To(Equal(1.0))
	})
	It("Computing the magnitude of a vector(0,0,1) should return 1", func() {
		v := tuple.NewVector(0.0, 0.0, 1.0)
		var m float64 = v.Mag()
		Expect(m).To(Equal(1.0))
	})
	It("Computing the magnitude of a vector(2,3,4) should return Sqrt(6.0)", func() {
		v := tuple.NewVector(2.0, 1.0, 1.0)
		var m float64 = v.Mag()
		Expect(m).To(Equal(math.Sqrt(6.0)))
	})

})

var _ = Describe("Normalize a vector", func() {
	It("Given a vector (4, 0, 0), Normal() should return a vector (1, 0, 0)", func() {
		v := tuple.NewVector(4.0, 0, 0)
		nv := v.Normal()
		Expect(nv.X).To(Equal(1.0))
		Expect(nv.Y).To(Equal(0.0))
		Expect(nv.Z).To(Equal(0.0))
		Expect(nv.W).To(Equal(v.W))
	})
	It("Given a vector (1, 2, 3), Normal() should return a vector (1/√14, 2/√14, 3/√14)", func() {
		v := tuple.NewVector(1.0, 2.0, 3.0)
		nv := v.Normal()
		Expect(nv.X).To(Equal(1.0 / math.Sqrt(14)))
		Expect(nv.Y).To(Equal(2.0 / math.Sqrt(14)))
		Expect(nv.Z).To(Equal(3.0 / math.Sqrt(14)))
		Expect(nv.W).To(Equal(v.W))
	})
	It("Given a vector (1, 2, 3), Normal() should return a vector (1/√14, 2/√14, 3/√14)", func() {
		v := tuple.NewVector(1.0, 2.0, 3.0)
		nv := v.Normal()
		Expect(nv.Mag()).To(Equal(1.0))
	})
})

var _ = Describe("Dot Product", func() {
	It("Given vector a (1, 2, 3) and b (2, 3, 4), dot product of a and b gives 20", func() {
		a := tuple.NewVector(1.0, 2.0, 3.0)
		b := tuple.NewVector(2.0, 3.0, 4.0)
		p := a.Dot(b)
		Expect(p).To(Equal(20.0))
	})

})

var _ = Describe("Cross Product", func() {
	It("Given vector a (1, 2, 3) and b (2, 3, 4), cross product of a X b gives a vector (-1, 2, -1)", func() {
		a := tuple.NewVector(1.0, 2.0, 3.0)
		b := tuple.NewVector(2.0, 3.0, 4.0)
		p := a.Cross(b)
		Expect(p.X).To(Equal(-1.0))
		Expect(p.Y).To(Equal(2.0))
		Expect(p.Z).To(Equal(-1.0))
		Expect(p.W).To(Equal(a.W))
		Expect(p.W).To(Equal(b.W))
	})
	It("Given vector a (1, 2, 3) and b (2, 3, 4), cross product of b X a gives a vector (1, -2, 1)", func() {
		a := tuple.NewVector(1.0, 2.0, 3.0)
		b := tuple.NewVector(2.0, 3.0, 4.0)
		p := b.Cross(a)
		Expect(p.X).To(Equal(1.0))
		Expect(p.Y).To(Equal(-2.0))
		Expect(p.Z).To(Equal(1.0))
		Expect(p.W).To(Equal(a.W))
		Expect(p.W).To(Equal(b.W))
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
