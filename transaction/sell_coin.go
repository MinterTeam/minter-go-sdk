package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction for selling one coin (owned by sender) in favour of another coin in a system.
type SellCoinData struct {
	CoinToSell        CoinID   // ID of a coin to give
	ValueToSell       *big.Int // Amount of CoinToSell to give
	CoinToBuy         CoinID   // ID of a coin to get
	MinimumValueToBuy *big.Int // Minimum value of coins to get
}

// New SellCoinData of Transaction for selling one coin (owned by sender) in favour of another coin in a system.
func NewSellCoinData() *SellCoinData {
	return &SellCoinData{}
}

// Set ID of a coin to give
func (d *SellCoinData) SetCoinToSell(id CoinID) *SellCoinData {
	d.CoinToSell = id
	return d
}

// Set ID of a coin to get
func (d *SellCoinData) SetCoinToBuy(id CoinID) *SellCoinData {
	d.CoinToBuy = id
	return d
}

// Set amount of CoinToSell to give
func (d *SellCoinData) SetValueToSell(value *big.Int) *SellCoinData {
	d.ValueToSell = value
	return d
}

// Set minimum value of coins to get
func (d *SellCoinData) SetMinimumValueToBuy(value *big.Int) *SellCoinData {
	d.MinimumValueToBuy = value
	return d
}

func (d *SellCoinData) Type() Type {
	return TypeSellCoin
}

func (d *SellCoinData) Fee() Fee {
	return feeTypeSellCoin
}

func (d *SellCoinData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
