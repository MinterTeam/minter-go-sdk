package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// DelegateData is a Data of Transaction for delegating funds to validator.
type DelegateData struct {
	PubKey [32]byte // Public key of a validator
	Coin   CoinID   // ID of coin to stake
	Value  *big.Int // Amount of coins to stake
}

// NewDelegateData returns new DelegateData of Transaction for delegating funds to validator
func NewDelegateData() *DelegateData {
	return &DelegateData{}
}

// SetPubKey sets public key of a validator.
func (d *DelegateData) SetPubKey(key string) (*DelegateData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// MustSetPubKey tries to set public key of validator and panics on error.
func (d *DelegateData) MustSetPubKey(key string) *DelegateData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// SetCoin sets ID of coin to stake.
func (d *DelegateData) SetCoin(id CoinID) *DelegateData {
	d.Coin = id
	return d
}

// SetValue sets amount of coins to stake.
func (d *DelegateData) SetValue(value *big.Int) *DelegateData {
	d.Value = value
	return d
}

// Type returns Data type of the transaction.
func (d *DelegateData) Type() Type {
	return TypeDelegate
}

// Fee returns commission of transaction Data
func (d *DelegateData) Fee() Fee {
	return feeTypeDelegate
}

// Encode returns the byte representation of a transaction Data.
func (d *DelegateData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
