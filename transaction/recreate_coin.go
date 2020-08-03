package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type RecreateCoinData struct {
	Symbol               CoinSymbol // Symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length
	InitialAmount        *big.Int   // Amount of coins to issue. Issued coins will be available to sender account. Should be between 1 and 1,000,000,000,000,000 coins.
	InitialReserve       *big.Int   // Initial reserve in BIP's
	ConstantReserveRatio uint       // ConstantReserveRatio (CRR), should be from 10 to 100.
	MaxSupply            *big.Int   // Max amount of coins that are allowed to be issued. Maximum is 1,000,000,000,000,000
}

func NewRecreateCoinData() *RecreateCoinData {
	return &RecreateCoinData{}
}

// Set symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length.
func (d *RecreateCoinData) SetSymbol(symbol string) *RecreateCoinData {
	copy(d.Symbol[:], symbol)
	return d
}

// Set initial reserve in BIP's.
func (d *RecreateCoinData) SetInitialReserve(value *big.Int) *RecreateCoinData {
	if big.NewInt(0).Mul(big.NewInt(10000), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)).Cmp(value) == -1 {
		d.InitialReserve = value
	}
	return d
}

// Set amount of coins to issue. Issued coins will be available to sender account.
func (d *RecreateCoinData) SetInitialAmount(value *big.Int) *RecreateCoinData {
	d.InitialAmount = value
	return d
}

// Set ConstantReserveRatio (CRR), uint, should be from 10 to 100.
func (d *RecreateCoinData) SetConstantReserveRatio(ratio uint) *RecreateCoinData {
	d.ConstantReserveRatio = ratio
	return d
}

// Set maximum amount of coins that are allowed to be issued.
func (d *RecreateCoinData) SetMaxSupply(maxSupply *big.Int) *RecreateCoinData {
	d.MaxSupply = maxSupply
	return d
}

func (d *RecreateCoinData) Type() Type {
	return TypeRecreateCoin
}

func (d *RecreateCoinData) Fee() Fee {
	return feeTypeRecreateCoin
}

func (d *RecreateCoinData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
