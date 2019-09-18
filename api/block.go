package api

import (
	"encoding/json"
	"strconv"
	"time"
)

type BlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Hash         string        `json:"hash"`
		Height       string        `json:"height"`
		Time         time.Time     `json:"time"`
		NumTxs       string        `json:"num_txs"`
		TotalTxs     string        `json:"total_txs"`
		Transactions []Transaction `json:"transactions"`
		BlockReward  string        `json:"block_reward"`
		Size         string        `json:"size"`
		Proposer     string        `json:"proposer"`
		Validators   []struct {
			PubKey string `json:"pub_key"`
			Signed bool   `json:"signed"`
		} `json:"validators"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

func (a *Api) Block(height int) (*BlockResponse, error) {

	params := make(map[string]string)
	params["height"] = strconv.Itoa(height)

	res, err := a.client.R().SetQueryParams(params).Get("/block")
	if err != nil {
		return nil, err
	}

	result := new(BlockResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
