package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction data for edit multisig owners.
type EditMultisigOwnersData struct {
	Threshold uint       // Threshold for the sums of signature weights.
	Weights   []uint     // Weights of signers
	Addresses [][20]byte // List of signed addresses
}

// Data of transaction for edit multisig owners.
func NewEditMultisigOwnersData() *EditMultisigOwnersData {
	return &EditMultisigOwnersData{}
}

// Set threshold for the sums of signature weights.
func (d *EditMultisigOwnersData) SetThreshold(threshold uint) *EditMultisigOwnersData {
	d.Threshold = threshold
	return d
}

// Set a set of signers with appropriate weights.
func (d *EditMultisigOwnersData) AddSigData(address string, weight uint) (*EditMultisigOwnersData, error) {
	_, err := d.addAddress(address)
	if err != nil {
		return nil, err
	}

	d.addWeight(weight)

	return d, nil
}

// Tries to set a set of signers with appropriate weights and panics on error.
func (d *EditMultisigOwnersData) MustAddSigData(address string, weight uint) *EditMultisigOwnersData {
	_, err := d.AddSigData(address, weight)
	if err != nil {
		panic(err)
	}

	return d
}

func (d *EditMultisigOwnersData) addAddress(address string) (*EditMultisigOwnersData, error) {
	hexAddress, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	var a [20]byte
	copy(a[:], hexAddress)
	d.Addresses = append(d.Addresses, a)
	return d, nil
}

func (d *EditMultisigOwnersData) addWeight(weight uint) *EditMultisigOwnersData {
	d.Weights = append(d.Weights, weight)
	return d
}

func (d *EditMultisigOwnersData) Type() Type {
	return TypeEditMultisigOwners
}

func (d *EditMultisigOwnersData) Fee() Fee {
	return feeEditMultisigOwners
}

func (d *EditMultisigOwnersData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
