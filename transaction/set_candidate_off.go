package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
)

type SetCandidateOffData struct {
	PubKey []byte
}

func NewSetCandidateOffData() *SetCandidateOffData {
	return &SetCandidateOffData{}
}

func (d *SetCandidateOffData) SetPubKey(symbol string) *SetCandidateOffData {
	d.PubKey = []byte(symbol)
	return d
}

func (d *SetCandidateOffData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
