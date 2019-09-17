package api

import (
	"encoding/json"
	"strconv"
)

type CandidateResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
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
		CreatedAtBlock string `json:"created_at_block"`
		Status         int    `json:"status"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

func (a *Api) Candidate(pubKey string, height int) (*CandidateResponse, error) {

	params := make(map[string]string)
	params["pub_key"] = pubKey
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/candidate")
	if err != nil {
		return nil, err
	}

	result := new(CandidateResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
