package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Data interface {
	data()
}

type SendData struct {
	Coin  *Coin  `json:"coin"`
	To    string `json:"to"`
	Value string `json:"value"`
}

type SellCoinData struct {
	CoinToSell        *Coin  `json:"coin_to_sell"`
	ValueToSell       string `json:"value_to_sell"`
	CoinToBuy         *Coin  `json:"coin_to_buy"`
	MinimumValueToBuy string `json:"minimum_value_to_buy"`
}

type SellAllCoinData struct {
	CoinToSell        *Coin  `json:"coin_to_sell,omitempty"`
	CoinToBuy         *Coin  `json:"coin_to_buy,omitempty"`
	MinimumValueToBuy string `json:"minimum_value_to_buy,omitempty"`
}

type BuyCoinData struct {
	CoinToBuy          *Coin  `json:"coin_to_buy,omitempty"`
	ValueToBuy         string `json:"value_to_buy,omitempty"`
	CoinToSell         *Coin  `json:"coin_to_sell,omitempty"`
	MaximumValueToSell string `json:"maximum_value_to_sell,omitempty"`
}

type CreateCoinData struct {
	Name                 string `json:"name,omitempty"`
	Symbol               string `json:"symbol,omitempty"`
	InitialAmount        string `json:"initial_amount,omitempty"`
	InitialReserve       string `json:"initial_reserve,omitempty"`
	ConstantReserveRatio uint64 `json:"constant_reserve_ratio,string,omitempty"`
	MaxSupply            string `json:"max_supply,omitempty"`
}

type DeclareCandidacyData struct {
	Address    string `json:"address,omitempty"`
	PubKey     string `json:"pub_key,omitempty"`
	Commission uint64 `json:"commission,string,omitempty"`
	Coin       *Coin  `json:"coin,omitempty"`
	Stake      string `json:"stake,omitempty"`
}

type DelegateData struct {
	PubKey string `json:"pub_key,omitempty"`
	Coin   *Coin  `json:"coin,omitempty"`
	Value  string `json:"value,omitempty"`
}

type UnbondData struct {
	PubKey string `json:"pub_key,omitempty"`
	Coin   *Coin  `json:"coin,omitempty"`
	Value  string `json:"value,omitempty"`
}

type RedeemCheckData struct {
	RawCheck string `json:"raw_check,omitempty"`
	Proof    string `json:"proof,omitempty"`
}

type SetCandidateOnData struct {
	PubKey string `json:"pub_key,omitempty"`
}

type SetCandidateOffData struct {
	PubKey string `json:"pub_key,omitempty"`
}

type CreateMultisigData struct {
	Threshold uint64        `json:"threshold,string,omitempty"`
	Weights   intsAsStrings `json:"weights,omitempty"`
	Addresses []string      `json:"addresses,omitempty"`
}

type MultiSendData struct {
	List []*SendData `json:"list,omitempty"`
}

type EditCandidateData struct {
	PubKey         string `json:"pub_key,omitempty"`
	RewardAddress  string `json:"reward_address,omitempty"`
	OwnerAddress   string `json:"owner_address,omitempty"`
	ControlAddress string `json:"control_address,omitempty"`
}

type SetHaltBlockData struct {
	PubKey string `json:"pub_key,omitempty"`
	Height uint64 `json:"height,string,omitempty"`
}

type RecreateCoinData struct {
	Name                 string `json:"name,omitempty"`
	Symbol               string `json:"symbol,omitempty"`
	InitialAmount        string `json:"initial_amount,omitempty"`
	InitialReserve       string `json:"initial_reserve,omitempty"`
	ConstantReserveRatio uint64 `json:"constant_reserve_ratio,string,omitempty"`
	MaxSupply            string `json:"max_supply,omitempty"`
}

type EditCoinOwnerData struct {
	Symbol   string `json:"symbol,omitempty"`
	NewOwner string `json:"new_owner,omitempty"`
}

type EditMultisigData CreateMultisigData

type PriceVoteData struct {
	Price string `json:"price,omitempty"`
}

