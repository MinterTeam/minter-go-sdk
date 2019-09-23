package api

import (
	"encoding/json"
	"strconv"
)

type EventsResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Events []struct {
			Type  string `json:"type"`
			Value struct {
				Role            string `json:"role"`
				Address         string `json:"address"`
				Amount          string `json:"amount"`
				ValidatorPubKey string `json:"validator_pub_key"`
			} `json:"value"`
		} `json:"events"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

// Returns events at given height.
func (a *Api) Events(height int) (*EventsResponse, error) {

	params := make(map[string]string)
	params["height"] = strconv.Itoa(height)

	res, err := a.client.R().SetQueryParams(params).Get("/events")
	if err != nil {
		return nil, err
	}

	result := new(EventsResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
