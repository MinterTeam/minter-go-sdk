package transaction

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type DelegateData struct {
	PubKey []byte
	Coin   [10]byte
	Stake  *big.Int
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
	return d, err
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

func (d *DelegateData) SetStake(value *big.Int) *DelegateData {
	d.Stake = big.NewInt(0).Mul(value, expPip)
	return d
}

func (d *DelegateData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
