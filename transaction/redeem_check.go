package transaction

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
)

type RedeemCheckData struct {
	RawCheck []byte
	Proof    [65]byte
}

func NewRedeemCheckData() *RedeemCheckData {
	return &RedeemCheckData{}
}

func (d *RedeemCheckData) SetRawCheck(raw string) (*RedeemCheckData, error) {
	var err error
	d.RawCheck, err = hex.DecodeString(raw[2:])
	if err != nil {
		return d, err
	}
	return d, nil
}

func (d *RedeemCheckData) MustSetRawCheck(raw string) *RedeemCheckData {
	_, err := d.SetRawCheck(raw)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *RedeemCheckData) SetProof(proof string) (*RedeemCheckData, error) {
	bytes, err := hex.DecodeString(proof)
	if err != nil {
		return nil, err
	}
	copy(d.Proof[:], bytes)
	return d, nil
}

func (d *RedeemCheckData) MustSetProof(proof string) *RedeemCheckData {
	_, err := d.SetProof(proof)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *RedeemCheckData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
