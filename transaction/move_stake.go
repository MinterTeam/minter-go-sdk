package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type MoveStakeData struct {
	From,
	To PublicKey
	Coin  CoinID
	Stake *big.Int
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
	copy(d.From[:], pubKey)
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
	copy(d.To[:], pubKey)
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

// SetStake sets amount of coins to stake
func (d *MoveStakeData) SetStake(value *big.Int) *MoveStakeData {
	d.Stake = value
	return d
}

// Type returns Data type of the transaction.
func (d *MoveStakeData) Type() Type {
	return TypeMoveStake
}

// Fee returns commission of transaction Data
func (d *MoveStakeData) Fee() Fee {
	return feeTypeMoveStake
}

// Encode returns the byte representation of a transaction Data.
func (d *MoveStakeData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
