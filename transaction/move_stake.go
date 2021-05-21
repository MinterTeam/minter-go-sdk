package transaction

import (
	"math/big"

	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
)

type MoveStakeData struct {
	From,
	To PublicKey
	Coin  CoinID
	Stake *big.Int
}

// Type returns Data type of the transaction.
func (d *MoveStakeData) Type() Type {
	return TypeMoveStake
}

// Encode returns the byte representation of a transaction Data.
func (d *MoveStakeData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
