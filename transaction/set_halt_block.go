package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// SetHaltBlockData is a Data of Transaction for voting to stop the network on block.
// This transaction should be sent from OwnerAddress which is set in the "Declare candidacy transaction".
type SetHaltBlockData struct {
	PubKey [32]byte
	Height uint64
}

// NewSetHaltBlockData returns new SetHaltBlockData of Transaction for voting to stop the network on block.
func NewSetHaltBlockData() *SetHaltBlockData {
	return &SetHaltBlockData{}
}

// SetPubKey sets public key.
func (d *SetHaltBlockData) SetPubKey(key string) (*SetHaltBlockData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// MustSetPubKey tries to set public key and panics on error.
func (d *SetHaltBlockData) MustSetPubKey(key string) *SetHaltBlockData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// SetHeight sets height
func (d *SetHaltBlockData) SetHeight(height uint64) *SetHaltBlockData {
	d.Height = height
	return d
}

// Type returns Data type of the transaction.
func (d *SetHaltBlockData) Type() Type {
	return TypeSetHaltBlock
}

// Fee returns commission of transaction Data
func (d *SetHaltBlockData) Fee() Fee {
	return feeTypeSetHaltBlock
}

// Encode returns the byte representation of a transaction Data.
func (d *SetHaltBlockData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
