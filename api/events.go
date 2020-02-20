package api

import (
	"encoding/json"
	"errors"
	"strconv"
)

type EventsResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      string        `json:"id,omitempty"`
	Result  *EventsResult `json:"result,omitempty"`
	Error   *Error        `json:"error,omitempty"`
}

type EventsResult struct {
	Events []Event `json:"events"`
}

type Event struct {
	Type  string            `json:"type"`
	Value map[string]string `json:"value"`
}

// Converting event map data to the structure interface regarding event type
func (e *Event) ValueStruct() (interface{}, error) {
	bytes, err := json.Marshal(e.Value)
	if err != nil {
		return nil, err
	}

	var value interface{}
	switch e.Type {
	case "minter/RewardEvent":
		value = &RewardEventValue{}
	case "minter/SlashEvent":
		value = &SlashEventValue{}
	case "minter/UnbondEvent":
		value = &UnbondEventValue{}
	default:
		return nil, errors.New("unknown event type")
	}

	err = json.Unmarshal(bytes, value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

type RewardEventValue struct {
	Role            string `json:"role"`
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

type SlashEventValue struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Coin            string `json:"coin"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

type UnbondEventValue struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Coin            string `json:"coin"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

// Returns events at given height.
func (a *Api) Events(height int) (*EventsResult, error) {

	params := make(map[string]string)
	params["height"] = strconv.Itoa(height)

	res, err := a.client.R().SetQueryParams(params).Get("/events")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, NewResponseError(res)
	}

	response := new(EventsResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
