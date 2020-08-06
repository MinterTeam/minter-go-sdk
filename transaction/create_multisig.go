package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction data for creating multisig wallet.
type CreateMultisigData struct {
	Threshold uint       // Threshold for the sums of signature weights.
	Weights   []uint     // Weights of signers
	Addresses [][20]byte // List of signed addresses
}

// Data of transaction for creating multisig wallet.
func NewCreateMultisigData() *CreateMultisigData {
	return &CreateMultisigData{}
}

// Set threshold for the sums of signature weights.
func (d *CreateMultisigData) SetThreshold(threshold uint) *CreateMultisigData {
	d.Threshold = threshold
	return d
}

// Set a set of signers with appropriate weights.
func (d *CreateMultisigData) AddSigData(address string, weight uint) (*CreateMultisigData, error) {
	_, err := d.addAddress(address)
	if err != nil {
		return nil, err
	}

	d.addWeight(weight)

	return d, nil
}

// Tries to set a set of signers with appropriate weights and panics on error.
func (d *CreateMultisigData) MustAddSigData(address string, weight uint) *CreateMultisigData {
	_, err := d.AddSigData(address, weight)
	if err != nil {
		panic(err)
	}

	return d
}

func (d *CreateMultisigData) Type() Type {
	return TypeCreateMultisig
}

func (d *CreateMultisigData) Fee() Fee {
	return feeTypeCreateMultisig
}

func (d *CreateMultisigData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (d *CreateMultisigData) addAddress(address string) (*CreateMultisigData, error) {
	hexAddress, err := wallet.AddressToHex(address)
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
