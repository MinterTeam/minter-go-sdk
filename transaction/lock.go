package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type LockData struct {
	DueBlock uint64
	Coin     CoinID
	Value    *big.Int
}

func NewLockData() *LockData {
	return &LockData{}
}

// SetDueBlock sets ID of coin to stake
func (d *LockData) SetDueBlock(height uint64) *LockData {
	d.DueBlock = height
	return d
}

// SetCoin sets ID of coin to stake
func (d *LockData) SetCoin(id uint64) *LockData {
	d.Coin = CoinID(id)
	return d
}

// SetValue sets amount of coins to stake
func (d *LockData) SetValue(value *big.Int) *LockData {
	d.Value = value
	return d
}

// Type returns Data type of the transaction.
func (d *LockData) Type() Type {
	return TypeLock
}

// Encode returns the byte representation of a transaction Data.
func (d *LockData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
