package handshake

import (
	"time"

	"github.com/haterb4/mp-quic/internal/crypto"
	"github.com/haterb4/mp-quic/internal/protocol"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ephermal KEX", func() {
	It("has a consistent KEX", func() {
		kex1 := getEphermalKEX()
		Expect(kex1).ToNot(BeNil())
		kex2 := getEphermalKEX()
		Expect(kex2).ToNot(BeNil())
		Expect(kex1).To(Equal(kex2))
	})

	It("changes KEX", func() {
		kexLifetime = time.Millisecond
		defer func() {
			kexLifetime = protocol.EphermalKeyLifetime
		}()
		kex := getEphermalKEX()
		Expect(kex).ToNot(BeNil())
		Eventually(func() crypto.KeyExchange { return getEphermalKEX() }).ShouldNot(Equal(kex))
	})
})
