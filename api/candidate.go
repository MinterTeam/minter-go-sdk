package api

import "strconv"

type CandidateResponse struct {
}

func (a *Api) Candidate(pubKey string, height int) (*CandidateResponse, error) {
	result := new(CandidateResponse)

	params := make(map[string]string)
	params["pub_key"] = pubKey
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	_, err := a.client.R().SetResult(result).SetQueryParams(params).Get("/candidate")
	if err != nil {
		return nil, err
	}

	return result, nil
}
