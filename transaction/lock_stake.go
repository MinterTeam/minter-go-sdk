package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
)

type LockStakeData struct {
}

func NewLockStakeData() *LockStakeData {
	return &LockStakeData{}
}

// Type returns Data type of the transaction.
func (d *LockStakeData) Type() Type {
	return TypeLockStake
}

// Encode returns the byte representation of a transaction Data.
func (d *LockStakeData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
