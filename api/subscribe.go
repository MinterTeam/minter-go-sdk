package api

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"strconv"
	"strings"
	"time"
)

const (
	// EventTypeKey is a reserved composite key for event name.
	eventTypeKey = "tm.event"
	// TxHashKey is a reserved key, used to specify transaction's hash.
	// see EventBus#PublishEventTx
	txHashKey = "tx.hash"
	// TxHeightKey is a reserved key, used to specify transaction block's height.
	// see EventBus#PublishEventTx
	txHeightKey = "tx.height"

	// BlockHeightKey is a reserved key used for indexing BeginBlock and Endblock
	// events.
	blockHeightKey = "block.height"
)

type SubscribeEventType string

// Reserved event types.
// Block level events for mass consumption by users.
// These events are triggered from the state package,
// after a block has been committed.
// These are also used by the tx indexer for async indexing.
const (
	EventNewBlock            SubscribeEventType = "NewBlock"
	EventNewBlockHeader      SubscribeEventType = "NewBlockHeader"
	EventNewEvidence         SubscribeEventType = "NewEvidence"
	EventTx                  SubscribeEventType = "Tx"
	EventValidatorSetUpdates SubscribeEventType = "ValidatorSetUpdates"
)

func QueryFrom(address string) string {
	if strings.HasPrefix(address, "Mx") {
		address = address[2:]
	}
	return QueryTag("from", address)
}

func QueryTo(address string) string {
	if strings.HasPrefix(address, "Mx") {
		address = address[2:]
	}
	return QueryTag("to", address)
}

func QueryType(t byte) string {
	return QueryTag("type", hex.EncodeToString([]byte{t}))
}

func QueryTag(tag, param string) string {
	return fmt.Sprintf("tags.tx.%s = '%s'", tag, param)
}

func QueryHash(hash string) string {
	if strings.HasPrefix(hash, "Mt") {
		hash = hash[2:]
	}
	return fmt.Sprintf("%s = '%s'", txHashKey, strings.ToUpper(hash))
}

func QueryFail() string {
	return QueryTag("fail", "1")
}

func QueryCommissionCoin(coinID uint64) string {
	return QueryTag("commission_coin", strconv.Itoa(int(coinID)))
}
func QueryCoinID(coinID uint64) string {
	return QueryTag("coin_id", strconv.Itoa(int(coinID)))
}
func QueryCoinSymbol(symbol string) string {
	return QueryTag("coin_symbol", symbol)
}
func QueryPublicKey(publicKey string) string {
	return QueryTag("public_key", publicKey)
}
func QueryCommissionConversion(isPool bool) string {
	conv := "bancor"
	if isPool {
		conv = "pool"
	}
	return QueryTag("commission_conversion", conv)
}

func QueryEvent(event SubscribeEventType) string {
	return fmt.Sprintf("%s = '%s'", eventTypeKey, event)
}

func QueryTxHeight(height uint64) string {
	return fmt.Sprintf("%s = '%d'", txHeightKey, height)
}

func QueryBlockHeight(height uint64) string {
	return fmt.Sprintf("%s = '%d'", blockHeightKey, height)
}

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
				Height        uint64 `json:"height"`
				LastBlockID   struct {
					Hash  string `json:"hash"`
					Parts struct {
						Hash  string `json:"hash"`
						Total uint64 `json:"total"`
					} `json:"parts"`
				} `json:"last_block_id"`
				LastCommitHash     string    `json:"last_commit_hash"`
				LastResultsHash    string    `json:"last_results_hash"`
				NextValidatorsHash string    `json:"next_validators_hash"`
				ProposerAddress    string    `json:"proposer_address"`
				Time               time.Time `json:"time"`
				ValidatorsHash     string    `json:"validators_hash"`
				Version            struct {
					Block uint64 `json:"block"`
				} `json:"version"`
			} `json:"header"`
			LastCommit struct {
				BlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Hash  string `json:"hash"`
						Total uint64 `json:"total"`
					} `json:"parts"`
				} `json:"block_id"`
				Height     uint64 `json:"height"`
				Round      uint64 `json:"round"`
				Signatures []struct {
					BlockIDFlag      uint64    `json:"block_id_flag"`
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
					MaxBytes uint64 `json:"max_bytes"`
					MaxGas   uint64 `json:"max_gas"`
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
		Height uint64 `json:"height"`
		Result struct {
			Events []struct {
				Attributes []struct {
					Index bool   `json:"index"`
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
				Type string `json:"type"`
			} `json:"events"`
			GasUsed   uint64 `json:"gas_used"`
			GasWanted uint64 `json:"gas_wanted"`
		} `json:"result"`
		Tx string `json:"tx"`
	} `json:"data"`
	Events []struct {
		Key    string   `json:"key"`
		Events []string `json:"events"`
	} `json:"events"`
}

func SubscribeNewTxFindTag(message, tag string) (string, error) {
	recv, err := SubscribeNewTxParse(message)
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

func SubscribeNewTxParse(message string) (*SubscribeNewTxResult, error) {
	var recv SubscribeNewTxResult
	err := json.Unmarshal([]byte(message), &recv)
	if err != nil {
		return nil, err
	}

	return &recv, nil
}

func SubscribeNewTxToTx(message string) (transaction.Signed, error) {
	recv, err := SubscribeNewTxParse(message)
	if err != nil {
		return nil, err
	}

	decodeString, err := base64.StdEncoding.DecodeString(recv.Data.Tx)
	if err != nil {
		return nil, err
	}

	signed, err := transaction.Decode("0x" + hex.EncodeToString(decodeString))
	if err != nil {
		return nil, err
	}

	return signed, nil
}
func SubscribeNewBlockParse(message string) (*SubscribeNewBlockResult, error) {
	var recv SubscribeNewBlockResult
	err := json.Unmarshal([]byte(message), &recv)
	if err != nil {
		return nil, err
	}

	return &recv, nil
}
