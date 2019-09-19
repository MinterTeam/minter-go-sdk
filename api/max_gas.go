package api

import (
	"encoding/json"
)

type MaxGasResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result,omitempty"`
	Error   struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

func (a *Api) MaxGas() (*MaxGasResponse, error) {

	res, err := a.client.R().Get("/max_gas")
	if err != nil {
		return nil, err
	}

	result := new(MaxGasResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
