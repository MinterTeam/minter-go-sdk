package transaction

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
)

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
	return d, err
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
