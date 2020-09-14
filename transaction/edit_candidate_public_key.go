package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction for editing candidate public key.
type EditCandidatePublicKeyData struct {
	PubKey    [32]byte // Public key of a validator
	NewPubKey [32]byte // New public key for change.
}

// New data of transaction for editing existing candidate.
func NewEditCandidatePublicKeyData() *EditCandidatePublicKeyData {
	return &EditCandidatePublicKeyData{}
}

// Set public key of a validator.
func (d *EditCandidatePublicKeyData) SetPubKey(key string) (*EditCandidatePublicKeyData, error) {
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
func (d *EditCandidatePublicKeyData) MustSetPubKey(key string) *EditCandidatePublicKeyData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Set new public key for change.
func (d *EditCandidatePublicKeyData) SetNewPubKey(key string) (*EditCandidatePublicKeyData, error) {
	newPubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	var pubKey [32]byte
	copy(pubKey[:], newPubKey)
	d.NewPubKey = pubKey
	return d, nil
}

// Tries to set new public key and panics on error.
func (d *EditCandidatePublicKeyData) MustSetNewPubKey(key string) *EditCandidatePublicKeyData {
	_, err := d.SetNewPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *EditCandidatePublicKeyData) Type() Type {
	return TypeEditCandidatePublicKey
}

func (d *EditCandidatePublicKeyData) Fee() Fee {
	return feeTypeEditCandidatePublicKey
}

func (d *EditCandidatePublicKeyData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
