package transaction

import (
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

func (d *UnbondData) SetPubKey(key string) *UnbondData {
	d.PubKey = []byte(key)
	return d
}

func (d *UnbondData) SetCoin(symbol string) *UnbondData {
	copy(d.Coin[:], symbol)
	return d
}

func (d *UnbondData) SetValue(value *big.Int) *UnbondData {
	d.Value = big.NewInt(0).Mul(value, expPip)
	return d
}

func (d *UnbondData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
