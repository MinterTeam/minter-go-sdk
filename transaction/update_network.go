package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
)

type UpdateNetworkData struct {
	Version string
	PubKey  PublicKey
	Height  uint64
}

// Type returns Data type of the transaction.
func (d *UpdateNetworkData) Type() Type {
	return TypeUpdateNetwork
}

// Fee returns commission of transaction Data
func (d *UpdateNetworkData) Fee() Fee {
	return 0 // todo
}

// Encode returns the byte representation of a transaction Data.
func (d *UpdateNetworkData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
