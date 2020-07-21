package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction for selling one coin (owned by sender) in favour of another coin in a system.
type SellAllCoinData struct {
	CoinToSell        CoinID   // ID of a coin to give
	CoinToBuy         CoinID   // ID of a coin to get
	MinimumValueToBuy *big.Int // Minimum value of coins to get
}

// New SellAllCoinData of transaction for selling one coin (owned by sender) in favour of another coin in a system.
func NewSellAllCoinData() *SellAllCoinData {
	return &SellAllCoinData{}
}

// Set ID of a coin to give.
func (d *SellAllCoinData) SetCoinToSell(id CoinID) *SellAllCoinData {
	d.CoinToSell = id
	return d
}

// Set ID of a coin to get.
func (d *SellAllCoinData) SetCoinToBuy(id CoinID) *SellAllCoinData {
	d.CoinToBuy = id
	return d
}

// Set minimum value of coins to get
func (d *SellAllCoinData) SetMinimumValueToBuy(value *big.Int) *SellAllCoinData {
	d.MinimumValueToBuy = value
	return d
}

func (d *SellAllCoinData) Type() Type {
	return TypeSellAllCoin
}

func (d *SellAllCoinData) Fee() Fee {
	return feeTypeSellAllCoin
}

func (d *SellAllCoinData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
