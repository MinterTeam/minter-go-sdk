package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
)

type RedeemCheckData struct {
	RawCheck []byte
	Proof    [65]byte
}

func NewRedeemCheckData() *RedeemCheckData {
	return &RedeemCheckData{}
}

func (d *RedeemCheckData) SetRawCheck(raw []byte) *RedeemCheckData {
	d.RawCheck = raw
	return d
}
func (d *RedeemCheckData) SetProof(proof string) *RedeemCheckData {
	copy(d.Proof[:], proof)
	return d
}

func (d *RedeemCheckData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
