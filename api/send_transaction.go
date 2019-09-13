package api

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"io/ioutil"
	"net/http"
)

func (a *Api) Send(transaction transaction.SignedTransaction) ([]byte, error) {
	bytes, err := transaction.Encode()
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(fmt.Sprintf(a.hostUrl+"/send_transaction?tx=%s", bytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
