package api

import (
	"encoding/json"
	"strconv"
)

type CoinInfoResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      string          `json:"id,omitempty"`
	Result  *CoinInfoResult `json:"result,omitempty"`
	Error   *Error          `json:"error,omitempty"`
}

type CoinInfoResult struct {
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Volume         string `json:"volume"`
	Crr            string `json:"crr"`
	ReserveBalance string `json:"reserve_balance"`
}

// Returns information about coin. Note: this method does not return information about base coins (MNT and BIP).
func (a *Api) CoinInfo(symbol string, height int) (*CoinInfoResult, error) {

	params := make(map[string]string)
	params["symbol"] = symbol
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/coin_info")
	if err != nil {
		return nil, err
	}

	response := new(CoinInfoResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
