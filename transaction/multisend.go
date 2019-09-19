package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type MultiMultisendDataItem struct {
	List []MultisendDataItem
}

type MultisendDataItem SendData

func NewMultiMultisendDataItem() *MultiMultisendDataItem {
	return &MultiMultisendDataItem{}
}

func NewMultisendDataItem() *MultisendDataItem {
	return &MultisendDataItem{}
}

func (d *MultisendDataItem) SetCoin(symbol string) *MultisendDataItem {
	copy(d.Coin[:], symbol)
	return d
}

func (d *MultisendDataItem) SetTo(address string) (*MultisendDataItem, error) {
	bytes, err := AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.To[:], bytes)
	return d, nil
}

func (d *MultisendDataItem) MustSetTo(address string) *MultisendDataItem {
	_, err := d.SetTo(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *MultisendDataItem) SetValue(value *big.Int) *MultisendDataItem {
	d.Value = big.NewInt(0).Mul(value, expPip)
	return d
}

func (d *MultiMultisendDataItem) AddItem(item MultisendDataItem) *MultiMultisendDataItem {
	d.List = append(d.List, item)
	return d
}

func (d *MultiMultisendDataItem) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
