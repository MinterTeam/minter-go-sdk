package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Data interface {
	clone() Data
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

var data = map[string]Data{
	"type.googleapis.com/api_pb.SendData":                   &SendData{},
	"type.googleapis.com/api_pb.SellCoinData":               &SellCoinData{},
	"type.googleapis.com/api_pb.SellAllCoinData":            &SellAllCoinData{},
	"type.googleapis.com/api_pb.BuyCoinData":                &BuyCoinData{},
	"type.googleapis.com/api_pb.CreateCoinData":             &CreateCoinData{},
	"type.googleapis.com/api_pb.DeclareCandidacyData":       &DeclareCandidacyData{},
	"type.googleapis.com/api_pb.DelegateData":               &DelegateData{},
	"type.googleapis.com/api_pb.UnbondData":                 &UnbondData{},
	"type.googleapis.com/api_pb.RedeemCheckData":            &RedeemCheckData{},
	"type.googleapis.com/api_pb.SetCandidateOnData":         &SetCandidateOnData{},
	"type.googleapis.com/api_pb.SetCandidateOffData":        &SetCandidateOffData{},
	"type.googleapis.com/api_pb.CreateMultisigData":         &CreateMultisigData{},
	"type.googleapis.com/api_pb.MultiSendData":              &MultiSendData{},
	"type.googleapis.com/api_pb.EditCandidateData":          &EditCandidateData{},
	"type.googleapis.com/api_pb.SetHaltBlockData":           &SetHaltBlockData{},
	"type.googleapis.com/api_pb.RecreateCoinData":           &RecreateCoinData{},
	"type.googleapis.com/api_pb.EditCoinOwnerData":          &EditCoinOwnerData{},
	"type.googleapis.com/api_pb.EditMultisigData":           &EditMultisigData{},
	"type.googleapis.com/api_pb.PriceVoteData":              &PriceVoteData{},
	"type.googleapis.com/api_pb.EditCandidatePublicKeyData": &EditCandidatePublicKeyData{},
}

func (e *SendData) clone() Data {
	c := *e
	return &c
}
func (e *SellCoinData) clone() Data {
	c := *e
	return &c
}
func (e *SellAllCoinData) clone() Data {
	c := *e
	return &c
}
func (e *BuyCoinData) clone() Data {
	c := *e
	return &c
}
func (e *CreateCoinData) clone() Data {
	c := *e
	return &c
}
func (e *DeclareCandidacyData) clone() Data {
	c := *e
	return &c
}
func (e *DelegateData) clone() Data {
	c := *e
	return &c
}
func (e *UnbondData) clone() Data {
	c := *e
	return &c
}
func (e *RedeemCheckData) clone() Data {
	c := *e
	return &c
}
func (e *SetCandidateOnData) clone() Data {
	c := *e
	return &c
}
func (e *SetCandidateOffData) clone() Data {
	c := *e
	return &c
}
func (e *CreateMultisigData) clone() Data {
	c := *e
	return &c
}
func (e *MultiSendData) clone() Data {
	c := *e
	return &c
}
func (e *EditCandidateData) clone() Data {
	c := *e
	return &c
}
func (e *SetHaltBlockData) clone() Data {
	c := *e
	return &c
}
func (e *RecreateCoinData) clone() Data {
	c := *e
	return &c
}
func (e *EditCoinOwnerData) clone() Data {
	c := *e
	return &c
}
func (e *EditMultisigData) clone() Data {
	c := *e
	return &c
}
func (e *PriceVoteData) clone() Data {
	c := *e
	return &c
}
func (e *EditCandidatePublicKeyData) clone() Data {
	c := *e
	return &c
}

// convertToData returns Transaction Data model
func convertToData(value *ProtobufAny) (Data, error) {
	var v map[string]interface{} = *value
	t := v["@type"].(string)
	data, ok := data[t]
	if !ok {
		return nil, fmt.Errorf("data type unknown: %s", t)
	}

	clone := data.clone()
	err := value.UnmarshalTo(clone)
	if err != nil {
		return nil, err
	}

	return clone, nil
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
