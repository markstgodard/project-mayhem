package durden_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDurden(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Durden Suite")
}
