package api

import (
	"encoding/json"
	"strconv"
	"time"
)

type BlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Hash         string    `json:"hash"`
		Height       string    `json:"height"`
		Time         time.Time `json:"time"`
		NumTxs       string    `json:"num_txs"`
		TotalTxs     string    `json:"total_txs"`
		Transactions []struct {
			Hash       string `json:"hash"`
			RawTx      string `json:"raw_tx"`
			From       string `json:"from"`
			Nonce      string `json:"nonce"`
			GasPrice   int    `json:"gas_price"`
			Type       int    `json:"type"`
			DataType13 struct {
				List []struct {
					Coin  string `json:"coin"`
					To    string `json:"to"`
					Value string `json:"value"`
				} `json:"list"`
			} `json:"data,omitempty"`
			Payload     string `json:"payload"`
			ServiceData string `json:"service_data"`
			Gas         string `json:"gas"`
			GasCoin     string `json:"gas_coin"`
			Tags        struct {
				TxFrom string `json:"tx.from"`
				TxTo   string `json:"tx.to"`
				TxType string `json:"tx.type"`
			} `json:"tags,omitempty"`
			DataType6 struct {
				Address    string `json:"address"`
				PubKey     string `json:"pub_key"`
				Commission string `json:"commission"`
				Coin       string `json:"coin"`
				Stake      string `json:"stake"`
			} `json:"data,omitempty"`
			DataType2 struct {
				CoinToSell        string `json:"coin_to_sell"`
				ValueToSell       string `json:"value_to_sell"`
				CoinToBuy         string `json:"coin_to_buy"`
				MinimumValueToBuy string `json:"minimum_value_to_buy"`
			} `json:"data,omitempty"`
		} `json:"transactions"`
		BlockReward string `json:"block_reward"`
		Size        string `json:"size"`
		Proposer    string `json:"proposer"`
		Validators  []struct {
			PubKey string `json:"pub_key"`
			Signed bool   `json:"signed"`
		} `json:"validators"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

func (a *Api) Block(height int) (*BlockResponse, error) {

	params := make(map[string]string)
	params["height"] = strconv.Itoa(height)

	res, err := a.client.R().SetQueryParams(params).Get("/block")
	if err != nil {
		return nil, err
	}

	result := new(BlockResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
