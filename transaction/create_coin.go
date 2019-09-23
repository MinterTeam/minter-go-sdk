package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction for creating new coin in a system.
// Name - Name of a coin. Arbitrary string up to 64 letters length.
// Symbol - Symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length.
// InitialAmount - Amount of coins to issue. Issued coins will be available to sender account.
// InitialReserve - Initial reserve in BIP's. ConstantReserveRatio - CRR, uint, should be from 10 to 100.
type CreateCoinData struct {
	Name                 string
	Symbol               [10]byte
	InitialAmount        *big.Int
	InitialReserve       *big.Int
	ConstantReserveRatio uint
}

func NewCreateCoinData() *CreateCoinData {
	return &CreateCoinData{}
}

func (d *CreateCoinData) SetName(name string) *CreateCoinData {
	d.Name = name
	return d
}

func (d *CreateCoinData) SetSymbol(symbol string) *CreateCoinData {
	copy(d.Symbol[:], symbol)
	return d
}

func (d *CreateCoinData) SetInitialReserve(value *big.Int) *CreateCoinData {
	d.InitialReserve = value
	return d
}

func (d *CreateCoinData) SetInitialAmount(value *big.Int) *CreateCoinData {
	d.InitialAmount = value
	return d
}

func (d *CreateCoinData) SetConstantReserveRatio(ratio uint) *CreateCoinData {
	d.ConstantReserveRatio = ratio
	return d
}

func (d *CreateCoinData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (d *CreateCoinData) fee() Fee {
	return feeTypeCreateCoin
}
