package api

import (
	"encoding/json"
	"strconv"
)

type MissedBlocksResponse struct {
	Jsonrpc string              `json:"jsonrpc"`
	ID      string              `json:"id,omitempty"`
	Result  *MissedBlocksResult `json:"result,omitempty"`
	Error   *Error              `json:"error,omitempty"`
}

type MissedBlocksResult struct {
	MissedBlocks      string `json:"missed_blocks"`
	MissedBlocksCount string `json:"missed_blocks_count"`
}

// Returns missed blocks by validator public key.
func (a *Api) MissedBlocks(pubKey string) (*MissedBlocksResult, error) {
	return a.MissedBlocksAtHeight(pubKey, LatestBlockHeight)
}

// Returns missed blocks by validator public key.
func (a *Api) MissedBlocksAtHeight(pubKey string, height int) (*MissedBlocksResult, error) {

	params := make(map[string]string)
	params["pub_key"] = pubKey
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/missed_blocks")
	if err != nil {
		return nil, err
	}

	response := new(MissedBlocksResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
