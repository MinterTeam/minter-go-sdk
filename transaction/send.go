package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type SendData struct {
	Coin  [10]byte
	To    [20]byte
	Value *big.Int
}

func NewSendData() *SendData {
	return &SendData{}
}

func (d *SendData) SetCoin(symbol string) *SendData {
	copy(d.Coin[:], symbol)
	return d
}

func (d *SendData) SetTo(address string) (*SendData, error) {
	bytes, err := AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.To[:], bytes)
	return d, nil
}

func (d *SendData) MustSetTo(address string) *SendData {
	_, err := d.SetTo(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *SendData) SetValue(value *big.Int) *SendData {
	d.Value = value
	return d
}

func (d *SendData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
