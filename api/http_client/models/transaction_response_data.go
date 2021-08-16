package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Data interface {
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

type CreateSwapPoolData struct {
	Coin0   *Coin  `json:"coin0,omitempty"`
	Coin1   *Coin  `json:"coin1,omitempty"`
	Volume0 string `json:"volume0,omitempty"`
	Volume1 string `json:"volume1,omitempty"`
}

type AddLiquidityData struct {
	Coin0          *Coin  `json:"coin0,omitempty"`
	Coin1          *Coin  `json:"coin1,omitempty"`
	Volume0        string `json:"volume0,omitempty"`
	MaximumVolume1 string `json:"maximum_volume1,omitempty"`
}

type RemoveLiquidityData struct {
	Coin0          *Coin  `json:"coin0,omitempty"`
	Coin1          *Coin  `json:"coin1,omitempty"`
	Liquidity      string `json:"liquidity,omitempty"`
	MinimumVolume0 string `json:"minimum_volume0,omitempty"`
	MinimumVolume1 string `json:"minimum_volume1,omitempty"`
}

type SellSwapPoolData struct {
	Coins             []*Coin `json:"coins,omitempty"`
	ValueToSell       string  `json:"value_to_sell,omitempty"`
	MinimumValueToBuy string  `json:"minimum_value_to_buy,omitempty"`
}

type SellAllSwapPoolData struct {
	Coins             []*Coin `json:"coins,omitempty"`
	MinimumValueToBuy string  `json:"minimum_value_to_buy,omitempty"`
}

type BuySwapPoolData struct {
	Coins              []*Coin `json:"coins,omitempty"`
	ValueToBuy         string  `json:"value_to_buy,omitempty"`
	MaximumValueToSell string  `json:"maximum_value_to_sell,omitempty"`
}

type EditCandidateCommission struct {
	PubKey     string `json:"pub_key,omitempty"`
	Commission uint64 `json:"commission,omitempty"`
}

type MoveStakeData struct {
	From  string `json:"from,omitempty"`
	To    string `json:"to,omitempty"`
	Coin  *Coin  `json:"coin,omitempty"`
	Stake string `json:"stake,omitempty"`
}

type MintTokenData struct {
	Coin  *Coin  `json:"coin,omitempty"`
	Value string `json:"value,omitempty"`
}

type BurnTokenData struct {
	Coin  *Coin  `json:"coin,omitempty"`
	Value string `json:"value,omitempty"`
}

type CreateTokenData struct {
	Name          string `json:"name,omitempty"`
	Symbol        string `json:"symbol,omitempty"`
	InitialAmount string `json:"initial_amount,omitempty"`
	MaxSupply     string `json:"max_supply,omitempty"`
	Mintable      bool   `json:"mintable,omitempty"`
	Burnable      bool   `json:"burnable,omitempty"`
}

type RecreateTokenData struct {
	Name          string `json:"name,omitempty"`
	Symbol        string `json:"symbol,omitempty"`
	InitialAmount string `json:"initial_amount,omitempty"`
	MaxSupply     string `json:"max_supply,omitempty"`
	Mintable      bool   `json:"mintable,omitempty"`
	Burnable      bool   `json:"burnable,omitempty"`
}

type VoteCommissionData struct {
	PubKey                  string `json:"pub_key,omitempty"`
	Height                  uint64 `json:"height,string,omitempty"`
	Coin                    *Coin  `json:"coin,omitempty"`
	PayloadByte             string `json:"payload_byte,omitempty"`
	Send                    string `json:"send,omitempty"`
	BuyBancor               string `json:"buy_bancor,omitempty"`
	SellBancor              string `json:"sell_bancor,omitempty"`
	SellAllBancor           string `json:"sell_all_bancor,omitempty"`
	BuyPoolBase             string `json:"buy_pool_base,omitempty"`
	BuyPoolDelta            string `json:"buy_pool_delta,omitempty"`
	SellPoolBase            string `json:"sell_pool_base,omitempty"`
	SellPoolDelta           string `json:"sell_pool_delta,omitempty"`
	SellAllPoolBase         string `json:"sell_all_pool_base,omitempty"`
	SellAllPoolDelta        string `json:"sell_all_pool_delta,omitempty"`
	CreateTicker3           string `json:"create_ticker3,omitempty"`
	CreateTicker4           string `json:"create_ticker4,omitempty"`
	CreateTicker5           string `json:"create_ticker5,omitempty"`
	CreateTicker6           string `json:"create_ticker6,omitempty"`
	CreateTicker7_10        string `json:"create_ticker7_10,omitempty"`
	CreateCoin              string `json:"create_coin,omitempty"`
	CreateToken             string `json:"create_token,omitempty"`
	RecreateCoin            string `json:"recreate_coin,omitempty"`
	RecreateToken           string `json:"recreate_token,omitempty"`
	DeclareCandidacy        string `json:"declare_candidacy,omitempty"`
	Delegate                string `json:"delegate,omitempty"`
	Unbond                  string `json:"unbond,omitempty"`
	RedeemCheck             string `json:"redeem_check,omitempty"`
	SetCandidateOn          string `json:"set_candidate_on,omitempty"`
	SetCandidateOff         string `json:"set_candidate_off,omitempty"`
	CreateMultisig          string `json:"create_multisig,omitempty"`
	MultisendBase           string `json:"multisend_base,omitempty"`
	MultisendDelta          string `json:"multisend_delta,omitempty"`
	EditCandidate           string `json:"edit_candidate,omitempty"`
	SetHaltBlock            string `json:"set_halt_block,omitempty"`
	EditTickerOwner         string `json:"edit_ticker_owner,omitempty"`
	EditMultisig            string `json:"edit_multisig,omitempty"`
	PriceVote               string `json:"price_vote,omitempty"`
	EditCandidatePublicKey  string `json:"edit_candidate_public_key,omitempty"`
	CreateSwapPool          string `json:"create_swap_pool,omitempty"`
	AddLiquidity            string `json:"add_liquidity,omitempty"`
	RemoveLiquidity         string `json:"remove_liquidity,omitempty"`
	EditCandidateCommission string `json:"edit_candidate_commission,omitempty"`
	MoveStake               string `json:"move_stake,omitempty"`
	MintToken               string `json:"mint_token,omitempty"`
	BurnToken               string `json:"burn_token,omitempty"`
	VoteCommission          string `json:"vote_commission,omitempty"`
	VoteUpdate              string `json:"vote_update,omitempty"`
	FailedTx                string `json:"failed_tx,omitempty"`
	AddLimitOrder           string `json:"add_limit_order,omitempty"`
	RemoveLimitOrder        string `json:"remove_limit_order,omitempty"`
}

type VoteUpdateData struct {
	PubKey  string `json:"pub_key,omitempty"`
	Height  uint64 `json:"height,string,omitempty"`
	Version string `json:"version,omitempty"`
}

type AddLimitOrderData struct {
	CoinToSell  *Coin  `json:"coin_to_sell,omitempty"`
	ValueToSell string `json:"value_to_sell,omitempty"`
	CoinToBuy   *Coin  `json:"coin_to_buy,omitempty"`
	ValueToBuy  string `json:"value_to_buy,omitempty"`
}

type RemoveLimitOrderData struct {
	Id uint64 `json:"id,string,omitempty"`
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
	case "type.googleapis.com/api_pb.CreateSwapPoolData":
		return &CreateSwapPoolData{}
	case "type.googleapis.com/api_pb.AddLiquidityData":
		return &AddLiquidityData{}
	case "type.googleapis.com/api_pb.RemoveLiquidityData":
		return &RemoveLiquidityData{}
	case "type.googleapis.com/api_pb.CreateTokenData":
		return &CreateTokenData{}
	case "type.googleapis.com/api_pb.RecreateTokenData":
		return &RecreateTokenData{}
	case "type.googleapis.com/api_pb.MoveStakeData":
		return &MoveStakeData{}
	case "type.googleapis.com/api_pb.EditCandidateCommission":
		return &EditCandidateCommission{}
	case "type.googleapis.com/api_pb.MintTokenData":
		return &MintTokenData{}
	case "type.googleapis.com/api_pb.BurnTokenData":
		return &BurnTokenData{}
	case "type.googleapis.com/api_pb.VoteUpdateData":
		return &VoteUpdateData{}
	case "type.googleapis.com/api_pb.VoteCommissionData":
		return &VoteCommissionData{}
	case "type.googleapis.com/api_pb.BuySwapPoolData":
		return &BuySwapPoolData{}
	case "type.googleapis.com/api_pb.SellSwapPoolData":
		return &SellSwapPoolData{}
	case "type.googleapis.com/api_pb.SellAllSwapPoolData":
		return &SellAllSwapPoolData{}
	case "type.googleapis.com/api_pb.AddLimitOrderData":
		return &AddLimitOrderData{}
	case "type.googleapis.com/api_pb.RemoveLimitOrderData":
		return &RemoveLimitOrderData{}
	default:
		return nil
	}
}

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
