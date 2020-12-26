package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"math/big"
)

// UnbondData is a Data of Transaction for unbonding funds from validator's stake.
type UnbondData struct {
	PubKey PublicKey // Public key of a validator
	Coin   CoinID    // ID of coin to stake
	Value  *big.Int  // Amount of coins to stake
}

// NewUnbondData create data of Transaction for unbonding funds from validator's stake
func NewUnbondData() *UnbondData {
	return &UnbondData{}
}

// SetPubKey sets Public key of a validator
func (d *UnbondData) SetPubKey(key string) (*UnbondData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// MustSetPubKey tries to set public key of validator and panics on error.
func (d *UnbondData) MustSetPubKey(key string) *UnbondData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// SetCoin sets ID of coin to stake
func (d *UnbondData) SetCoin(id uint64) *UnbondData {
	d.Coin = CoinID(id)
	return d
}

// SetValue sets amount of coins to stake
func (d *UnbondData) SetValue(value *big.Int) *UnbondData {
	d.Value = value
	return d
}

// Type returns Data type of the transaction.
func (d *UnbondData) Type() Type {
	return TypeUnbond
}

// Fee returns commission of transaction Data
func (d *UnbondData) Fee() Fee {
	return feeTypeUnbond
}

// Encode returns the byte representation of a transaction Data.
func (d *UnbondData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
