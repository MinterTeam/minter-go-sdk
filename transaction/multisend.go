package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction for sending coins to multiple addresses.
type MultisendData struct {
	List []MultisendDataItem
}

type MultisendDataItem SendData

func NewMultiMultisendDataItem() *MultisendData {
	return &MultisendData{}
}

func NewMultisendDataItem() *MultisendDataItem {
	return &MultisendDataItem{}
}

func (d *MultisendDataItem) SetCoin(symbol string) *MultisendDataItem {
	copy(d.Coin[:], symbol)
	return d
}

func (d *MultisendDataItem) SetTo(address string) (*MultisendDataItem, error) {
	bytes, err := addressToHex(address)
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
	d.Value = value
	return d
}

func (d *MultisendData) AddItem(item MultisendDataItem) *MultisendData {
	d.List = append(d.List, item)
	return d
}

func (d *MultisendData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
func (d *MultisendData) fee() Fee {
	return Fee(10 + (len(d.List)-1)*5)
}
