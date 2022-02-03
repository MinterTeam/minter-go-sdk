package transaction

import (
	"math/big"

	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

type MoveStakeData struct {
	FromPubKey,
	ToPubKey PublicKey
	Coin  CoinID
	Value *big.Int
}

func NewMoveStakeData() *MoveStakeData {
	return &MoveStakeData{}
}

// SetFrom sets Public key of current validator
func (d *MoveStakeData) SetFrom(key string) (*MoveStakeData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.FromPubKey[:], pubKey)
	return d, nil
}

// MustSetFrom tries to set public key of current validator and panics on error.
func (d *MoveStakeData) MustSetFrom(key string) *MoveStakeData {
	_, err := d.SetFrom(key)
	if err != nil {
		panic(err)
	}
	return d
}

// SetTo sets Public key of new validator
func (d *MoveStakeData) SetTo(key string) (*MoveStakeData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.ToPubKey[:], pubKey)
	return d, nil
}

// MustSetTo tries to set public key of new validator and panics on error.
func (d *MoveStakeData) MustSetTo(key string) *MoveStakeData {
	_, err := d.SetTo(key)
	if err != nil {
		panic(err)
	}
	return d
}

// SetCoin sets ID of coin to stake
func (d *MoveStakeData) SetCoin(id uint64) *MoveStakeData {
	d.Coin = CoinID(id)
	return d
}

// SetValue sets amount of coins to stake
func (d *MoveStakeData) SetValue(value *big.Int) *MoveStakeData {
	d.Value = value
	return d
}

// Type returns Data type of the transaction.
func (d *MoveStakeData) Type() Type {
	return TypeMoveStake
}

// Encode returns the byte representation of a transaction Data.
func (d *MoveStakeData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
