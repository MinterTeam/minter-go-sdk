package models

import (
	"encoding/json"
)

// EventItem is the structure of the EventsResponse view of list items
type EventItem struct {
	Type  string
	Value map[string]string
}

// NewEventItem returns an EventItem from the EventsResponse list item
func NewEventItem(i interface{}) *EventItem {
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil
	}

	var event EventItem
	err = json.Unmarshal(marshal, &event)
	if err != nil {
		return nil
	}

	return &event
}

// Event type names
const (
	TypeRewardEvent    = "minter/RewardEvent"
	TypeSlashEvent     = "minter/SlashEvent"
	TypeUnbondEvent    = "minter/UnbondEvent"
	TypeStakeKickEvent = "minter/StakeKickEvent"
)

// Events
type (
	RewardEvent struct {
		Role            string `json:"role"`
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}
	SlashEvent struct {
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		Coin            string `json:"coin"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}
	UnbondEvent struct {
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		Coin            string `json:"coin"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}
	StakeKickEvent struct {
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		Coin            string `json:"coin"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}
)
