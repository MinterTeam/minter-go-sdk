package api

import (
	"encoding/json"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

type EstimateTxCommissionResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Commission string `json:"commission"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

func (a *Api) EstimateTxCommission(transaction transaction.SignedTransaction) (*EstimateTxCommissionResponse, error) {
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

	return response, nil
}
