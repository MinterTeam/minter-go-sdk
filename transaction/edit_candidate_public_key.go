package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// EditCandidatePublicKeyData is Data of Transaction for editing candidate public key.
// This transaction should be sent from OwnerAddress which is set in the "Declare candidacy transaction".
type EditCandidatePublicKeyData struct {
	PubKey    PublicKey // Public key of a validator
	NewPubKey PublicKey // New public key for change.
}

// NewEditCandidatePublicKeyData returns new EditCandidatePublicKeyData of Transaction for editing existing candidate.
func NewEditCandidatePublicKeyData() *EditCandidatePublicKeyData {
	return &EditCandidatePublicKeyData{}
}

// SetPubKey sets public key of a validator.
func (d *EditCandidatePublicKeyData) SetPubKey(key string) (*EditCandidatePublicKeyData, error) {
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
func (d *EditCandidatePublicKeyData) MustSetPubKey(key string) *EditCandidatePublicKeyData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// SetNewPubKey sets new public key for change.
func (d *EditCandidatePublicKeyData) SetNewPubKey(key string) (*EditCandidatePublicKeyData, error) {
	newPubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	var pubKey PublicKey
	copy(pubKey[:], newPubKey)
	d.NewPubKey = pubKey
	return d, nil
}

// MustSetNewPubKey tries to set new public key and panics on error.
func (d *EditCandidatePublicKeyData) MustSetNewPubKey(key string) *EditCandidatePublicKeyData {
	_, err := d.SetNewPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Type returns Data type of the transaction.
func (d *EditCandidatePublicKeyData) Type() Type {
	return TypeEditCandidatePublicKey
}

// Fee returns commission of transaction Data
func (d *EditCandidatePublicKeyData) Fee() Fee {
	return feeTypeEditCandidatePublicKey
}

// Encode returns the byte representation of a transaction Data.
func (d *EditCandidatePublicKeyData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
