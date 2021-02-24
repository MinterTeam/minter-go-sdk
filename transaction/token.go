package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

var MaxCoinSupply = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(15+18), nil)

type CreateTokenData struct {
	Name          string
	Symbol        CoinSymbol
	InitialAmount *big.Int
	MaxSupply     *big.Int
	Mintable      bool
	Burnable      bool
}

func NewCreateTokenData() *CreateTokenData {
	return &CreateTokenData{}
}

// SetName sets name of a coin. Arbitrary string up to 64 letters length.
func (d *CreateTokenData) SetName(name string) *CreateTokenData {
	d.Name = name
	return d
}

// SetSymbol sets symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length.
func (d *CreateTokenData) SetSymbol(symbol string) *CreateTokenData {
	copy(d.Symbol[:], symbol)
	return d
}

// SetInitialAmount sets amount of coins to issue. Issued coins will be available to sender account.
func (d *CreateTokenData) SetInitialAmount(value *big.Int) *CreateTokenData {
	d.InitialAmount = value
	return d
}

// SetMaxSupply sets maximum amount of coins that are allowed to be issued.
func (d *CreateTokenData) SetMaxSupply(maxSupply *big.Int) *CreateTokenData {
	d.MaxSupply = maxSupply
	return d
}

func (d *CreateTokenData) SetMintable(mintable bool) *CreateTokenData {
	d.Mintable = mintable
	return d
}
func (d *CreateTokenData) SetBurnable(burnable bool) *CreateTokenData {
	d.Burnable = burnable
	return d
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

func NewRecreateTokenData() *RecreateTokenData {
	return &RecreateTokenData{}
}

// SetName sets name of a coin. Arbitrary string up to 64 letters length.
func (d *RecreateTokenData) SetName(name string) *RecreateTokenData {
	d.Name = name
	return d
}

// SetSymbol sets symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length.
func (d *RecreateTokenData) SetSymbol(symbol string) *RecreateTokenData {
	copy(d.Symbol[:], symbol)
	return d
}

// SetInitialAmount sets amount of coins to issue. Issued coins will be available to sender account.
func (d *RecreateTokenData) SetInitialAmount(value *big.Int) *RecreateTokenData {
	d.InitialAmount = value
	return d
}

// SetMaxSupply sets maximum amount of coins that are allowed to be issued.
func (d *RecreateTokenData) SetMaxSupply(maxSupply *big.Int) *RecreateTokenData {
	d.MaxSupply = maxSupply
	return d
}

func (d *RecreateTokenData) SetMintable(mintable bool) *RecreateTokenData {
	d.Mintable = mintable
	return d
}
func (d *RecreateTokenData) SetBurnable(burnable bool) *RecreateTokenData {
	d.Burnable = burnable
	return d
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

func NewBurnTokenData() *BurnTokenData {
	return &BurnTokenData{}
}

func (d *BurnTokenData) SetValue(value *big.Int) *BurnTokenData {
	d.Value = value
	return d
}

func (d *BurnTokenData) SetCoin(id uint64) *BurnTokenData {
	d.Coin = CoinID(id)
	return d
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

func NewMintTokenData() *MintTokenData {
	return &MintTokenData{}
}

func (d *MintTokenData) SetCoin(id uint64) *MintTokenData {
	d.Coin = CoinID(id)
	return d
}

func (d *MintTokenData) SetValue(value *big.Int) *MintTokenData {
	d.Value = value
	return d
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
