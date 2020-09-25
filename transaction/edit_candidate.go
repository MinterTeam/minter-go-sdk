package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// EditCandidateData is Data of Transaction for editing existing candidate.
// This transaction should be sent from OwnerAddress which is set in the "Declare candidacy transaction".
type EditCandidateData struct {
	PubKey         PublicKey // Public key of a validator
	RewardAddress  Address   // Address where validator’s rewards go to.
	OwnerAddress   Address   // Address that allows one to start the validator by sending the SetCandidateOnline transaction or stop it by sending the SetCandidateOffline transaction. It also enables the owner to edit the node by sending EditCandidate.
	ControlAddress Address   // Address that allows one to start the validator by sending the SetCandidateOnline transaction or stop it by sending the SetCandidateOffline transaction.
}

// NewEditCandidateData returns new EditCandidateData of Transaction for editing existing candidate.
func NewEditCandidateData() *EditCandidateData {
	return &EditCandidateData{}
}

// SetPubKey sets public key of a validator.
func (d *EditCandidateData) SetPubKey(key string) (*EditCandidateData, error) {
	pk, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	var pubKey PublicKey
	copy(pubKey[:], pk)
	d.PubKey = pubKey
	return d, nil
}

// MustSetPubKey tries to set public key of validator and panics on error.
func (d *EditCandidateData) MustSetPubKey(key string) *EditCandidateData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// MustSetRewardAddress tries to set reward address of validator and panics on error.
func (d *EditCandidateData) MustSetRewardAddress(address string) *EditCandidateData {
	_, err := d.SetRewardAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

// SetRewardAddress sets address where validator’s rewards go to.
func (d *EditCandidateData) SetRewardAddress(address string) (*EditCandidateData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.RewardAddress[:], bytes)
	return d, nil
}

// MustSetOwnerAddress tries to set owner address of validator and panics on error.
func (d *EditCandidateData) MustSetOwnerAddress(address string) *EditCandidateData {
	_, err := d.SetOwnerAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

// SetOwnerAddress sets address for managing SetCandidateOnline, SetCandidateOffline and EditCandidate data of transaction
func (d *EditCandidateData) SetOwnerAddress(address string) (*EditCandidateData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.OwnerAddress[:], bytes)
	return d, nil
}

// MustSetControlAddress tries to set control address of validator and panics on error.
func (d *EditCandidateData) MustSetControlAddress(address string) *EditCandidateData {
	_, err := d.SetControlAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

// SetControlAddress sets address for managing SetCandidateOnline and SetCandidateOffline data of transaction
func (d *EditCandidateData) SetControlAddress(address string) (*EditCandidateData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.ControlAddress[:], bytes)
	return d, nil
}

// Type returns Data type of the transaction.
func (d *EditCandidateData) Type() Type {
	return TypeEditCandidate
}

// Fee returns commission of transaction Data
func (d *EditCandidateData) Fee() Fee {
	return feeTypeEditCandidate
}

// Encode returns the byte representation of a transaction Data.
func (d *EditCandidateData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
