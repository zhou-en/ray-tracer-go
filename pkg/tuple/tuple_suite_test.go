package tuple_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTuple(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tuple Suite")
}
