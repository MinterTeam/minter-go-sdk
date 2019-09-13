package api

import (
	"fmt"
	"strconv"
)

type AddressResponse struct {
	Result struct {
		Balance struct {
			MNT       string `json:"MNT"`
			KLM0VCOIN string `json:"KLM0VCOIN"`
		} `json:"balance"`
		TransactionCount string `json:"transaction_count"`
	} `json:"result"`
}

func (a *Api) GetAddress(address []byte) (*AddressResponse, error) {
	result := new(AddressResponse)
	_, err := a.client.R().SetResult(result).Get(fmt.Sprintf("/address?address=%s", address))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *Api) GetAddressNonce(address []byte) (uint64, error) { //todo: change receiving address to privateKey
	response, err := a.GetAddress(address)
	if err != nil {
		return 0, err
	}
	nonce, err := strconv.ParseUint(response.Result.TransactionCount, 10, 64)
	if err != nil {
		return 0, err
	}
	return nonce + 1, nil
}
