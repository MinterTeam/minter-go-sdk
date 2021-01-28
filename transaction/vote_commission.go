package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type VoteCommissionData struct {
	PubKey                  PublicKey
	Height                  uint64
	Coin                    CoinID
	PayloadByte             *big.Int
	Send                    *big.Int
	BuyBancor               *big.Int
	SellBancor              *big.Int
	SellAllBancor           *big.Int
	BuyPool                 *big.Int
	SellPool                *big.Int
	SellAllPool             *big.Int
	CreateTicker3           *big.Int
	CreateTicker4           *big.Int
	CreateTicker5           *big.Int
	CreateTicker6           *big.Int
	CreateTicker7to10       *big.Int
	CreateCoin              *big.Int
	CreateToken             *big.Int
	RecreateCoin            *big.Int
	RecreateToken           *big.Int
	DeclareCandidacy        *big.Int
	Delegate                *big.Int
	Unbond                  *big.Int
	RedeemCheck             *big.Int
	SetCandidateOn          *big.Int
	SetCandidateOff         *big.Int
	CreateMultisig          *big.Int
	MultisendBase           *big.Int
	MultisendDelta          *big.Int
	EditCandidate           *big.Int
	SetHaltBlock            *big.Int
	EditTickerOwner         *big.Int
	EditMultisig            *big.Int
	PriceVote               *big.Int
	EditCandidatePublicKey  *big.Int
	CreateSwapPool          *big.Int
	AddLiquidity            *big.Int
	RemoveLiquidity         *big.Int
	EditCandidateCommission *big.Int
	MoveStake               *big.Int
	MintToken               *big.Int
	BurnToken               *big.Int
	PriceCommission         *big.Int
	UpdateNetwork           *big.Int
	_                       []*big.Int `rlp:"tail"`
}

// Type returns Data type of the transaction.
func (d *VoteCommissionData) Type() Type {
	return TypeVoteCommission
}

// Fee returns commission of transaction Data
func (d *VoteCommissionData) Fee() Fee {
	return 0 // todo
}

// Encode returns the byte representation of a transaction Data.
func (d *VoteCommissionData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
