package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
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

// Fee returns commission of transaction Data
func (d *MoveStakeData) Fee() Fee {
	return feeTypeMoveStake
}

// Encode returns the byte representation of a transaction Data.
func (d *MoveStakeData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
