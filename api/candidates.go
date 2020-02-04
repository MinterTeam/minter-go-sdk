package api

import (
	"encoding/json"
	"strconv"
)

type CandidatesResponse struct {
	Jsonrpc string             `json:"jsonrpc"`
	ID      string             `json:"id,omitempty"`
	Result  []*CandidateResult `json:"result,omitempty"`
	Error   *Error             `json:"error,omitempty"`
}

// Returns list of candidates.
func (a *Api) Candidates(height int, includeStakes bool) ([]*CandidateResult, error) {

	params := make(map[string]string)
	if includeStakes {
		params["include_stakes"] = "true"
	}
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/candidates")
	if err := hasError(res, err); err != nil {
		return nil, err
	}

	response := new(CandidatesResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
