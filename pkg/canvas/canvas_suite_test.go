package canvas_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCanvas(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Canvas Suite")
}
