package transaction

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
)

// RedeemCheckData is a Data of Transaction for redeeming a check.
// Note that maximum GasPrice is limited to 1 to prevent fraud,
// because GasPrice is set by redeem tx sender but commission is charded from check issuer.
type RedeemCheckData struct {
	RawCheck []byte   // Check received from sender
	Proof    [65]byte // Proof of owning a check: password signed with recipient's address
}

// NewRedeemCheckData returns new RedeemCheckData of Transaction for redeeming a check.
func NewRedeemCheckData() *RedeemCheckData {
	return &RedeemCheckData{}
}

// SetRawCheck sets check received from sender.
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

// MustSetRawCheck tries to set check received from sender and panics on error.
func (d *RedeemCheckData) MustSetRawCheck(raw string) *RedeemCheckData {
	_, err := d.SetRawCheck(raw)
	if err != nil {
		panic(err)
	}
	return d
}

// SetProof sets proof of owning a check.
func (d *RedeemCheckData) SetProof(proof string) (*RedeemCheckData, error) {
	bytes, err := hex.DecodeString(proof)
	if err != nil {
		return nil, err
	}
	copy(d.Proof[:], bytes)
	return d, nil
}

// MustSetProof tries to set proof of owning a check and panics on error.
func (d *RedeemCheckData) MustSetProof(proof string) *RedeemCheckData {
	_, err := d.SetProof(proof)
	if err != nil {
		panic(err)
	}
	return d
}

// Type returns Data type of the transaction.
func (d *RedeemCheckData) Type() Type {
	return TypeRedeemCheck
}

// Encode returns the byte representation of a transaction Data.
func (d *RedeemCheckData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
