package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

type VoteUpdateData struct {
	Version string
	PubKey  PublicKey
	Height  uint64
}

func NewVoteUpdateData() *VoteUpdateData {
	return &VoteUpdateData{}
}

// SetVersion sets version name.
func (d *VoteUpdateData) SetVersion(version string) *VoteUpdateData {
	d.Version = version
	return d
}

// SetPubKey sets public key.
func (d *VoteUpdateData) SetPubKey(key string) (*VoteUpdateData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// MustSetPubKey tries to set public key and panics on error.
func (d *VoteUpdateData) MustSetPubKey(key string) *VoteUpdateData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// SetHeight sets height
func (d *VoteUpdateData) SetHeight(height uint64) *VoteUpdateData {
	d.Height = height
	return d
}

// Type returns Data type of the transaction.
func (d *VoteUpdateData) Type() Type {
	return TypeVoteUpdate
}

// Fee returns commission of transaction Data
func (d *VoteUpdateData) Fee() Fee {
	return 0 // todo
}

// Encode returns the byte representation of a transaction Data.
func (d *VoteUpdateData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
