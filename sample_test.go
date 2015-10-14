package go_ginkgo_test

import (
	. "github.com/135yshr/go_ginkgo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {
	Describe("#say", func() {
		It("Hello TDD BootCampを返すこと", func() {
			Expect(Say()).To(Equal("Hello TDD BootCamp!!"))
		})
	})
})
