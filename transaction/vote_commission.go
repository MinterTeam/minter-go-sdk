package transaction

import (
	"math/big"

	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// VoteCommissionData is a Data of Transaction for enabling validators to vote for the fees to be changed. The change comes into force once a two-thirds majority is reached. The vote can be sent from the validator owner address.
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
	VoteCommission          *big.Int
	VoteUpdate              *big.Int
	More                    []*big.Int `rlp:"tail"`
	//FailedTX         *big.Int
	//AddLimitOrder    *big.Int
	//RemoveLimitOrder *big.Int
	//MoveStake        *big.Int
	//LockStake        *big.Int
	//Lock             *big.Int
}

// NewVoteCommissionData creates VoteCommissionData
func NewVoteCommissionData() *VoteCommissionData {
	return &VoteCommissionData{}
}

func (d *VoteCommissionData) FailedTxPrice() *big.Int {
	if len(d.More) > 0 {
		return d.More[0]
	}
	return big.NewInt(0)
}

func (d *VoteCommissionData) AddLimitOrderPrice() *big.Int {
	if len(d.More) > 1 {
		return d.More[1]
	}
	return big.NewInt(0)
}

func (d *VoteCommissionData) RemoveLimitOrderPrice() *big.Int {
	if len(d.More) > 2 {
		return d.More[2]
	}
	return big.NewInt(0)
}

func (d *VoteCommissionData) MoveStakePrice() *big.Int {
	if len(d.More) > 3 {
		return d.More[3]
	}
	return big.NewInt(0)
}
func (d *VoteCommissionData) LockStakePrice() *big.Int {
	if len(d.More) > 4 {
		return d.More[4]
	}
	return big.NewInt(0)
}
func (d *VoteCommissionData) LockPrice() *big.Int {
	if len(d.More) > 5 {
		return d.More[5]
	}
	return big.NewInt(0)
}

// SetCoin sets coin for calculate commission price in base coin
func (d *VoteCommissionData) SetCoin(id uint64) *VoteCommissionData {
	d.Coin = CoinID(id)
	return d
}

// SetPubKey sets public key.
func (d *VoteCommissionData) SetPubKey(key string) (*VoteCommissionData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// MustSetPubKey tries to set public key and panics on error.
func (d *VoteCommissionData) MustSetPubKey(key string) *VoteCommissionData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// SetHeight sets height
func (d *VoteCommissionData) SetHeight(height uint64) *VoteCommissionData {
	d.Height = height
	return d
}

// Type returns Data type of the transaction.
func (d *VoteCommissionData) Type() Type {
	return TypeVoteCommission
}

// Encode returns the byte representation of a transaction Data.
func (d *VoteCommissionData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
