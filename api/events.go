package api

import (
	"encoding/json"
	"fmt"
)

// Event type names
const (
	TypeRewardEvent    = "minter/RewardEvent"
	TypeSlashEvent     = "minter/SlashEvent"
	TypeUnbondEvent    = "minter/UnbondEvent"
	TypeStakeKickEvent = "minter/StakeKickEvent"
)

// Event interface
type Event interface {
	GetAddress() string
	GetValidatorPublicKey() string
	clone() Event
}

// Events
type RewardEvent struct {
	Role            string `json:"role"`
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

func (e *RewardEvent) clone() Event {
	c := *e
	return &c
}

func (e *RewardEvent) GetAddress() string {
	return e.Address
}
func (e *RewardEvent) GetValidatorPublicKey() string {
	return e.ValidatorPubKey
}

type SlashEvent struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Coin            string `json:"coin"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

func (e *SlashEvent) clone() Event {
	c := *e
	return &c
}
func (e *SlashEvent) GetAddress() string {
	return e.Address
}
func (e *SlashEvent) GetValidatorPublicKey() string {
	return e.ValidatorPubKey
}

type UnbondEvent struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Coin            string `json:"coin"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

func (e *UnbondEvent) clone() Event {
	c := *e
	return &c
}
func (e *UnbondEvent) GetAddress() string {
	return e.Address
}
func (e *UnbondEvent) GetValidatorPublicKey() string {
	return e.ValidatorPubKey
}

type StakeKickEvent struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Coin            string `json:"coin"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

func (e *StakeKickEvent) clone() Event {
	c := *e
	return &c
}
func (e *StakeKickEvent) GetAddress() string {
	return e.Address
}
func (e *StakeKickEvent) GetValidatorPublicKey() string {
	return e.ValidatorPubKey
}

var em = map[string]Event{
	TypeRewardEvent:    &RewardEvent{},
	TypeSlashEvent:     &SlashEvent{},
	TypeUnbondEvent:    &UnbondEvent{},
	TypeStakeKickEvent: &StakeKickEvent{},
}

func ConvertToEvent(typeName string, value []byte) (Event, error) {
	eventStruct, ok := em[typeName]
	if !ok {
		return nil, fmt.Errorf("event type unknown: %s", typeName)
	}

	clone := eventStruct.clone()
	err := json.Unmarshal(value, clone)
	if err != nil {
		return nil, err
	}

	return clone, nil
}
