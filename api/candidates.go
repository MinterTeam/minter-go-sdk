package api

import "strconv"

type CandidatesResponse struct {
}

func (a *Api) Candidates(height int, includeStakes bool) (*CandidateResponse, error) {
	result := new(CandidateResponse)

	params := make(map[string]string)
	if includeStakes {
		params["include_stakes"] = "true"
	}
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	_, err := a.client.R().SetResult(result).SetQueryParams(params).Get("/candidates")
	if err != nil {
		return nil, err
	}

	return result, nil
}
