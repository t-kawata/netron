package crypto_test

import (
	. "github.com/mudler/netron/pkg/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mudler/netron/pkg/crypto"
)

var _ = Describe("Crypto utilities", func() {
	Context("AES", func() {
		It("Encode/decode", func() {
			key := RandStringRunes(32)
			message := "foo"
			k := [32]byte{}
			copy([]byte(key)[:], k[:32])

			encoded, err := AESEncrypt(message, &k)
			Expect(err).ToNot(HaveOccurred())
			Expect(encoded).ToNot(Equal(key))
			Expect(len(encoded)).To(Equal(62))

			// Encode again
			encoded2, err := AESEncrypt(message, &k)
			Expect(err).ToNot(HaveOccurred())

			// should differ
			Expect(encoded2).ToNot(Equal(encoded))

			// Decrypt and check
			decoded, err := AESDecrypt(encoded, &k)
			Expect(err).ToNot(HaveOccurred())
			Expect(decoded).To(Equal(message))

			decoded, err = AESDecrypt(encoded2, &k)
			Expect(err).ToNot(HaveOccurred())
			Expect(decoded).To(Equal(message))
		})
	})
})
