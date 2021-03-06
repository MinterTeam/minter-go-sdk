package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// CreateMultisigData is a Data of Transaction for creating multisig wallet.
type CreateMultisigData struct {
	Threshold uint32    // Threshold for the sums of signature weights.
	Weights   []uint32  // Weights of signers
	Addresses []Address // List of signed addresses
}

// NewCreateMultisigData returns new CreateMultisigData of Transaction for creating multisig wallet.
func NewCreateMultisigData() *CreateMultisigData {
	return &CreateMultisigData{}
}

// SetThreshold sets threshold for the sums of signature weights.
func (d *CreateMultisigData) SetThreshold(threshold uint32) *CreateMultisigData {
	d.Threshold = threshold
	return d
}

// AddSigData sets a set of signers with appropriate weights.
func (d *CreateMultisigData) AddSigData(address string, weight uint32) (*CreateMultisigData, error) {
	_, err := d.addAddress(address)
	if err != nil {
		return nil, err
	}

	d.addWeight(weight)

	return d, nil
}

// MustAddSigData tries to set a set of signers with appropriate weights and panics on error.
func (d *CreateMultisigData) MustAddSigData(address string, weight uint32) *CreateMultisigData {
	_, err := d.AddSigData(address, weight)
	if err != nil {
		panic(err)
	}

	return d
}

// Type returns Data type of the transaction.
func (d *CreateMultisigData) Type() Type {
	return TypeCreateMultisig
}

func (d *CreateMultisigData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (d *CreateMultisigData) addAddress(address string) (*CreateMultisigData, error) {
	hexAddress, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	var a Address
	copy(a[:], hexAddress)
	d.Addresses = append(d.Addresses, a)
	return d, nil
}

// Encode returns the byte representation of a transaction Data.
func (d *CreateMultisigData) addWeight(weight uint32) *CreateMultisigData {
	d.Weights = append(d.Weights, weight)
	return d
}
