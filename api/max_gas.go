package api

import (
	"encoding/json"
)

type MaxGasResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id,omitempty"`
	Result  string `json:"result,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

// Returns current max gas.
func (a *Api) MaxGas() (string, error) {

	res, err := a.client.R().Get("/max_gas")
	if err := hasError(res, err); err != nil {
		return "", err
	}

	response := new(MaxGasResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return "", err
	}

	if response.Error != nil {
		return "", response.Error
	}

	return response.Result, nil
}
