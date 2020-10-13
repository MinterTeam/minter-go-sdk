package api

import (
	"encoding/json"
	"errors"
)

type TransactionResponse struct {
	Jsonrpc string             `json:"jsonrpc"`
	ID      string             `json:"id,omitempty"`
	Result  *TransactionResult `json:"result,omitempty"`
	Error   *Error             `json:"error,omitempty"`
}

type Coin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
}

type TransactionResult struct {
	Hash        string          `json:"hash"`
	RawTx       string          `json:"raw_tx"`
	Height      string          `json:"height"`
	Index       int             `json:"index,omitempty"`
	From        string          `json:"from"`
	Nonce       string          `json:"nonce"`
	Gas         string          `json:"gas"`
	GasPrice    int             `json:"gas_price"`
	GasCoin     Coin            `json:"gas_coin"`
	Type        int             `json:"type"`
	Data        transactionData `json:"data"`
	Payload     []byte          `json:"payload"`
	ServiceData []byte          `json:"service_data"`
	Tags        struct {
		TxCoinToBuy       string `json:"tx.coin_to_buy,omitempty"`
		TxCoinToSell      string `json:"tx.coin_to_sell,omitempty"`
		TxReturn          string `json:"tx.return,omitempty"`
		TxType            string `json:"tx.type,omitempty"`
		TxFrom            string `json:"tx.from,omitempty"`
		TxTo              string `json:"tx.to,omitempty"`
		TxCoinID          string `json:"tx.coin_id,omitempty"`
		TxCoinSymbol      string `json:"tx.coin_symbol,omitempty"`
		TxSellAmount      string `json:"tx.sell_amount,omitempty"`
		TxCreatedMultisig string `json:"tx.created_multisig,omitempty"`
	} `json:"tags,omitempty"`
	Code uint32 `json:"code,omitempty"`
	Log  string `json:"log,omitempty"`
}

type transactionData map[string]interface{}

func (dt *transactionData) FillStruct(data tdi) error {
	b, err := json.Marshal(dt)
	if err != nil {
		return err
	}

	return data.fill(b)
}

func (t *TransactionResult) IsValid() bool {
	return t.Code == 0
}

func (t *TransactionResult) ErrorLog() error {
	if t.IsValid() {
		return nil
	}
	return errors.New(t.Log)
}

// Converting transaction map data to the structure interface regarding transaction type
func (t *TransactionResult) DataStruct() (tdi, error) {

	var data tdi
	switch t.Type {
	case 1:
		data = &SendData{}
	case 2:
		data = &SellCoinData{}
	case 3:
		data = &SellAllCoinData{}
	case 4:
		data = &BuyCoinData{}
	case 5:
		data = &CreateCoinData{}
	case 6:
		data = &DeclareCandidacyData{}
	case 7:
		data = &DelegateData{}
	case 8:
		data = &UnbondData{}
	case 9:
		data = &RedeemCheckData{}
	case 10:
		data = &SetCandidateOnData{}
	case 11:
		data = &SetCandidateOffData{}
	case 12:
		data = &CreateMultisigData{}
	case 13:
		data = &MultisendData{}
	case 14:
		data = &EditCandidateData{}
	default:
		return nil, errors.New("unknown transaction type")
	}

	return data, t.Data.FillStruct(data)
}

type tdi interface {
	fill([]byte) error
}

type SendData struct {
	Coin  Coin   `json:"coin"`
	To    string `json:"to"`
	Value string `json:"value"`
}

func (s *SendData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type SellCoinData struct {
	CoinToSell        Coin   `json:"coin_to_sell"`
	ValueToSell       string `json:"value_to_sell"`
	CoinToBuy         Coin   `json:"coin_to_buy"`
	MinimumValueToBuy string `json:"minimum_value_to_buy"`
}

func (s *SellCoinData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type SellAllCoinData struct {
	CoinToSell        Coin   `json:"coin_to_sell"`
	CoinToBuy         Coin   `json:"coin_to_buy"`
	MinimumValueToBuy string `json:"minimum_value_to_buy"`
}

func (s *SellAllCoinData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type BuyCoinData struct {
	CoinToBuy          Coin   `json:"coin_to_buy"`
	ValueToBuy         string `json:"value_to_buy"`
	CoinToSell         Coin   `json:"coin_to_sell"`
	MaximumValueToSell string `json:"maximum_value_to_sell"`
}

func (s *BuyCoinData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type CreateCoinData struct {
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	InitialAmount        string `json:"initial_amount"`
	InitialReserve       string `json:"initial_reserve"`
	ConstantReserveRatio string `json:"constant_reserve_ratio"`
	MaxSupply            string `json:"max_supply"`
}

func (s *CreateCoinData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type DeclareCandidacyData struct {
	Address    string `json:"address"`
	PubKey     string `json:"pub_key"`
	Commission string `json:"commission"`
	Coin       Coin   `json:"coin"`
	Stake      string `json:"stake"`
}

func (s *DeclareCandidacyData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type DelegateData struct {
	PubKey string `json:"pub_key"`
	Coin   Coin   `json:"coin"`
	Value  string `json:"value"`
}

func (s *DelegateData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type UnbondData struct {
	PubKey string `json:"pub_key"`
	Coin   Coin   `json:"coin"`
	Value  string `json:"value"`
}

// todo add more types

func (s *UnbondData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type RedeemCheckData struct {
	RawCheck string `json:"raw_check"`
	Proof    string `json:"proof"`
}

func (s *RedeemCheckData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type SetCandidateOnData struct {
	PubKey string `json:"pub_key"`
}

func (s *SetCandidateOnData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type SetCandidateOffData struct {
	PubKey string `json:"pub_key"`
}

func (s *SetCandidateOffData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type EditCandidateData struct {
	PubKey        string `json:"pub_key"`
	RewardAddress string `json:"reward_address"`
	OwnerAddress  string `json:"owner_address"`
}

func (s *EditCandidateData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type CreateMultisigData struct {
	Threshold uint       `json:"threshold"`
	Weights   []uint     `json:"weights"`
	Addresses [][20]byte `json:"addresses"`
}

func (s *CreateMultisigData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type MultisendData struct {
	List []MultisendDataItem
}

func (s *MultisendData) fill(b []byte) error {
	return json.Unmarshal(b, s)
}

type MultisendDataItem SendData

// Returns transaction info.
func (a *Api) Transaction(hash string) (*TransactionResult, error) {

	params := make(map[string]string)
	params["hash"] = hash

	res, err := a.client.R().SetQueryParams(params).Get("/transaction")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, NewResponseError(res)
	}

	response := new(TransactionResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
