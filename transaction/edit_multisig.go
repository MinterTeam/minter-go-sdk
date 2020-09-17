package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction data for edit multisig owners.
type EditMultisigData struct {
	Threshold uint       // Threshold for the sums of signature weights.
	Weights   []uint     // Weights of signers
	Addresses [][20]byte // List of signed addresses
}

// Data of transaction for edit multisig owners.
func NewEditMultisigData() *EditMultisigData {
	return &EditMultisigData{}
}

// Set threshold for the sums of signature weights.
func (d *EditMultisigData) SetThreshold(threshold uint) *EditMultisigData {
	d.Threshold = threshold
	return d
}

// Set a set of signers with appropriate weights.
func (d *EditMultisigData) AddSigData(address string, weight uint) (*EditMultisigData, error) {
	_, err := d.addAddress(address)
	if err != nil {
		return nil, err
	}

	d.addWeight(weight)

	return d, nil
}

// Tries to set a set of signers with appropriate weights and panics on error.
func (d *EditMultisigData) MustAddSigData(address string, weight uint) *EditMultisigData {
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
	var a [20]byte
	copy(a[:], hexAddress)
	d.Addresses = append(d.Addresses, a)
	return d, nil
}

func (d *EditMultisigData) addWeight(weight uint) *EditMultisigData {
	d.Weights = append(d.Weights, weight)
	return d
}

func (d *EditMultisigData) Type() Type {
	return TypeEditMultisig
}

func (d *EditMultisigData) Fee() Fee {
	return feeEditMultisig
}

func (d *EditMultisigData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
