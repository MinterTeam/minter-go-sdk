package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
)

// LockStakeData is a Data of Transaction for lock stake.
type LockStakeData struct {
}

// NewLockStakeData returns new LockStakeData of Transaction for lock stake.
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
