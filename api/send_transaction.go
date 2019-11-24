package api

import (
	"encoding/json"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

type SendTransactionResponse struct {
	Jsonrpc string                 `json:"jsonrpc"`
	ID      string                 `json:"id,omitempty"`
	Result  *SendTransactionResult `json:"result,omitempty"`
	Error   *TxError               `json:"error,omitempty"`
}

type SendTransactionResult struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Log  string `json:"log"`
	Hash string `json:"hash"`
}

type TxError struct {
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message,omitempty"`
	Data     string `json:"data,omitempty"`
	TxResult struct {
		Code int    `json:"code,omitempty"`
		Log  string `json:"log,omitempty"`
	} `json:"tx_result,omitempty"`
}

func (e *TxError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, data: \"%s\", tx_result.code: %d, tx_result.log: \"%s\"", e.Code, e.Message, e.Message, e.TxResult.Code, e.TxResult.Log)
}

// Returns the result of sending signed tx.
func (a *Api) SendTransaction(transaction transaction.SignedTransaction) (*SendTransactionResult, error) {
	tx, err := transaction.Encode()
	if err != nil {
		return nil, err
	}

	res, err := a.client.R().SetQueryParam("tx", tx).Get("/send_transaction")
	if err != nil {
		return nil, err
	}

	response := new(SendTransactionResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
