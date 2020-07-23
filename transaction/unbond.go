package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction data for unbonding funds from validator's stake.
type UnbondData struct {
	PubKey [32]byte // Public key of a validator
	Coin   CoinID   // ID of coin to stake
	Value  *big.Int // Amount of coins to stake
}

// New data of Transaction for unbonding funds from validator's stake
func NewUnbondData() *UnbondData {
	return &UnbondData{}
}

// Set Public key of a validator
func (d *UnbondData) SetPubKey(key string) (*UnbondData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// Tries to set public key of validator and panics on error.
func (d *UnbondData) MustSetPubKey(key string) *UnbondData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Set ID of coin to stake
func (d *UnbondData) SetCoin(id CoinID) *UnbondData {
	d.Coin = id
	return d
}

// Set amount of coins to stake
func (d *UnbondData) SetValue(value *big.Int) *UnbondData {
	d.Value = value
	return d
}

func (d *UnbondData) Type() Type {
	return TypeUnbond
}

func (d *UnbondData) Fee() Fee {
	return feeTypeUnbond
}

func (d *UnbondData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
