package transaction

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction for turning candidate off. This transaction should be sent from address which is set in the "Declare candidacy transaction".
// PubKey - Public key of a validator.
type SetCandidateOffData struct {
	PubKey []byte
}

func NewSetCandidateOffData() *SetCandidateOffData {
	return &SetCandidateOffData{}
}

func (d *SetCandidateOffData) SetPubKey(key string) (*SetCandidateOffData, error) {
	var err error
	d.PubKey, err = hex.DecodeString(key[2:])
	if err != nil {
		return d, err
	}
	return d, nil
}

func (d *SetCandidateOffData) MustSetPubKey(key string) *SetCandidateOffData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *SetCandidateOffData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
func (d *SetCandidateOffData) fee() Fee {
	return feeTypeSetCandidateOffline
}
