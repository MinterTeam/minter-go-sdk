package api

import (
	"encoding/json"
	"strconv"
)

type CandidateResponse struct {
	Jsonrpc string           `json:"jsonrpc"`
	ID      string           `json:"id,omitempty"`
	Result  *CandidateResult `json:"result,omitempty"`
	Error   *Error           `json:"error,omitempty"`
}

type CandidateResult struct {
	RewardAddress string `json:"reward_address"`
	OwnerAddress  string `json:"owner_address"`
	TotalStake    string `json:"total_stake"`
	PubKey        string `json:"pub_key"`
	Commission    string `json:"commission"`
	Stakes        []struct {
		Owner    string `json:"owner"`
		Coin     string `json:"coin"`
		Value    string `json:"value"`
		BipValue string `json:"bip_value"`
	} `json:"stakes"`
	Status int `json:"status"`
}

// Returns candidate’s info by provided public_key. It will respond with 404 code if candidate is not found.
func (a *Api) Candidate(pubKey string, height int) (*CandidateResult, error) {

	params := make(map[string]string)
	params["pub_key"] = pubKey
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/candidate")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, NewResponseError(res)
	}

	response := new(CandidateResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
