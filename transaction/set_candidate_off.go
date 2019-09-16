package transaction

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
)

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
