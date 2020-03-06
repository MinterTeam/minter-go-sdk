package api

import (
	"encoding/json"
	"strconv"
	"time"
)

type BlockResponse struct {
	Jsonrpc string       `json:"jsonrpc"`
	ID      string       `json:"id,omitempty"`
	Result  *BlockResult `json:"result,omitempty"`
	Error   *Error       `json:"error,omitempty"`
}

type BlockResult struct {
	Hash         string              `json:"hash"`
	Height       string              `json:"height"`
	Time         time.Time           `json:"time"`
	NumTxs       string              `json:"num_txs"`
	TotalTxs     string              `json:"total_txs"`
	Transactions []TransactionResult `json:"transactions"`
	BlockReward  string              `json:"block_reward"`
	Size         string              `json:"size"`
	Proposer     string              `json:"proposer"`
	Validators   []struct {
		PubKey string `json:"pub_key"`
		Signed bool   `json:"signed"`
	} `json:"validators"`
}

// Returns block data at given height.
func (a *Api) Block(height int) (*BlockResult, error) {

	params := make(map[string]string)
	params["height"] = strconv.Itoa(height)

	res, err := a.client.R().SetQueryParams(params).Get("/block")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, NewResponseError(res)
	}

	response := new(BlockResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
