package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"strings"
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
	MaxSupply            *big.Int
}

func NewCreateCoinData() *CreateCoinData {
	return &CreateCoinData{
		InitialReserve: big.NewInt(0).Mul(big.NewInt(10000), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
	}
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
	if big.NewInt(0).Mul(big.NewInt(10000), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)).Cmp(value) == -1 {
		d.InitialReserve = value
	}
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
	switch strings.Index(string(d.Symbol[:]), "\x00") {
	case 3:
		return 1000000 * feeTypeCreateCoin
	case 4:
		return 100000 * feeTypeCreateCoin
	case 5:
		return 10000 * feeTypeCreateCoin
	case 6:
		return 1000 * feeTypeCreateCoin
	case 7, 8, 9, 10:
		return 100 * feeTypeCreateCoin
	default:
		return feeTypeCreateCoin
	}
}
