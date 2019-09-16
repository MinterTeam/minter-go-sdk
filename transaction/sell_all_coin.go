package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

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
	d.MinimumValueToBuy = big.NewInt(0).Mul(value, expPip)
	return d
}

func (d *SellAllCoinData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
