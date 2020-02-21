package api

import (
	"encoding/json"
)

type MinGasPriceResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id,omitempty"`
	Result  string `json:"result,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

// Returns current min gas price.
func (a *Api) MinGasPrice() (string, error) {

	res, err := a.client.R().Get("/min_gas_price")
	if err != nil {
		return "", err
	}
	if res.IsError() {
		return "", NewResponseError(res)
	}

	response := new(MinGasPriceResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return "", err
	}

	if response.Error != nil {
		return "", response.Error
	}

	return response.Result, nil
}
