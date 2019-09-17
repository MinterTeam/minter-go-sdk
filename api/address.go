package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type AddressResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Balance          map[string]string `json:"balance"`
		TransactionCount string            `json:"transaction_count"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

func (a *Api) Address(address []byte) (*AddressResponse, error) {

	res, err := a.client.R().Get(fmt.Sprintf("/address?address=%s", address))
	if err != nil {
		return nil, err
	}

	result := new(AddressResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *Api) Nonce(address []byte) (uint64, error) {
	response, err := a.Address(address)
	if err != nil {
		return 0, err
	}

	nonce, err := strconv.ParseUint(response.Result.TransactionCount, 10, 64)
	if err != nil {
		return 0, err
	}

	return nonce + 1, nil
}