type EditCandidatePublicKeyData struct {
	PubKey    string `json:"pub_key,omitempty"`
	NewPubKey string `json:"new_pub_key,omitempty"`
}

func newData(t string) Data {
	switch t {
	case "type.googleapis.com/api_pb.SendData":
		return &SendData{}
	case "type.googleapis.com/api_pb.SellCoinData":
		return &SellCoinData{}
	case "type.googleapis.com/api_pb.SellAllCoinData":
		return &SellAllCoinData{}
	case "type.googleapis.com/api_pb.BuyCoinData":
		return &BuyCoinData{}
	case "type.googleapis.com/api_pb.CreateCoinData":
		return &CreateCoinData{}
	case "type.googleapis.com/api_pb.DeclareCandidacyData":
		return &DeclareCandidacyData{}
	case "type.googleapis.com/api_pb.DelegateData":
		return &DelegateData{}
	case "type.googleapis.com/api_pb.UnbondData":
		return &UnbondData{}
	case "type.googleapis.com/api_pb.RedeemCheckData":
		return &RedeemCheckData{}
	case "type.googleapis.com/api_pb.SetCandidateOnData":
		return &SetCandidateOnData{}
	case "type.googleapis.com/api_pb.SetCandidateOffData":
		return &SetCandidateOffData{}
	case "type.googleapis.com/api_pb.CreateMultisigData":
		return &CreateMultisigData{}
	case "type.googleapis.com/api_pb.MultiSendData":
		return &MultiSendData{}
	case "type.googleapis.com/api_pb.EditCandidateData":
		return &EditCandidateData{}
	case "type.googleapis.com/api_pb.SetHaltBlockData":
		return &SetHaltBlockData{}
	case "type.googleapis.com/api_pb.RecreateCoinData":
		return &RecreateCoinData{}
	case "type.googleapis.com/api_pb.EditCoinOwnerData":
		return &EditCoinOwnerData{}
	case "type.googleapis.com/api_pb.EditMultisigData":
		return &EditMultisigData{}
	case "type.googleapis.com/api_pb.PriceVoteData":
		return &PriceVoteData{}
	case "type.googleapis.com/api_pb.EditCandidatePublicKeyData":
		return &EditCandidatePublicKeyData{}
	default:
		return nil
	}
}

func (e *SendData) data()                   {}
func (e *SellCoinData) data()               {}
func (e *SellAllCoinData) data()            {}
func (e *BuyCoinData) data()                {}
func (e *CreateCoinData) data()             {}
func (e *DeclareCandidacyData) data()       {}
func (e *DelegateData) data()               {}
func (e *UnbondData) data()                 {}
func (e *RedeemCheckData) data()            {}
func (e *SetCandidateOnData) data()         {}
func (e *SetCandidateOffData) data()        {}
func (e *CreateMultisigData) data()         {}
func (e *MultiSendData) data()              {}
func (e *EditCandidateData) data()          {}
func (e *SetHaltBlockData) data()           {}
func (e *RecreateCoinData) data()           {}
func (e *EditCoinOwnerData) data()          {}
func (e *EditMultisigData) data()           {}
func (e *PriceVoteData) data()              {}
func (e *EditCandidatePublicKeyData) data() {}

// convertToData returns Transaction Data model
func convertToData(value *ProtobufAny) (Data, error) {
	var v map[string]interface{} = *value
	t := v["@type"].(string)
	data := newData(t)
	if data == nil {
		return nil, fmt.Errorf("data type unknown: %s", t)
	}

	err := value.UnmarshalTo(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type intsAsStrings []uint64

func (d *intsAsStrings) MarshalJSON() ([]byte, error) {
	var weights []string
	for _, weight := range *d {
		weights = append(weights, strconv.Itoa(int(weight)))
	}
	return json.Marshal(weights)
}

func (d *intsAsStrings) UnmarshalJSON(data []byte) error {
	var adx []string
	if err := json.Unmarshal(data, &adx); err != nil {
		return err
	}

	weights := make([]uint64, 0, len(adx))
	for _, strWeight := range adx {
		weight, _ := strconv.Atoi(strWeight)
		weights = append(weights, uint64(weight))
	}
	*d = weights
	return nil
}
