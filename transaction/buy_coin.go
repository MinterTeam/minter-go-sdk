package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type BuyCoinData struct {
	CoinToBuy          [10]byte
	ValueToBuy         *big.Int
	CoinToSell         [10]byte
	MaximumValueToSell *big.Int
}

func NewBuyCoinData() *BuyCoinData {
	return &BuyCoinData{}
}

func (d *BuyCoinData) SetCoinToSell(symbol string) *BuyCoinData {
	copy(d.CoinToSell[:], symbol)
	return d
}

func (d *BuyCoinData) SetCoinToBuy(symbol string) *BuyCoinData {
	copy(d.CoinToBuy[:], symbol)
	return d
}

func (d *BuyCoinData) SetValueToBuy(value *big.Int) *BuyCoinData {
	d.ValueToBuy = big.NewInt(0).Mul(value, expPip)
	return d
}

func (d *BuyCoinData) SetMaximumValueToSell(value *big.Int) *BuyCoinData {
	d.MaximumValueToSell = big.NewInt(0).Mul(value, expPip)
	return d
}

func (d *BuyCoinData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
