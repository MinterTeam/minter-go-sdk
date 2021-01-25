package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type PriceCommissionData struct {
	PubKey                  PublicKey
	Height                  uint64
	Coin                    CoinID
	PayloadByte             *big.Int
	Send                    *big.Int
	Convert                 *big.Int
	CreateTicker3           *big.Int
	CreateTicker4           *big.Int
	CreateTicker5           *big.Int
	CreateTicker6           *big.Int
	CreateTicker7to10       *big.Int
	RecreateTicker          *big.Int
	DeclareCandidacy        *big.Int
	Delegate                *big.Int
	Unbond                  *big.Int
	RedeemCheck             *big.Int
	ToggleCandidateStatus   *big.Int
	CreateMultisig          *big.Int
	MultisendDelta          *big.Int
	EditCandidate           *big.Int
	SetHaltBlock            *big.Int
	EditTickerOwner         *big.Int
	EditMultisig            *big.Int
	PriceVote               *big.Int
	EditCandidatePublicKey  *big.Int
	AddLiquidity            *big.Int
	RemoveLiquidity         *big.Int
	EditCandidateCommission *big.Int
	MoveStake               *big.Int
	EditTokenEmission       *big.Int
	PriceCommission         *big.Int
	UpdateNetwork           *big.Int
	_                       []*big.Int `rlp:"tail"`
}

// Type returns Data type of the transaction.
func (d *PriceCommissionData) Type() Type {
	return TypePriceCommission
}

// Fee returns commission of transaction Data
func (d *PriceCommissionData) Fee() Fee {
	return 0 // todo
}

// Encode returns the byte representation of a transaction Data.
func (d *PriceCommissionData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
