package tuple

import (
	"reflect"
	"testing"
)

func TestTupleOperations(t *testing.T) {
	tests := []struct {
		name     string
		a        Tuple
		b        Tuple
		scalar   float64
		expected interface{}
		method   string
	}{
		{
			name:     "Add tuples",
			a:        NewTuple(1, 2, 3, 1),
			b:        NewTuple(4, 5, 6, 0),
			expected: NewTuple(5, 7, 9, 1),
			method:   "Add",
		},
		{
			name:     "Subtract tuples",
			a:        NewTuple(1, 2, 3, 1),
			b:        NewTuple(4, 5, 6, 0),
			expected: NewTuple(-3, -3, -3, 1),
			method:   "Sub",
		},
		{
			name:     "Negate tuple",
			a:        NewTuple(1, -2, 3, 0),
			expected: NewTuple(-1, 2, -3, 0),
			method:   "Negate",
		},
		{
			name:     "Scale tuple",
			a:        NewTuple(1, -2, 3, 0),
			scalar:   2,
			expected: NewTuple(2, -4, 6, 0),
			method:   "Scale",
		},
		{
			name:     "Divide tuple",
			a:        NewTuple(1, -2, 3, 0),
			scalar:   2,
			expected: NewTuple(0.5, -1, 1.5, 0),
			method:   "Divide",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result interface{}
			switch tt.method {
			case "Add":
				result = tt.a.Add(tt.b)
			case "Sub":
				result = tt.a.Sub(tt.b)
			case "Negate":
				result = tt.a.Negate()
			case "Scale":
				result = tt.a.Scale(tt.scalar)
			case "Divide":
				result = tt.a.Divide(tt.scalar)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestVectorOperations(t *testing.T) {
	tests := []struct {
		name     string
		v1       Vector
		v2       Vector
		scalar   float64
		expected interface{}
		method   string
	}{
		{
			name:     "Dot product",
			v1:       NewVector(1, 2, 3),
			v2:       NewVector(4, 5, 6),
			expected: 32.0,
			method:   "Dot",
		},
		{
			name:     "Magnitude",
			v1:       NewVector(3, 4, 0),
			expected: 5.0,
			method:   "Magnitude",
		},
		{
			name:     "Normalize",
			v1:       NewVector(4, 0, 0),
			expected: NewVector(1, 0, 0),
			method:   "Normalize",
		},
		{
			name:     "Cross product",
			v1:       NewVector(1, 0, 0),
			v2:       NewVector(0, 1, 0),
			expected: NewVector(0, 0, 1),
			method:   "Cross",
		},
		{
			name:     "Add vectors",
			v1:       NewVector(1, 2, 3),
			v2:       NewVector(4, 5, 6),
			expected: NewVector(5, 7, 9),
			method:   "Add",
		},
		{
			name:     "Subtract vectors",
			v1:       NewVector(1, 2, 3),
			v2:       NewVector(4, 5, 6),
			expected: NewVector(-3, -3, -3),
			method:   "Sub",
		},
		{
			name:     "Scale vector",
			v1:       NewVector(1, 2, 3),
			scalar:   2,
			expected: NewVector(2, 4, 6),
			method:   "Scale",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result interface{}
			switch tt.method {
			case "Dot":
				result = tt.v1.Dot(tt.v2)
			case "Magnitude":
				result = tt.v1.Magnitude()
			case "Normalize":
				result = tt.v1.Normalize()
			case "Cross":
				result = tt.v1.Cross(tt.v2)
			case "Add":
				result = tt.v1.Add(tt.v2)
			case "Sub":
				result = tt.v1.Sub(tt.v2)
			case "Scale":
				result = tt.v1.Scale(tt.scalar)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPointOperations(t *testing.T) {
	tests := []struct {
		name     string
		p1       Point
		v1       Vector
		expected Point
	}{
		{
			name:     "Add vector to point",
			p1:       NewPoint(1, 2, 3),
			v1:       NewVector(4, 5, 6),
			expected: NewPoint(5, 7, 9),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p1.AddVector(tt.v1)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
