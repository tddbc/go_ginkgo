package go_ginkgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoGinkgo Suite")
}
