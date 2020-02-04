package api

import (
	"encoding/json"
	"strconv"
)

type TransactionsResponse struct {
	Jsonrpc string               `json:"jsonrpc"`
	ID      string               `json:"id,omitempty"`
	Result  []*TransactionResult `json:"result,omitempty"`
	Error   *Error               `json:"error,omitempty"`
}

// Return transactions by query.
func (a *Api) Transactions(query string, page int, perPage int) ([]*TransactionResult, error) {

	params := make(map[string]string)
	params["query"] = query
	if page != 0 {
		params["page"] = strconv.Itoa(page)
	}
	if perPage != 0 {
		params["perPage"] = strconv.Itoa(perPage)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/transactions")
	if err := hasError(res, err); err != nil {
		return nil, err
	}

	response := new(TransactionsResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
