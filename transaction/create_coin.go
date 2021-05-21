package transaction

import (
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"
)

// CreateCoinData is a Data of Transaction for creating new coin.
type CreateCoinData struct {
	Name                 string     // Name of a coin
	Symbol               CoinSymbol // Symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length
	InitialAmount        *big.Int   // Amount of coins to issue. Issued coins will be available to sender account. Should be between 1 and 1,000,000,000,000,000 coins.
	InitialReserve       *big.Int   // Initial reserve in BIP's
	ConstantReserveRatio uint32     // ConstantReserveRatio (CRR), should be from 10 to 100.
	MaxSupply            *big.Int   // Max amount of coins that are allowed to be issued. Maximum is 1,000,000,000,000,000
}

// NewCreateCoinData returns new CreateCoinData of Transaction for creating new coin.
func NewCreateCoinData() *CreateCoinData {
	return &CreateCoinData{}
}

// SetName sets name of a coin. Arbitrary string up to 64 letters length.
func (d *CreateCoinData) SetName(name string) *CreateCoinData {
	d.Name = name
	return d
}

// SetSymbol sets symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length.
func (d *CreateCoinData) SetSymbol(symbol string) *CreateCoinData {
	copy(d.Symbol[:], symbol)
	return d
}

// SetInitialReserve sets initial reserve in BIP's.
func (d *CreateCoinData) SetInitialReserve(value *big.Int) *CreateCoinData {
	d.InitialReserve = value
	return d
}

// SetInitialAmount sets amount of coins to issue. Issued coins will be available to sender account.
func (d *CreateCoinData) SetInitialAmount(value *big.Int) *CreateCoinData {
	d.InitialAmount = value
	return d
}

// SetConstantReserveRatio sets CRR, uint, should be from 10 to 100.
func (d *CreateCoinData) SetConstantReserveRatio(ratio uint32) *CreateCoinData {
	d.ConstantReserveRatio = ratio
	return d
}

// SetMaxSupply sets maximum amount of coins that are allowed to be issued.
func (d *CreateCoinData) SetMaxSupply(maxSupply *big.Int) *CreateCoinData {
	d.MaxSupply = maxSupply
	return d
}

// Type returns Data type of the transaction.
func (d *CreateCoinData) Type() Type {
	return TypeCreateCoin
}

// Encode returns the byte representation of a transaction Data.
func (d *CreateCoinData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
