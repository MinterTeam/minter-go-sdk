package api

import (
	"encoding/json"
)

type MinGasPriceResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

func (a *Api) MinGasPrice() (*MinGasPriceResponse, error) {

	res, err := a.client.R().Get("/min_gas_price")
	if err != nil {
		return nil, err
	}

	result := new(MinGasPriceResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
