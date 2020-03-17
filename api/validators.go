package api

import (
	"encoding/json"
	"strconv"
)

type ValidatorsResponse struct {
	Jsonrpc string             `json:"jsonrpc"`
	ID      string             `json:"id,omitempty"`
	Result  []*ValidatorResult `json:"result,omitempty"`
	Error   *Error             `json:"error,omitempty"`
}

type ValidatorResult struct {
	PubKey      string `json:"pub_key"`
	VotingPower string `json:"voting_power"`
}

// Returns list of active validators.
func (a *Api) Validators() ([]*ValidatorResult, error) {
	return a.ValidatorsAtHeight(LatestBlockHeight)
}

// Returns list of active validators.
func (a *Api) ValidatorsAtHeight(height int) ([]*ValidatorResult, error) {
	return a.ValidatorsPage(height, 1, 100)
}

// Returns list of active validators with custom paging.
func (a *Api) ValidatorsPage(height, page, perPage int) ([]*ValidatorResult, error) {

	params := make(map[string]string)
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	params["page"] = strconv.Itoa(page)
	params["perPage"] = strconv.Itoa(perPage)

	res, err := a.client.R().SetQueryParams(params).Get("/validators")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, NewResponseError(res)
	}

	response := new(ValidatorsResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
