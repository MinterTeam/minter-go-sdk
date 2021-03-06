package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// EditMultisigData is a Data of Transaction for editing multisig address.
type EditMultisigData struct {
	Threshold uint32    // Threshold for the sums of signature weights.
	Weights   []uint32  // Weights of signers
	Addresses []Address // List of signed addresses
}

// NewEditMultisigData returns new EditMultisigData of Transaction for editing multisig address.
func NewEditMultisigData() *EditMultisigData {
	return &EditMultisigData{}
}

// SetThreshold sets threshold for the sums of signature weights.
func (d *EditMultisigData) SetThreshold(threshold uint32) *EditMultisigData {
	d.Threshold = threshold
	return d
}

// AddSigData sets a set of signers with appropriate weights.
func (d *EditMultisigData) AddSigData(address string, weight uint32) (*EditMultisigData, error) {
	_, err := d.addAddress(address)
	if err != nil {
		return nil, err
	}

	d.addWeight(weight)

	return d, nil
}

// MustAddSigData tries to set a set of signers with appropriate weights and panics on error.
func (d *EditMultisigData) MustAddSigData(address string, weight uint32) *EditMultisigData {
	_, err := d.AddSigData(address, weight)
	if err != nil {
		panic(err)
	}

	return d
}

func (d *EditMultisigData) addAddress(address string) (*EditMultisigData, error) {
	hexAddress, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	var a Address
	copy(a[:], hexAddress)
	d.Addresses = append(d.Addresses, a)
	return d, nil
}

func (d *EditMultisigData) addWeight(weight uint32) *EditMultisigData {
	d.Weights = append(d.Weights, weight)
	return d
}

// Type returns Data type of the transaction.
func (d *EditMultisigData) Type() Type {
	return TypeEditMultisig
}

// Encode returns the byte representation of a transaction Data.
func (d *EditMultisigData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
