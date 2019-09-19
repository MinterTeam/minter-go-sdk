package transaction

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type UnbondData struct {
	PubKey []byte
	Coin   [10]byte
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
