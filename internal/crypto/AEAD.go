package crypto

import "github.com/haterb4/mp-quic/internal/protocol"

// An AEAD implements QUIC's authenticated encryption and associated data
type AEAD interface {
	Open(dst, src []byte, packetNumber protocol.PacketNumber, associatedData []byte) ([]byte, error)
	Seal(dst, src []byte, packetNumber protocol.PacketNumber, associatedData []byte) []byte
	Overhead() int
	GetOtherIV() []byte
}
