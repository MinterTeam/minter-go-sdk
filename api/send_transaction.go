package api

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

func (a *Api) Send(transaction transaction.SignedTransaction) ([]byte, error) {
	bytes, err := transaction.Encode()
	if err != nil {
		return nil, err
	}
	resp, err := a.client.R().Get(fmt.Sprintf("/send_transaction?tx=%s", bytes))
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
