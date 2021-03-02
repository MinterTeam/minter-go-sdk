package api

import (
	"encoding/json"
)

type SubscribeResult struct {
	Query string `json:"query"`
	Data  struct {
		Height int `json:"height"`
		Result struct {
			Events []struct {
				Attributes []struct {
					Index bool   `json:"index"`
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
				Type string `json:"type"`
			} `json:"events"`
			GasUsed   int `json:"gas_used"`
			GasWanted int `json:"gas_wanted"`
		} `json:"result"`
		Tx string `json:"tx"`
	} `json:"data"`
	Events []struct {
		Key    string   `json:"key"`
		Events []string `json:"events"`
	} `json:"events"`
}

func FindNewBlockTags(message, tag string) (string, error) {
	var recv SubscribeResult
	err := json.Unmarshal([]byte(message), &recv)
	if err != nil {
		return "", err
	}

	for _, event := range recv.Events {
		if event.Key == tag {
			if len(event.Events) == 0 {
				break
			}
			return event.Events[0], nil
		}
	}

	return "", nil
}
