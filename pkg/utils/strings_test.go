package utils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mudler/netron/pkg/utils"
)

var _ = Describe("String utilities", func() {
	Context("RandStringRunes", func() {
		It("returns a string with the correct length", func() {
			Expect(len(RandStringRunes(10))).To(Equal(10))
		})
	})
})
