package api

import (
	"encoding/json"
	"strconv"
)

type TransactionsResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Result  []Transaction `json:"result,omitempty"`
	Error   struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

// Return transactions by query.
func (a *Api) Transactions(query string, page int, perPage int) (*TransactionsResponse, error) {

	params := make(map[string]string)
	params["query"] = query
	if page != 0 {
		params["page"] = strconv.Itoa(page)
	}
	if perPage != 0 {
		params["perPage"] = strconv.Itoa(perPage)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/transactions")
	if err != nil {
		return nil, err
	}

	result := new(TransactionsResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
