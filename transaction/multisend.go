package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction data for sending coins to multiple addresses.
type MultisendData struct {
	List []*SendData
}

// Data of transaction for sending coins to multiple addresses
func NewMultisendData() *MultisendData {
	return &MultisendData{}
}

// New item for Multisend list
func NewMultisendDataItem() *SendData {
	return NewSendData()
}

// Add SendData to Multisend list
func (d *MultisendData) AddItem(item *SendData) *MultisendData {
	d.List = append(d.List, item)
	return d
}

func (d *MultisendData) Type() Type {
	return TypeMultisend
}

func (d *MultisendData) Fee() Fee {
	return Fee(10 + (len(d.List)-1)*5)
}

func (d *MultisendData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
