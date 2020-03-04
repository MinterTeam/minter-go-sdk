package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"github.com/ethereum/go-ethereum/crypto"
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

func (d *CreateMultisigData) addAddress(address string) (*CreateMultisigData, error) {
	hexAddress, err := addressToHex(address)
	if err != nil {
		return d, err
	}
	var a [20]byte
	copy(a[:], hexAddress)
	d.Addresses = append(d.Addresses, a)
	return d, nil
}

func (d *CreateMultisigData) addWeight(weight uint) *CreateMultisigData {
	d.Weights = append(d.Weights, weight)
	return d
}

func (d *CreateMultisigData) AddSigData(address string, weight uint) (*CreateMultisigData, error) {
	_, err := d.addAddress(address)
	if err != nil {
		return nil, err
	}

	d.addWeight(weight)

	return d, nil
}

func (d *CreateMultisigData) MustAddSigData(address string, weight uint) *CreateMultisigData {
	_, err := d.AddSigData(address, weight)
	if err != nil {
		panic(err)
	}

	return d
}

func (d *CreateMultisigData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (d *CreateMultisigData) fee() Fee {
	return feeTypeCreateMultisig
}

func (d *CreateMultisigData) MultisigAddress() [20]byte {
	b, err := rlp.EncodeToBytes(d)
	if err != nil {
		panic(err)
	}

	var addr [20]byte
	copy(addr[:], crypto.Keccak256(b)[12:])

	return addr
}

func (d *CreateMultisigData) MultisigAddressString() string {
	return wallet.BytesToAddress(d.MultisigAddress())
}
