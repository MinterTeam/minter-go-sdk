package api

import (
	"encoding/json"
	"strconv"
)

type MissedBlocksResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  []struct {
		MissedBlocks      string `json:"missed_blocks"`
		MissedBlocksCount string `json:"missed_blocks_count"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

func (a *Api) MissedBlocks(pubKey string, height int) (*MissedBlocksResponse, error) {

	params := make(map[string]string)
	params["pub_key"] = pubKey
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/missed_blocks")
	if err != nil {
		return nil, err
	}

	result := new(MissedBlocksResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
