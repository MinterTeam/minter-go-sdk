package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction for selling one coin (owned by sender) in favour of another coin in a system.
// CoinToSell - Symbol of a coin to give. ValueToSell - Amount of CoinToSell to give.
// CoinToBuy - Symbol of a coin to get. MinimumValueToBuy - Minimum value of coins to get.
type SellCoinData struct {
	CoinToSell        Coin
	ValueToSell       *big.Int
	CoinToBuy         Coin
	MinimumValueToBuy *big.Int
}

func NewSellCoinData() *SellCoinData {
	return &SellCoinData{}
}

func (d *SellCoinData) SetCoinToSell(symbol string) *SellCoinData {
	copy(d.CoinToSell[:], symbol)
	return d
}

func (d *SellCoinData) SetCoinToBuy(symbol string) *SellCoinData {
	copy(d.CoinToBuy[:], symbol)
	return d
}

func (d *SellCoinData) SetValueToSell(value *big.Int) *SellCoinData {
	d.ValueToSell = value
	return d
}

func (d *SellCoinData) SetMinimumValueToBuy(value *big.Int) *SellCoinData {
	d.MinimumValueToBuy = value
	return d
}

func (d *SellCoinData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (d *SellCoinData) fee() fee {
	return feeTypeSellCoin
}
