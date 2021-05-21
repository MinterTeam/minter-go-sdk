package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
)

// MultisendData is a Data of Transaction for sending coins to multiple addresses.
type MultisendData struct {
	List []*SendData // List of SendData
}

// NewMultisendData returns new MultisendData of Transaction for sending coins to multiple addresses
func NewMultisendData() *MultisendData {
	return &MultisendData{}
}

// AddItem adds SendData to Multisend list
func (d *MultisendData) AddItem(item *SendData) *MultisendData {
	d.List = append(d.List, item)
	return d
}

// Type returns Data type of the transaction.
func (d *MultisendData) Type() Type {
	return TypeMultisend
}

// Encode returns the byte representation of a transaction Data.
func (d *MultisendData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
