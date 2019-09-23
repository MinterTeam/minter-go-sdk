package api

import (
	"encoding/json"
	"strconv"
)

type ValidatorsResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  []struct {
		PubKey      string `json:"pub_key"`
		VotingPower string `json:"voting_power"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

// Returns list of active validators.
func (a *Api) Validators(height int) (*ValidatorsResponse, error) {

	params := make(map[string]string)
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/validators")
	if err != nil {
		return nil, err
	}

	result := new(ValidatorsResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
