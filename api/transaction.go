package api

import (
	"encoding/json"
	"errors"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

type TransactionResponse struct {
	Jsonrpc string             `json:"jsonrpc"`
	ID      string             `json:"id,omitempty"`
	Result  *TransactionResult `json:"result,omitempty"`
	Error   *Error             `json:"error,omitempty"`
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
	GasCoin     string          `json:"gas_coin"`
	Type        int             `json:"type"`
	Data        dataTransaction `json:"data"`
	Payload     []byte          `json:"payload"`
	ServiceData []byte          `json:"service_data"`
	Tags        struct {
		TxCoinToBuy  string `json:"tx.coin_to_buy,omitempty"`
		TxCoinToSell string `json:"tx.coin_to_sell,omitempty"`
		TxReturn     string `json:"tx.return,omitempty"`
		TxType       string `json:"tx.type,omitempty"`
		TxFrom       string `json:"tx.from,omitempty"`
		TxTo         string `json:"tx.to,omitempty"`
		TxCoin       string `json:"tx.coin,omitempty"`
		TxSellAmount string `json:"tx.sell_amount,omitempty"`
	} `json:"tags,omitempty"`
	Code uint32 `json:"code,omitempty"`
	Log  string `json:"log,omitempty"`
}

type dataTransaction map[string]interface{}

func (dt *dataTransaction) FillStruct(data interface{}) error {
	bytes, err := json.Marshal(dt)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	return nil
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
func (t *TransactionResult) DataStruct() (interface{}, error) {

	var data interface{}
	switch transaction.Type(t.Type) {
	case transaction.TypeSend:
		data = &SendData{}
	case transaction.TypeSellCoin:
		data = &SellCoinData{}
	case transaction.TypeSellAllCoin:
		data = &SellAllCoinData{}
	case transaction.TypeBuyCoin:
		data = &SellCoinData{}
	case transaction.TypeCreateCoin:
		data = &BuyCoinData{}
	case transaction.TypeDeclareCandidacy:
		data = &DeclareCandidacyData{}
	case transaction.TypeDelegate:
		data = &DelegateData{}
	case transaction.TypeUnbond:
		data = &UnbondData{}
	case transaction.TypeRedeemCheck:
		data = &RedeemCheckData{}
	case transaction.TypeSetCandidateOnline:
		data = &SetCandidateOnData{}
	case transaction.TypeSetCandidateOffline:
		data = &SetCandidateOffData{}
	case transaction.TypeCreateMultisig:
		data = &CreateMultisigData{}
	case transaction.TypeMultisend:
		data = &MultisendData{}
	case transaction.TypeEditCandidate:
		data = &EditCandidateData{}
	default:
		return nil, errors.New("unknown transaction type")
	}

	return data, t.Data.FillStruct(data)
}

type SendData struct {
	Coin  string `json:"coin"`
	To    string `json:"to"`
	Value string `json:"value"`
}

type SellCoinData struct {
	CoinToSell        string `json:"coin_to_sell"`
	ValueToSell       string `json:"value_to_sell"`
	CoinToBuy         string `json:"coin_to_buy"`
	MinimumValueToBuy string `json:"minimum_value_to_buy"`
}

type SellAllCoinData struct {
	CoinToSell        string `json:"coin_to_sell"`
	CoinToBuy         string `json:"coin_to_buy"`
	MinimumValueToBuy string `json:"minimum_value_to_buy"`
}

type BuyCoinData struct {
	CoinToBuy          string `json:"coin_to_buy"`
	ValueToBuy         string `json:"value_to_buy"`
	CoinToSell         string `json:"coin_to_sell"`
	MaximumValueToSell string `json:"maximum_value_to_sell"`
}

type CreateCoinData struct {
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	InitialAmount        string `json:"initial_amount"`
	InitialReserve       string `json:"initial_reserve"`
	ConstantReserveRatio string `json:"constant_reserve_ratio"`
}

type DeclareCandidacyData struct {
	Address    string `json:"address"`
	PubKey     string `json:"pub_key"`
	Commission string `json:"commission"`
	Coin       string `json:"coin"`
	Stake      string `json:"stake"`
}

type DelegateData struct {
	PubKey string `json:"pub_key"`
	Coin   string `json:"coin"`
	Value  string `json:"value"`
}

type UnbondData struct {
	PubKey string `json:"pub_key"`
	Coin   string `json:"coin"`
	Value  string `json:"value"`
}

type RedeemCheckData struct {
	RawCheck string `json:"raw_check"`
	Proof    string `json:"proof"`
}

type SetCandidateOnData struct {
	PubKey string `json:"pub_key"`
}

type SetCandidateOffData struct {
	PubKey string `json:"pub_key"`
}

type EditCandidateData struct {
	PubKey        string `json:"pub_key"`
	RewardAddress string `json:"reward_address"`
	OwnerAddress  string `json:"owner_address"`
}

type CreateMultisigData struct {
	Threshold uint       `json:"threshold"`
	Weights   []uint     `json:"weights"`
	Addresses [][20]byte `json:"addresses"`
}

type MultisendData struct {
	List []MultisendDataItem
}

type MultisendDataItem SendData

// Returns transaction info.
func (a *Api) Transaction(hash string) (*TransactionResult, error) {

	params := make(map[string]string)
	params["hash"] = hash

	res, err := a.client.R().SetQueryParams(params).Get("/transaction")
	if err := hasError(res, err); err != nil {
		return nil, err
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
