package api

import (
	"encoding/json"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
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
func (a *Api) EstimateTxCommission(transaction transaction.SignedTransaction) (*EstimateTxCommissionResult, error) {
	bytes, err := transaction.Encode()
	if err != nil {
		return nil, err
	}

	res, err := a.client.R().Get(fmt.Sprintf("/estimate_tx_commission?tx=%s", bytes))
	if err != nil {
		return nil, err
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
