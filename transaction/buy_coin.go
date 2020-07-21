package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction data for buy a coin paying another coin (owned by sender).
type BuyCoinData struct {
	CoinToBuy          CoinID   // ID of a coin to get
	ValueToBuy         *big.Int // Amount of CoinToBuy to get
	CoinToSell         CoinID   // ID of a coin to give
	MaximumValueToSell *big.Int // Maximum value of coins to sell
}

// New data of transaction for buy a coin paying another coin (owned by sender).
func NewBuyCoinData() *BuyCoinData {
	return &BuyCoinData{}
}

// Set ID of a coin to get.
func (d *BuyCoinData) SetCoinToSell(id CoinID) *BuyCoinData {
	d.CoinToSell = id
	return d
}

// Set ID of a coin to give.
func (d *BuyCoinData) SetCoinToBuy(id CoinID) *BuyCoinData {
	d.CoinToBuy = id
	return d
}

// Set amount of CoinToBuy to get.
func (d *BuyCoinData) SetValueToBuy(value *big.Int) *BuyCoinData {
	d.ValueToBuy = value
	return d
}

// Set maximum value of coins to sell.
func (d *BuyCoinData) SetMaximumValueToSell(value *big.Int) *BuyCoinData {
	d.MaximumValueToSell = value
	return d
}

func (d *BuyCoinData) Type() Type {
	return TypeBuyCoin
}

func (d *BuyCoinData) Fee() Fee {
	return feeTypeBuyCoin
}

func (d *BuyCoinData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
