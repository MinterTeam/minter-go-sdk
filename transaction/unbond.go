package transaction

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction for unbonding funds from validator's stake.
// PubKey - Public key of a validator. Coin - Symbol of coin to stake. Value - Amount of coins to stake.
type UnbondData struct {
	PubKey []byte
	Coin   Coin
	Value  *big.Int
}

func NewUnbondData() *UnbondData {
	return &UnbondData{}
}

func (d *UnbondData) SetPubKey(key string) (*UnbondData, error) {
	var err error
	d.PubKey, err = hex.DecodeString(key[2:])
	if err != nil {
		return d, err
	}
	return d, nil
}

func (d *UnbondData) MustSetPubKey(key string) *UnbondData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *UnbondData) SetCoin(symbol string) *UnbondData {
	copy(d.Coin[:], symbol)
	return d
}

func (d *UnbondData) SetValue(value *big.Int) *UnbondData {
	d.Value = value
	return d
}

func (d *UnbondData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (d *UnbondData) fee() fee {
	return feeTypeUnbond
}
