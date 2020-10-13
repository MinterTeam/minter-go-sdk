package api

import (
	"encoding/json"
	"fmt"
)

type EstimateTxCommissionResponse struct {
	Jsonrpc string                      `json:"jsonrpc"`
	ID      string                      `json:"id,omitempty"`
	Result  *EstimateTxCommissionResult `json:"result,omitempty"`
	Error   *Error                      `json:"error,omitempty"`
}

type EstimateTxCommissionResult struct {
	Commission string `json:"commission"`
}

// Return estimate of transaction.
func (a *Api) EstimateTxCommission(transaction interface{ Encode() (string, error) }) (*EstimateTxCommissionResult, error) {
	bytes, err := transaction.Encode()
	if err != nil {
		return nil, err
	}

	res, err := a.client.R().Get(fmt.Sprintf("/estimate_tx_commission?tx=%s", bytes))
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, NewResponseError(res)
	}

	response := new(EstimateTxCommissionResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
