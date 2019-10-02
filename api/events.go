package api

import (
	"encoding/json"
	"strconv"
)

type EventsResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Result  *EventsResult `json:"result,omitempty"`
	Error   *Error        `json:"error,omitempty"`
}

type EventsResult struct {
	Events []struct {
		Type  string `json:"type"`
		Value struct {
			Role            string `json:"role"`
			Address         string `json:"address"`
			Amount          string `json:"amount"`
			ValidatorPubKey string `json:"validator_pub_key"`
		} `json:"value"`
	} `json:"events"`
}

// Returns events at given height.
func (a *Api) Events(height int) (*EventsResult, error) {

	params := make(map[string]string)
	params["height"] = strconv.Itoa(height)

	res, err := a.client.R().SetQueryParams(params).Get("/events")
	if err != nil {
		return nil, err
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
