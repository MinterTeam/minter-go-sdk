package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type CreateTokenData struct {
	Name          string
	Symbol        CoinSymbol
	InitialAmount *big.Int
	MaxSupply     *big.Int
	Mintable      bool
	Burnable      bool
}

// Type returns Data type of the transaction.
func (d *CreateTokenData) Type() Type {
	return TypeCreateToken
}

// Fee returns commission of transaction Data
func (d *CreateTokenData) Fee() Fee {
	return feeTypeCreateCoin
}

// Encode returns the byte representation of a transaction Data.
func (d *CreateTokenData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

type RecreateTokenData struct {
	Name          string
	Symbol        CoinSymbol
	InitialAmount *big.Int
	MaxSupply     *big.Int
	Mintable      bool
	Burnable      bool
}

// Type returns Data type of the transaction.
func (d *RecreateTokenData) Type() Type {
	return TypeRecreateToken
}

// Fee returns commission of transaction Data
func (d *RecreateTokenData) Fee() Fee {
	return feeTypeRecreateCoin
}

// Encode returns the byte representation of a transaction Data.
func (d *RecreateTokenData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

type BurnTokenData struct {
	Coin  CoinID
	Value *big.Int
}

// Type returns Data type of the transaction.
func (d *BurnTokenData) Type() Type {
	return TypeBurnToken
}

// Fee returns commission of transaction Data
func (d *BurnTokenData) Fee() Fee {
	return feeTypeEditEmissionData
}

// Encode returns the byte representation of a transaction Data.
func (d *BurnTokenData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

type MintTokenData struct {
	Coin  CoinID
	Value *big.Int
}

// Type returns Data type of the transaction.
func (d *MintTokenData) Type() Type {
	return TypeMintToken
}

// Fee returns commission of transaction Data
func (d *MintTokenData) Fee() Fee {
	return feeTypeEditEmissionData
}

// Encode returns the byte representation of a transaction Data.
func (d *MintTokenData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
