package api

type TransactionResponse struct {
}

func (a *Api) Transaction(hash string) (*TransactionResponse, error) {
	result := new(TransactionResponse)

	params := make(map[string]string)
	if hash != "" {
		params["hash"] = hash
	}

	_, err := a.client.R().SetResult(result).SetQueryParams(params).Get("/transaction")
	if err != nil {
		return nil, err
	}

	return result, nil
}
