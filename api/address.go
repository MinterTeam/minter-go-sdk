package api

import (
	"encoding/json"
	"strconv"
)

type AddressResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Balance          map[string]string `json:"balance"`
		TransactionCount string            `json:"transaction_count"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

func (a *Api) Address(address string) (*AddressResponse, error) {
	res, err := a.client.R().SetQueryParam("address", address).Get("/address")
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
func (a *Api) Balance(address string) (map[string]string, error) {
	response, err := a.Address(address)
	if err != nil {
		return nil, err
	}

	return response.Result.Balance, nil
}

func (a *Api) Nonce(address string) (uint64, error) {
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
