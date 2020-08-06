package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction for editing existing candidate.
type EditCandidateData struct {
	PubKey         [32]byte  // Public key of a validator
	NewPubKey      *[32]byte `rlp:"nil"` // New public key for change.
	RewardAddress  [20]byte  // Address where validator’s rewards go to.
	OwnerAddress   [20]byte  // Address that allows one to start the validator by sending the SetCandidateOnline transaction or stop it by sending the SetCandidateOffline transaction. It also enables the owner to edit the node by sending EditCandidate.
	ControlAddress [20]byte  // Address that allows one to start the validator by sending the SetCandidateOnline transaction or stop it by sending the SetCandidateOffline transaction.
}

// New data of transaction for editing existing candidate.
func NewEditCandidateData() *EditCandidateData {
	return &EditCandidateData{}
}

// Set public key of a validator.
func (d *EditCandidateData) SetPubKey(key string) (*EditCandidateData, error) {
	pk, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	var pubKey [32]byte
	copy(pubKey[:], pk)
	d.PubKey = pubKey
	return d, nil
}

// Tries to set public key of validator and panics on error.
func (d *EditCandidateData) MustSetPubKey(key string) *EditCandidateData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Set new public key for change.
func (d *EditCandidateData) SetNewPubKey(key string) (*EditCandidateData, error) {
	newPubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	var pubKey [32]byte
	copy(pubKey[:], newPubKey)
	d.NewPubKey = &pubKey
	return d, nil
}

// Tries to set new public key and panics on error.
func (d *EditCandidateData) MustSetNewPubKey(key string) *EditCandidateData {
	_, err := d.SetNewPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Tries to set reward address of validator and panics on error.
func (d *EditCandidateData) MustSetRewardAddress(address string) *EditCandidateData {
	_, err := d.SetRewardAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

// Set address where validator’s rewards go to.
func (d *EditCandidateData) SetRewardAddress(address string) (*EditCandidateData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.RewardAddress[:], bytes)
	return d, nil
}

// Tries to set owner address of validator and panics on error.
func (d *EditCandidateData) MustSetOwnerAddress(address string) *EditCandidateData {
	_, err := d.SetOwnerAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

// Set address for managing SetCandidateOnline, SetCandidateOffline and EditCandidate data of transaction
func (d *EditCandidateData) SetOwnerAddress(address string) (*EditCandidateData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.OwnerAddress[:], bytes)
	return d, nil
}

// Tries to set control address of validator and panics on error.
func (d *EditCandidateData) MustSetControlAddress(address string) *EditCandidateData {
	_, err := d.SetControlAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

// Set address for managing SetCandidateOnline and SetCandidateOffline data of transaction
func (d *EditCandidateData) SetControlAddress(address string) (*EditCandidateData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.ControlAddress[:], bytes)
	return d, nil
}

func (d *EditCandidateData) Type() Type {
	return TypeEditCandidate
}

func (d *EditCandidateData) Fee() Fee {
	return feeTypeEditCandidate
}

func (d *EditCandidateData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
