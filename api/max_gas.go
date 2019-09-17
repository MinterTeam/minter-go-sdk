package api

import (
	"encoding/json"
)

type MaxGasResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
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
