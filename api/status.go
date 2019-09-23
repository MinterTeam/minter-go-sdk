package api

import (
	"encoding/json"
	"time"
)

type StatusResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
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
				ID         string `json:"id"`
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
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

// Returns node status info.
func (a *Api) Status() (*StatusResponse, error) {

	res, err := a.client.R().Get("/status")
	if err != nil {
		return nil, err
	}

	result := new(StatusResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
