package api

import (
	"encoding/json"
	"time"
)

type SubscribeNewBlockResult struct {
	Query string `json:"query"`
	Data  struct {
		Block struct {
			Data struct {
				Txs []interface{} `json:"txs"`
			} `json:"data"`
			Evidence struct {
				Evidence []interface{} `json:"evidence"`
			} `json:"evidence"`
			Header struct {
				AppHash       string `json:"app_hash"`
				ChainID       string `json:"chain_id"`
				ConsensusHash string `json:"consensus_hash"`
				DataHash      string `json:"data_hash"`
				EvidenceHash  string `json:"evidence_hash"`
				Height        int    `json:"height"`
				LastBlockID   struct {
					Hash  string `json:"hash"`
					Parts struct {
						Hash  string `json:"hash"`
						Total int    `json:"total"`
					} `json:"parts"`
				} `json:"last_block_id"`
				LastCommitHash     string    `json:"last_commit_hash"`
				LastResultsHash    string    `json:"last_results_hash"`
				NextValidatorsHash string    `json:"next_validators_hash"`
				ProposerAddress    string    `json:"proposer_address"`
				Time               time.Time `json:"time"`
				ValidatorsHash     string    `json:"validators_hash"`
				Version            struct {
					Block int `json:"block"`
				} `json:"version"`
			} `json:"header"`
			LastCommit struct {
				BlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Hash  string `json:"hash"`
						Total int    `json:"total"`
					} `json:"parts"`
				} `json:"block_id"`
				Height     int `json:"height"`
				Round      int `json:"round"`
				Signatures []struct {
					BlockIDFlag      int       `json:"block_id_flag"`
					Signature        string    `json:"signature"`
					Timestamp        time.Time `json:"timestamp"`
					ValidatorAddress string    `json:"validator_address"`
				} `json:"signatures"`
			} `json:"last_commit"`
		} `json:"block"`
		ResultBeginBlock struct {
		} `json:"result_begin_block"`
		ResultEndBlock struct {
			ConsensusParamUpdates struct {
				Block struct {
					MaxBytes int `json:"max_bytes"`
					MaxGas   int `json:"max_gas"`
				} `json:"block"`
			} `json:"consensus_param_updates"`
			ValidatorUpdates interface{} `json:"validator_updates"`
		} `json:"result_end_block"`
	} `json:"data"`
	Events []struct {
		Key    string   `json:"key"`
		Events []string `json:"events"`
	} `json:"events"`
}

type SubscribeNewTxResult struct {
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
	var recv SubscribeNewTxResult
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
