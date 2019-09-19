package api

import (
	"encoding/json"
	"strconv"
)

type UnconfirmedTxsResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		NTxs       string   `json:"n_txs"`
		Total      string   `json:"total"`
		TotalBytes string   `json:"total_bytes"`
		Txs        []string `json:"txs"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

func (a *Api) UnconfirmedTxs(limit int) (*UnconfirmedTxsResponse, error) {

	params := make(map[string]string)
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/unconfirmed_txs")
	if err != nil {
		return nil, err
	}

	result := new(UnconfirmedTxsResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
