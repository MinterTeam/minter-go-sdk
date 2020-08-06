package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction data for delegating funds to validator.
type DelegateData struct {
	PubKey [32]byte // Public key of a validator
	Coin   CoinID   // ID of coin to stake
	Value  *big.Int // Amount of coins to stake
}

// New data of transaction for delegating funds to validator
func NewDelegateData() *DelegateData {
	return &DelegateData{}
}

// Set public key of a validator.
func (d *DelegateData) SetPubKey(key string) (*DelegateData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

//  Tries to set public key of validator and panics on error.
func (d *DelegateData) MustSetPubKey(key string) *DelegateData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Set ID of coin to stake.
func (d *DelegateData) SetCoin(id CoinID) *DelegateData {
	d.Coin = id
	return d
}

// Set amount of coins to stake.
func (d *DelegateData) SetValue(value *big.Int) *DelegateData {
	d.Value = value
	return d
}

func (d *DelegateData) Type() Type {
	return TypeDelegate
}

func (d *DelegateData) Fee() Fee {
	return feeTypeDelegate
}

func (d *DelegateData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
