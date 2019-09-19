package api

import (
	"encoding/json"
	"strconv"
)

type CoinInfoResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Name           string `json:"name"`
		Symbol         string `json:"symbol"`
		Volume         string `json:"volume"`
		Crr            string `json:"crr"`
		ReserveBalance string `json:"reserve_balance"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

func (a *Api) CoinInfo(symbol string, height int) (*CoinInfoResponse, error) {

	params := make(map[string]string)
	params["symbol"] = symbol
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/coin_info")
	if err != nil {
		return nil, err
	}

	result := new(CoinInfoResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
