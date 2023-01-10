package fiber

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testingObj *testing.T

func TestFiber(t *testing.T) {
	testingObj = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fiber Suite")
}

var _ = Describe("Fiber", Ordered, func() {

})
