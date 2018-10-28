package generator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/xplaceholder/ashandler/generator"
)

var _ = Describe("AsGenerator", func() {

	var (
		generator ASGenerator
	)

	BeforeEach(func() {
		generator = ASGenerator{}
	})
	Describe("Generate", func() {
		Context("Everything OK", func() {
			It("should succeed", func() {
				Expect(generator.Generate(nil)).To(BeNil())
			})
		})
	})
})
