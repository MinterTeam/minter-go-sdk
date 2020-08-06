package transaction

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/rlp"
	"strings"
)

// Transaction data for redeeming a check.
// Note that maximum GasPrice is limited to 1 to prevent fraud,
// because GasPrice is set by redeem tx sender but commission is charded from check issuer.
type RedeemCheckData struct {
	RawCheck []byte   // Check received from sender
	Proof    [65]byte // Proof of owning a check: password signed with recipient's address
}

// New data of transaction for redeeming a check.
func NewRedeemCheckData() *RedeemCheckData {
	return &RedeemCheckData{}
}

// Set check received from sender.
func (d *RedeemCheckData) SetRawCheck(raw string) (*RedeemCheckData, error) {
	raw = strings.Title(strings.ToLower(raw))
	if !strings.HasPrefix(raw, "Mc") {
		return nil, errors.New("raw check don't has prefix 'Mc'")
	}

	var err error
	d.RawCheck, err = hex.DecodeString(raw[2:])
	if err != nil {
		return d, err
	}
	return d, nil
}

// Tries to set check received from sender and panics on error.
func (d *RedeemCheckData) MustSetRawCheck(raw string) *RedeemCheckData {
	_, err := d.SetRawCheck(raw)
	if err != nil {
		panic(err)
	}
	return d
}

// Set proof of owning a check.
func (d *RedeemCheckData) SetProof(proof string) (*RedeemCheckData, error) {
	bytes, err := hex.DecodeString(proof)
	if err != nil {
		return nil, err
	}
	copy(d.Proof[:], bytes)
	return d, nil
}

// Tries to set proof of owning a check and panics on error.
func (d *RedeemCheckData) MustSetProof(proof string) *RedeemCheckData {
	_, err := d.SetProof(proof)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *RedeemCheckData) Type() Type {
	return TypeRedeemCheck
}

func (d *RedeemCheckData) Fee() Fee {
	return feeTypeRedeemCheck
}

func (d *RedeemCheckData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
