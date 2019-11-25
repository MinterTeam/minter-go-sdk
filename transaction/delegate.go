package transaction

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction for delegating funds to validator.
// PubKey - Public key of a validator. Coin - Symbol of coin to stake. Value - Amount of coins to stake.
type DelegateData struct {
	PubKey []byte
	Coin   [10]byte
	Value  *big.Int
}

func NewDelegateData() *DelegateData {
	return &DelegateData{}
}

func (d *DelegateData) SetPubKey(key string) (*DelegateData, error) {
	var err error
	d.PubKey, err = hex.DecodeString(key[2:])
	if err != nil {
		return d, err
	}
	return d, nil
}

func (d *DelegateData) MustSetPubKey(key string) *DelegateData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *DelegateData) SetCoin(symbol string) *DelegateData {
	copy(d.Coin[:], symbol)
	return d
}

func (d *DelegateData) SetValue(value *big.Int) *DelegateData {
	d.Value = value
	return d
}

func (d *DelegateData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (d *DelegateData) fee() Fee {
	return feeTypeDelegate
}
