package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
	"math/big"
)

// SellCoinData is Data of Transaction for selling one coin (owned by sender) in favour of another coin in a system.
type SellCoinData struct {
	CoinToSell        CoinID   // ID of a coin to give
	ValueToSell       *big.Int // Amount of CoinToSell to give
	CoinToBuy         CoinID   // ID of a coin to get
	MinimumValueToBuy *big.Int // Minimum value of coins to get
}

// NewSellCoinData returns new SellCoinData of Transaction for selling one coin (owned by sender) in favour of another coin in a system.
func NewSellCoinData() *SellCoinData {
	return &SellCoinData{}
}

// SetCoinToSell sets ID of a coin to give
func (d *SellCoinData) SetCoinToSell(id uint32) *SellCoinData {
	d.CoinToSell = CoinID(id)
	return d
}

// SetCoinToBuy sets ID of a coin to get
func (d *SellCoinData) SetCoinToBuy(id uint32) *SellCoinData {
	d.CoinToBuy = CoinID(id)
	return d
}

// SetValueToSell sets amount of CoinToSell to give
func (d *SellCoinData) SetValueToSell(value *big.Int) *SellCoinData {
	d.ValueToSell = value
	return d
}

// SetMinimumValueToBuy sets minimum value of coins to get
func (d *SellCoinData) SetMinimumValueToBuy(value *big.Int) *SellCoinData {
	d.MinimumValueToBuy = value
	return d
}

// Type returns Data type of the transaction.
func (d *SellCoinData) Type() Type {
	return TypeSellCoin
}

// Encode returns the byte representation of a transaction Data.
func (d *SellCoinData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
