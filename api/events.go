package api

import (
	"encoding/json"
	"fmt"
)

var (
	// DAOAddress is DAO's address for charging a 10% commission on all rewards
	DAOAddress = "Mx7f0fc21d932f38ca9444f61703174569066cfa50"

	// DevelopersAddress is developers' address for charging a 10% commission on all rewards
	DevelopersAddress = "Mx688568d9d70c57e71d0b9de6480afb0d317f885c"
)
var (
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
	TypeRewardEvent    EventType = "minter/RewardEvent"
	TypeSlashEvent     EventType = "minter/SlashEvent"
	TypeUnbondEvent    EventType = "minter/UnbondEvent"
	TypeStakeKickEvent EventType = "minter/StakeKickEvent"
)

// Event interface
type Event interface {
	// GetAddress return owner address
	GetAddress() string
	// GetValidatorPublicKey return validator public key
	GetValidatorPublicKey() string
	event()
}

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
func (e *RewardEvent) event() {}

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
func (e *SlashEvent) event() {}

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
func (e *UnbondEvent) event() {}

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
func (e *StakeKickEvent) event() {}

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
	default:
		return nil
	}
}

// ConvertToEvent returns interface of Event
func ConvertToEvent(typeName EventType, value []byte) (Event, error) {
	eventStruct := newEvent(typeName)
	if eventStruct == nil {
		return nil, fmt.Errorf("event type unknown: %s", typeName)
	}

	err := json.Unmarshal(value, eventStruct)
	if err != nil {
		return nil, err
	}

	return eventStruct, nil
}
