package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

type SetHaltBlockData struct {
	PubKey [32]byte
	Height uint64
}

func NewSetHaltBlockData() *SetHaltBlockData {
	return &SetHaltBlockData{}
}

// Set public key
func (d *SetHaltBlockData) SetPubKey(key string) (*SetHaltBlockData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// Tries to set public key and panics on error.
func (d *SetHaltBlockData) MustSetPubKey(key string) *SetHaltBlockData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Set height
func (d *SetHaltBlockData) SetHeight(height uint64) *SetHaltBlockData {
	d.Height = height
	return d
}

func (d *SetHaltBlockData) Type() Type {
	return TypeSetHaltBlock
}

func (d *SetHaltBlockData) Fee() Fee {
	return feeTypeSetHaltBlock
}

func (d *SetHaltBlockData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
