package api

import (
	"encoding/json"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

type SendResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Code int    `json:"code"`
		Data string `json:"data"`
		Log  string `json:"log"`
		Hash string `json:"hash"`
	} `json:"result,omitempty"`
	Error struct {
		Code     int    `json:"code,omitempty"`
		Message  string `json:"message"`
		Data     string `json:"data"`
		TxResult struct {
			Code int    `json:"code"`
			Log  string `json:"log"`
		} `json:"tx_result"`
	} `json:"error,omitempty"`
}

// Returns the result of sending signed tx.
func (a *Api) Send(transaction transaction.SignedTransaction) (*SendResponse, error) {
	bytes, err := transaction.Encode()
	if err != nil {
		return nil, err
	}

	res, err := a.client.R().Get(fmt.Sprintf("/send_transaction?tx=%s", bytes))
	if err != nil {
		return nil, err
	}

	response := new(SendResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
