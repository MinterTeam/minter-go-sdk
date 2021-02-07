package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
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
	VoteCommission          *big.Int
	VoteUpdate              *big.Int
	_                       []*big.Int `rlp:"tail"`
}

func NewVoteCommissionData() *VoteCommissionData {
	return &VoteCommissionData{}
}

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

// Fee returns commission of transaction Data
func (d *VoteCommissionData) Fee() Fee {
	return 0 // todo
}

// Encode returns the byte representation of a transaction Data.
func (d *VoteCommissionData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
