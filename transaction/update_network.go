package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
)

type VoteUpdateData struct {
	Version string
	PubKey  PublicKey
	Height  uint64
}

// Type returns Data type of the transaction.
func (d *VoteUpdateData) Type() Type {
	return TypeVoteUpdate
}

// Fee returns commission of transaction Data
func (d *VoteUpdateData) Fee() Fee {
	return 0 // todo
}

// Encode returns the byte representation of a transaction Data.
func (d *VoteUpdateData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
