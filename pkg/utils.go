package pkg

import (
	. "github.com/onsi/gomega"
)

// Custom matcher to compare floating-point numbers with a tolerance
func BeCloseTo(expected interface{}, tolerance float64) OmegaMatcher {
	return WithTransform(func(actual interface{}) float64 {
		return actual.(float64)
	}, BeNumerically("~", expected, tolerance))
}
