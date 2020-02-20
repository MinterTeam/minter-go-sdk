package api

import (
	"encoding/json"
	"strconv"
)

type AddressResponse struct {
	Jsonrpc string         `json:"jsonrpc"`
	ID      string         `json:"id,omitempty"`
	Result  *AddressResult `json:"result,omitempty"`
	Error   *Error         `json:"error,omitempty"`
}

type AddressResult struct {
	Balance          map[string]string `json:"balance"`
	TransactionCount string            `json:"transaction_count"`
}

// Returns coins list, balance and transaction count (for nonce) of an address.
func (a *Api) Address(address string, height int) (*AddressResult, error) {

	params := make(map[string]string)
	params["address"] = address
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/address")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, NewResponseError(res)
	}

	response := new(AddressResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}

// Returns balance of an address.
func (a *Api) Balance(address string, height int) (map[string]string, error) {
	response, err := a.Address(address, height)
	if err != nil {
		return nil, err
	}

	return response.Balance, nil
}

// Returns next transaction number (nonce) of an address.
func (a *Api) Nonce(address string) (uint64, error) {
	response, err := a.Address(address, LatestBlockHeight)
	if err != nil {
		return 0, err
	}

	nonce, err := strconv.ParseUint(response.TransactionCount, 10, 64)
	if err != nil {
		return 0, err
	}

	return nonce + 1, nil
}
