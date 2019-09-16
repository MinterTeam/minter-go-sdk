package api

import "strconv"

type EventsResponse struct {
}

func (a *Api) Events(height int) (*EventsResponse, error) {
	result := new(EventsResponse)

	params := make(map[string]string)
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	_, err := a.client.R().SetResult(result).SetQueryParams(params).Get("/events")
	if err != nil {
		return nil, err
	}

	return result, nil
}
