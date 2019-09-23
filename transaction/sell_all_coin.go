package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction for selling one coin (owned by sender) in favour of another coin in a system.
// CoinToSell - Symbol of a coin to give. ValueToSell - Amount of CoinToSell to give.
// CoinToBuy - Symbol of a coin to get. MinimumValueToBuy - Minimum value of coins to get.
type SellAllCoinData struct {
	CoinToSell        [10]byte
	CoinToBuy         [10]byte
	MinimumValueToBuy *big.Int
}

func NewSellAllCoinData() *SellAllCoinData {
	return &SellAllCoinData{}
}

func (d *SellAllCoinData) SetCoinToSell(symbol string) *SellAllCoinData {
	copy(d.CoinToSell[:], symbol)
	return d
}

func (d *SellAllCoinData) SetCoinToBuy(symbol string) *SellAllCoinData {
	copy(d.CoinToBuy[:], symbol)
	return d
}

func (d *SellAllCoinData) SetMinimumValueToBuy(value *big.Int) *SellAllCoinData {
	d.MinimumValueToBuy = value
	return d
}

func (d *SellAllCoinData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
func (d *SellAllCoinData) fee() Fee {
	return feeTypeSellAllCoin
}
