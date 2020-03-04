package transaction

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction for turning candidate on. This transaction should be sent from address which is set in the "Declare candidacy transaction".
// PubKey - Public key of a validator.
type SetCandidateOnData struct {
	PubKey []byte
}

func NewSetCandidateOnData() *SetCandidateOnData {
	return &SetCandidateOnData{}
}

func (d *SetCandidateOnData) SetPubKey(key string) (*SetCandidateOnData, error) {
	var err error
	d.PubKey, err = hex.DecodeString(key[2:])
	if err != nil {
		return d, err
	}
	return d, nil
}

func (d *SetCandidateOnData) MustSetPubKey(key string) *SetCandidateOnData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}
func (d *SetCandidateOnData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
func (d *SetCandidateOnData) fee() fee {
	return feeTypeSetCandidateOnline
}
