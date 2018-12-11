package helium_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHelium(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helium Suite")
}
