package api

import (
	"encoding/json"
	"fmt"
)

const (
	// DAOAddress is DAO's address for charging a 10% commission on all rewards
	DAOAddress = "Mx7f0fc21d932f38ca9444f61703174569066cfa50"

	// DevelopersAddress is developers' address for charging a 10% commission on all rewards
	DevelopersAddress = "Mx688568d9d70c57e71d0b9de6480afb0d317f885c"
)

const (
	// RoleValidator is validator role
	RoleValidator = "Validator"

	// RoleDelegator is delegator role
	RoleDelegator = "Delegator"

	// RoleDAO is DAO role
	RoleDAO = "DAO"

	// RoleDevelopers is developers role
	RoleDevelopers = "Developers"
)

// EventType is string name of events
type EventType string

// Event type names
const (
	TypeRewardEvent            EventType = "minter/RewardEvent"
	TypeSlashEvent             EventType = "minter/SlashEvent"
	TypeUnbondEvent            EventType = "minter/UnbondEvent"
	TypeStakeKickEvent         EventType = "minter/StakeKickEvent"
	TypeStakeMoveEvent         EventType = "minter/StakeMoveEvent"
	TypeUpdateNetworkEvent     EventType = "minter/UpdateNetworkEvent"
	TypeUpdateCommissionsEvent EventType = "minter/UpdateCommissionsEvent"
)

// Event interface
type StakeEvent interface {
	// GetAddress return owner address
	GetAddress() string
	// GetValidatorPublicKey return validator public key
	GetValidatorPublicKey() string
}

// Event interface
type Event interface {
	Type() EventType
}

type UpdateCommissionsEvent struct {
	Coin                    string `json:"coin"`
	PayloadByte             string `json:"payload_byte"`
	Send                    string `json:"send"`
	BuyBancor               string `json:"buy_bancor"`
	SellBancor              string `json:"sell_bancor"`
	SellAllBancor           string `json:"sell_all_bancor"`
	BuyPool                 string `json:"buy_pool"`
	SellPool                string `json:"sell_pool"`
	SellAllPool             string `json:"sell_all_pool"`
	CreateTicker3           string `json:"create_ticker3"`
	CreateTicker4           string `json:"create_ticker4"`
	CreateTicker5           string `json:"create_ticker5"`
	CreateTicker6           string `json:"create_ticker6"`
	CreateTicker7_10        string `json:"create_ticker7_10"`
	CreateCoin              string `json:"create_coin"`
	CreateToken             string `json:"create_token"`
	RecreateCoin            string `json:"recreate_coin"`
	RecreateToken           string `json:"recreate_token"`
	DeclareCandidacy        string `json:"declare_candidacy"`
	Delegate                string `json:"delegate"`
	Unbond                  string `json:"unbond"`
	RedeemCheck             string `json:"redeem_check"`
	SetCandidateOn          string `json:"set_candidate_on"`
	SetCandidateOff         string `json:"set_candidate_off"`
	CreateMultisig          string `json:"create_multisig"`
	MultisendDelta          string `json:"multisend_delta"`
	EditCandidate           string `json:"edit_candidate"`
	SetHaltBlock            string `json:"set_halt_block"`
	EditTickerOwner         string `json:"edit_ticker_owner"`
	EditMultisig            string `json:"edit_multisig"`
	PriceVote               string `json:"price_vote"`
	EditCandidatePublicKey  string `json:"edit_candidate_public_key"`
	AddLiquidity            string `json:"add_liquidity"`
	RemoveLiquidity         string `json:"remove_liquidity"`
	EditCandidateCommission string `json:"edit_candidate_commission"`
	MoveStake               string `json:"move_stake"`
	EditTokenEmission       string `json:"edit_token_emission"`
	PriceCommission         string `json:"price_commission"`
	UpdateNetwork           string `json:"update_network"`
}

func (e *UpdateCommissionsEvent) Type() EventType { return TypeUpdateCommissionsEvent }

type UpdateNetworkEvent struct {
	Version string `json:"version"`
}

func (e *UpdateNetworkEvent) Type() EventType { return TypeUpdateNetworkEvent }

// StakeMoveEvent ...
type StakeMoveEvent struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Coin            string `json:"coin"`
	ValidatorPubKey string `json:"validator_pub_key"`
	WaitList        bool   `json:"waitlist"`
}

// GetAddress return owner address
func (e *StakeMoveEvent) GetAddress() string {
	return e.Address
}

// GetValidatorPublicKey return validator public key
func (e *StakeMoveEvent) GetValidatorPublicKey() string {
	return e.ValidatorPubKey
}
func (e *StakeMoveEvent) Type() EventType { return TypeStakeMoveEvent }

// RewardEvent is the payment of rewards
type RewardEvent struct {
	Role            string `json:"role"`
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

// GetAddress return owner address
func (e *RewardEvent) GetAddress() string {
	return e.Address
}

// GetValidatorPublicKey return validator public key
func (e *RewardEvent) GetValidatorPublicKey() string {
	return e.ValidatorPubKey
}
func (e *RewardEvent) Type() EventType { return TypeRewardEvent }

// SlashEvent is the payment of the validator's penalty by this stake
type SlashEvent struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Coin            string `json:"coin"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

// GetAddress return owner address
func (e *SlashEvent) GetAddress() string {
	return e.Address
}

// GetValidatorPublicKey return validator public key
func (e *SlashEvent) GetValidatorPublicKey() string {
	return e.ValidatorPubKey
}
func (e *SlashEvent) Type() EventType { return TypeSlashEvent }

// UnbondEvent is the unbinding a stake from a validator
type UnbondEvent struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Coin            string `json:"coin"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

// GetAddress return owner address
func (e *UnbondEvent) GetAddress() string {
	return e.Address
}

// GetValidatorPublicKey return validator public key
func (e *UnbondEvent) GetValidatorPublicKey() string {
	return e.ValidatorPubKey
}
func (e *UnbondEvent) Type() EventType { return TypeUnbondEvent }

// StakeKickEvent is the knocking out a stake to the waiting list
type StakeKickEvent struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Coin            string `json:"coin"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

// GetAddress return owner address
func (e *StakeKickEvent) GetAddress() string {
	return e.Address
}

// GetValidatorPublicKey return validator public key
func (e *StakeKickEvent) GetValidatorPublicKey() string {
	return e.ValidatorPubKey
}
func (e *StakeKickEvent) Type() EventType { return TypeStakeKickEvent }

func newEvent(t EventType) Event {
	switch t {
	case TypeRewardEvent:
		return &RewardEvent{}
	case TypeSlashEvent:
		return &SlashEvent{}
	case TypeUnbondEvent:
		return &UnbondEvent{}
	case TypeStakeKickEvent:
		return &StakeKickEvent{}
	case TypeStakeMoveEvent:
		return &StakeMoveEvent{}
	case TypeUpdateCommissionsEvent:
		return &UpdateCommissionsEvent{}
	case TypeUpdateNetworkEvent:
		return &UpdateNetworkEvent{}
	default:
		return nil
	}
}

// ConvertToEvent returns interface of Event
func ConvertToEvent(typeName EventType, value []byte) (Event, error) {
	eventStruct := newEvent(typeName)
	if eventStruct == nil {
		return nil, fmt.Errorf("Type type unknown: %s", typeName)
	}

	err := json.Unmarshal(value, eventStruct)
	if err != nil {
		return nil, err
	}

	return eventStruct, nil
}
