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
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Volume         string `json:"volume"`
	Crr            int    `json:"crr"`
	ReserveBalance string `json:"reserve_balance"`
	MaxSupply      string `json:"max_supply"`
	OwnerAddress   string `json:"owner_address,omitempty"`
}

// Returns information about coin. Note: this method does not return information about base coins (MNT and BIP).
func (a *Api) CoinInfo(symbol string) (*CoinInfoResult, error) {
	return a.CoinInfoAtHeight(symbol, LatestBlockHeight)
}

// Returns information about coin. Note: this method does not return information about base coins (MNT and BIP).
func (a *Api) CoinInfoAtHeight(symbol string, height int) (*CoinInfoResult, error) {

	params := make(map[string]string)
	params["symbol"] = symbol
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/coin_info")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, NewResponseError(res)
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
