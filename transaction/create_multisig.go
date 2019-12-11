package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
)

type CreateMultisigData struct {
	Threshold uint
	Weights   []uint
	Addresses [][20]byte
}

func NewCreateMultisigData() *CreateMultisigData {
	return &CreateMultisigData{}
}

func (d *CreateMultisigData) SetThreshold(threshold uint) *CreateMultisigData {
	d.Threshold = threshold
	return d
}

func (d *CreateMultisigData) SetAddresses(addresses [][20]byte) *CreateMultisigData {
	d.Addresses = addresses
	return d
}

func (d *CreateMultisigData) AddAddress(address string) *CreateMultisigData {
	var a [20]byte
	copy(a[:], address)
	d.Addresses = append(d.Addresses, a)
	return d
}

func (d *CreateMultisigData) SetWeights(weights []uint) *CreateMultisigData {
	d.Weights = weights
	return d
}

func (d *CreateMultisigData) AddWeight(weight uint) *CreateMultisigData {
	d.Weights = append(d.Weights, weight)
	return d
}

func (d *CreateMultisigData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (d *CreateMultisigData) fee() Fee {
	return feeTypeCreateMultisig
}
