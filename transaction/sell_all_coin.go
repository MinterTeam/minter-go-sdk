package transaction

import (
	"math/big"

	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
)

// SellAllCoinData is Data of Transaction for selling one coin (owned by sender) in favour of another coin in a system.
type SellAllCoinData struct {
	CoinToSell        CoinID   // ID of a coin to give
	CoinToBuy         CoinID   // ID of a coin to get
	MinimumValueToBuy *big.Int // Minimum value of coins to get
}

// NewSellAllCoinData returns new SellAllCoinData of Transaction for selling one coin (owned by sender) in favour of another coin in a system.
func NewSellAllCoinData() *SellAllCoinData {
	return &SellAllCoinData{}
}

// SetCoinToSell sets ID of a coin to give.
func (d *SellAllCoinData) SetCoinToSell(id uint64) *SellAllCoinData {
	d.CoinToSell = CoinID(id)
	return d
}

// SetCoinToBuy sets ID of a coin to get.
func (d *SellAllCoinData) SetCoinToBuy(id uint64) *SellAllCoinData {
	d.CoinToBuy = CoinID(id)
	return d
}

// SetMinimumValueToBuy sets minimum value of coins to get
func (d *SellAllCoinData) SetMinimumValueToBuy(value *big.Int) *SellAllCoinData {
	d.MinimumValueToBuy = value
	return d
}

// Type returns Data type of the transaction.
func (d *SellAllCoinData) Type() Type {
	return TypeSellAllCoin
}

// Encode returns the byte representation of a transaction Data.
func (d *SellAllCoinData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
