package api

import (
	"encoding/json"
	"strconv"
)

type UnconfirmedTxsResponse struct {
	Jsonrpc string                `json:"jsonrpc"`
	ID      string                `json:"id,omitempty"`
	Result  *UnconfirmedTxsResult `json:"result,omitempty"`
	Error   *Error                `json:"error,omitempty"`
}

type UnconfirmedTxsResult struct {
	NTxs       string   `json:"n_txs"`
	Total      string   `json:"total"`
	TotalBytes string   `json:"total_bytes"`
	Txs        []string `json:"txs"`
}

// Returns unconfirmed transactions.
func (a *Api) UnconfirmedTxs(limit int) (*UnconfirmedTxsResult, error) {

	params := make(map[string]string)
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/unconfirmed_txs")
	if err := hasError(res, err); err != nil {
		return nil, err
	}

	response := new(UnconfirmedTxsResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
