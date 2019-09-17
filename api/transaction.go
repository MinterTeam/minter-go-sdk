package api

import "encoding/json"

type TransactionResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Hash     string `json:"hash"`
		RawTx    string `json:"raw_tx"`
		Height   string `json:"height"`
		Index    int    `json:"index"`
		From     string `json:"from"`
		Nonce    string `json:"nonce"`
		Gas      string `json:"gas"`
		GasPrice int    `json:"gas_price"`
		GasCoin  string `json:"gas_coin"`
		Type     int    `json:"type"`
		Data     struct {
			List []struct {
				Coin  string `json:"coin"`
				To    string `json:"to"`
				Value string `json:"value"`
			} `json:"list"`
		} `json:"data"`
		Payload string `json:"payload"`
		Tags    struct {
			TxType string `json:"tx.type"`
			TxFrom string `json:"tx.from"`
			TxTo   string `json:"tx.to"`
		} `json:"tags"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

func (a *Api) Transaction(hash string) (*TransactionResponse, error) {

	params := make(map[string]string)
	params["hash"] = hash

	res, err := a.client.R().SetQueryParams(params).Get("/transaction")
	if err != nil {
		return nil, err
	}

	result := new(TransactionResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
