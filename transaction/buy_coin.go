package transaction

import (
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"
)

// BuyCoinData is a Data of Transaction for buy a coin paying another coin (owned by sender).
type BuyCoinData struct {
	CoinToBuy          CoinID   // ID of a coin to get
	ValueToBuy         *big.Int // Amount of CoinToBuy to get
	CoinToSell         CoinID   // ID of a coin to give
	MaximumValueToSell *big.Int // Maximum value of coins to sell
}

// NewBuyCoinData returns new data of transaction for buy a coin paying another coin (owned by sender).
func NewBuyCoinData() *BuyCoinData {
	return &BuyCoinData{}
}

// SetCoinToSell sets ID of a coin to get.
func (d *BuyCoinData) SetCoinToSell(id uint64) *BuyCoinData {
	d.CoinToSell = CoinID(id)
	return d
}

// SetCoinToBuy sets ID of a coin to give.
func (d *BuyCoinData) SetCoinToBuy(id uint64) *BuyCoinData {
	d.CoinToBuy = CoinID(id)
	return d
}

// SetValueToBuy sets amount of CoinToBuy to get.
func (d *BuyCoinData) SetValueToBuy(value *big.Int) *BuyCoinData {
	d.ValueToBuy = value
	return d
}

// SetMaximumValueToSell sets maximum value of coins to sell.
func (d *BuyCoinData) SetMaximumValueToSell(value *big.Int) *BuyCoinData {
	d.MaximumValueToSell = value
	return d
}

// Type returns Data type of the transaction.
func (d *BuyCoinData) Type() Type {
	return TypeBuyCoin
}

// Encode returns the byte representation of a transaction Data.
func (d *BuyCoinData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
