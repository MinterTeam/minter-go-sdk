package api

import "strconv"

type ValidatorsResponse struct {
}

func (a *Api) Validators(height int) (*ValidatorsResponse, error) {
	result := new(ValidatorsResponse)

	params := make(map[string]string)
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	_, err := a.client.R().SetResult(result).SetQueryParams(params).Get("/validators")
	if err != nil {
		return nil, err
	}

	return result, nil
}
