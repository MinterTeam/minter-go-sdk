package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction data for creating new coin in a system.
type CreateCoinData struct {
	Name                 string     // Name of a coin
	Symbol               CoinSymbol // Symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length
	InitialAmount        *big.Int   // Amount of coins to issue. Issued coins will be available to sender account. Should be between 1 and 1,000,000,000,000,000 coins.
	InitialReserve       *big.Int   // Initial reserve in BIP's
	ConstantReserveRatio uint       // ConstantReserveRatio (CRR), should be from 10 to 100.
	MaxSupply            *big.Int   // Max amount of coins that are allowed to be issued. Maximum is 1,000,000,000,000,000
}

// New data of transaction for creating new coin in a system.
func NewCreateCoinData() *CreateCoinData {
	return &CreateCoinData{}
}

// Set name of a coin. Arbitrary string up to 64 letters length.
func (d *CreateCoinData) SetName(name string) *CreateCoinData {
	d.Name = name
	return d
}

// Set symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length.
func (d *CreateCoinData) SetSymbol(symbol string) *CreateCoinData {
	copy(d.Symbol[:], symbol)
	return d
}

// Set initial reserve in BIP's.
func (d *CreateCoinData) SetInitialReserve(value *big.Int) *CreateCoinData {
	d.InitialReserve = value
	return d
}

// Set amount of coins to issue. Issued coins will be available to sender account.
func (d *CreateCoinData) SetInitialAmount(value *big.Int) *CreateCoinData {
	d.InitialAmount = value
	return d
}

// Set ConstantReserveRatio (CRR), uint, should be from 10 to 100.
func (d *CreateCoinData) SetConstantReserveRatio(ratio uint) *CreateCoinData {
	d.ConstantReserveRatio = ratio
	return d
}

// Set maximum amount of coins that are allowed to be issued.
func (d *CreateCoinData) SetMaxSupply(maxSupply *big.Int) *CreateCoinData {
	d.MaxSupply = maxSupply
	return d
}

func (d *CreateCoinData) Type() Type {
	return TypeCreateCoin
}

func (d *CreateCoinData) Fee() Fee {
	switch len(d.Symbol.String()) {
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

func (d *CreateCoinData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
