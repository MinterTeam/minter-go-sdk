package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
)

type SetCandidateOnData struct {
	PubKey []byte
}

func NewSetCandidateOnData() *SetCandidateOnData {
	return &SetCandidateOnData{}
}

func (d *SetCandidateOnData) SetPubKey(symbol string) *SetCandidateOnData {
	d.PubKey = []byte(symbol)
	return d
}

func (d *SetCandidateOnData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
