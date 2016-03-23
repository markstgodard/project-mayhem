package spacemonkeys_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSpacemonkeys(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SpaceMonkeys Suite")
}
