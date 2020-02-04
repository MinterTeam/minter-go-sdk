package api

import (
	"encoding/json"
	"time"
)

type StatusResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      string        `json:"id,omitempty"`
	Result  *StatusResult `json:"result,omitempty"`
	Error   *Error        `json:"error,omitempty"`
}

type StatusResult struct {
	Version           string    `json:"version"`
	LatestBlockHash   string    `json:"latest_block_hash"`
	LatestAppHash     string    `json:"latest_app_hash"`
	LatestBlockHeight string    `json:"latest_block_height"`
	LatestBlockTime   time.Time `json:"latest_block_time"`
	StateHistory      string    `json:"state_history"`
	TmStatus          struct {
		NodeInfo struct {
			ProtocolVersion struct {
				P2P   string `json:"p2p"`
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"protocol_version"`
			ID         string `json:"id,omitempty"`
			ListenAddr string `json:"listen_addr"`
			Network    string `json:"network"`
			Version    string `json:"version"`
			Channels   string `json:"channels"`
			Moniker    string `json:"moniker"`
			Other      struct {
				TxIndex    string `json:"tx_index"`
				RPCAddress string `json:"rpc_address"`
			} `json:"other"`
		} `json:"node_info"`
		SyncInfo struct {
			LatestBlockHash   string    `json:"latest_block_hash"`
			LatestAppHash     string    `json:"latest_app_hash"`
			LatestBlockHeight string    `json:"latest_block_height"`
			LatestBlockTime   time.Time `json:"latest_block_time"`
			CatchingUp        bool      `json:"catching_up"`
		} `json:"sync_info"`
		ValidatorInfo struct {
			Address string `json:"address"`
			PubKey  struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"pub_key"`
			VotingPower string `json:"voting_power"`
		} `json:"validator_info"`
	} `json:"tm_status"`
}

// Returns node status info.
func (a *Api) Status() (*StatusResult, error) {

	res, err := a.client.R().Get("/status")
	if err := hasError(res, err); err != nil {
		return nil, err
	}

	response := new(StatusResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
