package crypto

import "github.com/haterb4/mp-quic/internal/protocol"

// NewNullAEAD creates a NullAEAD
func NewNullAEAD(p protocol.Perspective, v protocol.VersionNumber) AEAD {
	if v.UsesTLS() {
		return &nullAEADFNV64a{}
	}
	return &nullAEADFNV128a{perspective: p}
}
