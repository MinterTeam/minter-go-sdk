package transaction

import (
	"math/big"

	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
)

// MaxCoinSupply returns max available coin supply
func MaxCoinSupply() *big.Int {
	return big.NewInt(0).Exp(big.NewInt(10), big.NewInt(15+18), nil)
}

// CreateTokenData is a Data of Transaction for creation of a token (non-reserve coin).
type CreateTokenData struct {
	Name          string
	Symbol        CoinSymbol
	InitialAmount *big.Int // Number of tokens to be created at the start
	MaxSupply     *big.Int // Upper limit of the total number of tokens
	Mintable      bool     // Allow new tokens to be issued additionally
	Burnable      bool     // Allow all tokens to be burned
}

// NewCreateTokenData creates CreateTokenData
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

// Encode returns the byte representation of a transaction Data.
func (d *CreateTokenData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

// RecreateTokenData is a Data of Transaction for re-creation the coins (both backed and non-reserve).
type RecreateTokenData struct {
	Name          string
	Symbol        CoinSymbol
	InitialAmount *big.Int
	MaxSupply     *big.Int
	Mintable      bool
	Burnable      bool
}

// NewRecreateTokenData creates RecreateTokenData
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

// Encode returns the byte representation of a transaction Data.
func (d *RecreateTokenData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

// BurnTokenData is a Data of Transaction for decreasing the token's supply. Can be applied to tokens only and is executed from the address of the user who has the necessary amount of this coin. The new supply must be more than or equal to 0.
type BurnTokenData struct {
	Coin  CoinID
	Value *big.Int
}

// NewBurnTokenData creates BurnTokenData
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

// Encode returns the byte representation of a transaction Data.
func (d *BurnTokenData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

// MintTokenData is a Data of Transaction for increasing the token's supply. Can be applied to tokens only and is executed from the coin owner address. The new supply must not exceed the MaxSupply value.
type MintTokenData struct {
	Coin  CoinID
	Value *big.Int
}

// NewMintTokenData creates MintTokenData
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

// Encode returns the byte representation of a transaction Data.
func (d *MintTokenData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
