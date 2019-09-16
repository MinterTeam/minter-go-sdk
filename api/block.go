package api

import "strconv"

type BlockResponse struct {
}

func (a *Api) Block(height int) (*BlockResponse, error) {
	result := new(BlockResponse)

	params := make(map[string]string)
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	_, err := a.client.R().SetResult(result).SetQueryParams(params).Get("/block")
	if err != nil {
		return nil, err
	}

	return result, nil
}
