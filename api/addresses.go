package api

import (
	"encoding/json"
	"strconv"
	"strings"
)

type AddressesResponse struct {
	Jsonrpc string             `json:"jsonrpc"`
	ID      string             `json:"id,omitempty"`
	Result  []*AddressesResult `json:"result,omitempty"`
	Error   *Error             `json:"error,omitempty"`
}

type AddressesResult struct {
	Address          string    `json:"address"`
	Balance          []Balance `json:"balance"`
	TransactionCount string    `json:"transaction_count"`
}

func (a *Api) Addresses(addresses []string, height int) ([]*AddressesResult, error) {
	params := make(map[string]string)
	params["height"] = strconv.Itoa(height)
	params["addresses"] = "[" + strings.Join(addresses, ",") + "]"

	res, err := a.client.R().SetQueryParams(params).Get("/addresses")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, NewResponseError(res)
	}

	response := new(AddressesResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
