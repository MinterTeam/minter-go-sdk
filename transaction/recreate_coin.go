package transaction

import (
	"math/big"

	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
)

// RecreateCoinData is a Data of Transaction for recreating new coin.
type RecreateCoinData struct {
	Name                 string     // Name of a coin
	Symbol               CoinSymbol // Symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length
	InitialAmount        *big.Int   // Amount of coins to issue. Issued coins will be available to sender account. Should be between 1 and 1,000,000,000,000,000 coins.
	InitialReserve       *big.Int   // Initial reserve in BIP's
	ConstantReserveRatio uint32     // ConstantReserveRatio (CRR), should be from 10 to 100.
	MaxSupply            *big.Int   // Max amount of coins that are allowed to be issued. Maximum is 1,000,000,000,000,000
}

// NewRecreateCoinData returns new RecreateCoinData of Transaction for recreating coin
func NewRecreateCoinData() *RecreateCoinData {
	return &RecreateCoinData{}
}

// SetName sets name of a coin. Arbitrary string up to 64 letters length.
func (d *RecreateCoinData) SetName(name string) *RecreateCoinData {
	d.Name = name
	return d
}

// SetSymbol sets symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length.
func (d *RecreateCoinData) SetSymbol(symbol string) *RecreateCoinData {
	copy(d.Symbol[:], symbol)
	return d
}

// SetInitialReserve sets initial reserve in BIP's.
func (d *RecreateCoinData) SetInitialReserve(value *big.Int) *RecreateCoinData {
	if BipToPip(big.NewInt(10000)).Cmp(value) == -1 {
		d.InitialReserve = value
	}
	return d
}

// SetInitialAmount sets amount of coins to issue. Issued coins will be available to sender account.
func (d *RecreateCoinData) SetInitialAmount(value *big.Int) *RecreateCoinData {
	d.InitialAmount = value
	return d
}

// SetConstantReserveRatio sets CRR, uint, should be from 10 to 100.
func (d *RecreateCoinData) SetConstantReserveRatio(ratio uint32) *RecreateCoinData {
	d.ConstantReserveRatio = ratio
	return d
}

// SetMaxSupply sets maximum amount of coins that are allowed to be issued.
func (d *RecreateCoinData) SetMaxSupply(maxSupply *big.Int) *RecreateCoinData {
	d.MaxSupply = maxSupply
	return d
}

// Type returns Data type of the transaction.
func (d *RecreateCoinData) Type() Type {
	return TypeRecreateCoin
}

// Encode returns the byte representation of a transaction Data.
func (d *RecreateCoinData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
