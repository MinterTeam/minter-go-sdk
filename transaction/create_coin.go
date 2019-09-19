package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

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
