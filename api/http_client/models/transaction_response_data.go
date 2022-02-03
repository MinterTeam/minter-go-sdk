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
	CoinToSell        *Coin  `json:"coin_to_sell"`
	CoinToBuy         *Coin  `json:"coin_to_buy"`
	MinimumValueToBuy string `json:"minimum_value_to_buy"`
}

type BuyCoinData struct {
	CoinToBuy          *Coin  `json:"coin_to_buy"`
	ValueToBuy         string `json:"value_to_buy"`
	CoinToSell         *Coin  `json:"coin_to_sell"`
	MaximumValueToSell string `json:"maximum_value_to_sell"`
}

type CreateCoinData struct {
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	InitialAmount        string `json:"initial_amount"`
	InitialReserve       string `json:"initial_reserve"`
	ConstantReserveRatio uint64 `json:"constant_reserve_ratio,string"`
	MaxSupply            string `json:"max_supply"`
}

type DeclareCandidacyData struct {
	Address    string `json:"address"`
	PubKey     string `json:"pub_key"`
	Commission uint64 `json:"commission,string"`
	Coin       *Coin  `json:"coin"`
	Stake      string `json:"stake"`
}

type DelegateData struct {
	PubKey string `json:"pub_key"`
	Coin   *Coin  `json:"coin"`
	Value  string `json:"value"`
}

type UnbondData struct {
	PubKey string `json:"pub_key"`
	Coin   *Coin  `json:"coin"`
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

type CreateMultisigData struct {
	Threshold uint64        `json:"threshold,string"`
	Weights   intsAsStrings `json:"weights"`
	Addresses []string      `json:"addresses"`
}

type MultiSendData struct {
	List []*SendData `json:"list"`
}

type EditCandidateData struct {
	PubKey         string `json:"pub_key"`
	RewardAddress  string `json:"reward_address"`
	OwnerAddress   string `json:"owner_address"`
	ControlAddress string `json:"control_address"`
}

type SetHaltBlockData struct {
	PubKey string `json:"pub_key"`
	Height uint64 `json:"height,string"`
}

type RecreateCoinData struct {
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	InitialAmount        string `json:"initial_amount"`
	InitialReserve       string `json:"initial_reserve"`
	ConstantReserveRatio uint64 `json:"constant_reserve_ratio,string"`
	MaxSupply            string `json:"max_supply"`
}

type EditCoinOwnerData struct {
	Symbol   string `json:"symbol"`
	NewOwner string `json:"new_owner"`
}

type EditMultisigData CreateMultisigData

type PriceVoteData struct {
	Price string `json:"price"`
}

type EditCandidatePublicKeyData struct {
	PubKey    string `json:"pub_key"`
	NewPubKey string `json:"new_pub_key"`
}

type CreateSwapPoolData struct {
	Coin0   *Coin  `json:"coin0"`
	Coin1   *Coin  `json:"coin1"`
	Volume0 string `json:"volume0"`
	Volume1 string `json:"volume1"`
}

type AddLiquidityData struct {
	Coin0          *Coin  `json:"coin0"`
	Coin1          *Coin  `json:"coin1"`
	Volume0        string `json:"volume0"`
	MaximumVolume1 string `json:"maximum_volume1"`
}

type RemoveLiquidityData struct {
	Coin0          *Coin  `json:"coin0"`
	Coin1          *Coin  `json:"coin1"`
	Liquidity      string `json:"liquidity"`
	MinimumVolume0 string `json:"minimum_volume0"`
	MinimumVolume1 string `json:"minimum_volume1"`
}

type SellSwapPoolData struct {
	Coins             []*Coin `json:"coins"`
	ValueToSell       string  `json:"value_to_sell"`
	MinimumValueToBuy string  `json:"minimum_value_to_buy"`
}

type SellAllSwapPoolData struct {
	Coins             []*Coin `json:"coins"`
	MinimumValueToBuy string  `json:"minimum_value_to_buy"`
}

type BuySwapPoolData struct {
	Coins              []*Coin `json:"coins"`
	ValueToBuy         string  `json:"value_to_buy"`
	MaximumValueToSell string  `json:"maximum_value_to_sell"`
}

type EditCandidateCommission struct {
	PubKey     string `json:"pub_key"`
	Commission uint64 `json:"commission"`
}

type MoveStakeData struct {
	FromPubKey string `json:"from_pub_key"`
	ToPubKey   string `json:"to_pub_key"`
	Coin       *Coin  `json:"coin"`
	Value      string `json:"value"`
}

type MintTokenData struct {
	Coin  *Coin  `json:"coin"`
	Value string `json:"value"`
}

type BurnTokenData struct {
	Coin  *Coin  `json:"coin"`
	Value string `json:"value"`
}

type CreateTokenData struct {
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	InitialAmount string `json:"initial_amount"`
	MaxSupply     string `json:"max_supply"`
	Mintable      bool   `json:"mintable"`
	Burnable      bool   `json:"burnable"`
}

type RecreateTokenData struct {
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	InitialAmount string `json:"initial_amount"`
	MaxSupply     string `json:"max_supply"`
	Mintable      bool   `json:"mintable"`
	Burnable      bool   `json:"burnable"`
}

type VoteCommissionData struct {
	PubKey                  string `json:"pub_key"`
	Height                  uint64 `json:"height,string"`
	Coin                    *Coin  `json:"coin"`
	PayloadByte             string `json:"payload_byte"`
	Send                    string `json:"send"`
	BuyBancor               string `json:"buy_bancor"`
	SellBancor              string `json:"sell_bancor"`
	SellAllBancor           string `json:"sell_all_bancor"`
	BuyPoolBase             string `json:"buy_pool_base"`
	BuyPoolDelta            string `json:"buy_pool_delta"`
	SellPoolBase            string `json:"sell_pool_base"`
	SellPoolDelta           string `json:"sell_pool_delta"`
	SellAllPoolBase         string `json:"sell_all_pool_base"`
	SellAllPoolDelta        string `json:"sell_all_pool_delta"`
	CreateTicker3           string `json:"create_ticker3"`
	CreateTicker4           string `json:"create_ticker4"`
	CreateTicker5           string `json:"create_ticker5"`
	CreateTicker6           string `json:"create_ticker6"`
	CreateTicker7_10        string `json:"create_ticker7_10"`
	CreateCoin              string `json:"create_coin"`
	CreateToken             string `json:"create_token"`
	RecreateCoin            string `json:"recreate_coin"`
	RecreateToken           string `json:"recreate_token"`
	DeclareCandidacy        string `json:"declare_candidacy"`
	Delegate                string `json:"delegate"`
	Unbond                  string `json:"unbond"`
	RedeemCheck             string `json:"redeem_check"`
	SetCandidateOn          string `json:"set_candidate_on"`
	SetCandidateOff         string `json:"set_candidate_off"`
	CreateMultisig          string `json:"create_multisig"`
	MultisendBase           string `json:"multisend_base"`
	MultisendDelta          string `json:"multisend_delta"`
	EditCandidate           string `json:"edit_candidate"`
	SetHaltBlock            string `json:"set_halt_block"`
	EditTickerOwner         string `json:"edit_ticker_owner"`
	EditMultisig            string `json:"edit_multisig"`
	PriceVote               string `json:"price_vote"`
	EditCandidatePublicKey  string `json:"edit_candidate_public_key"`
	CreateSwapPool          string `json:"create_swap_pool"`
	AddLiquidity            string `json:"add_liquidity"`
	RemoveLiquidity         string `json:"remove_liquidity"`
	EditCandidateCommission string `json:"edit_candidate_commission"`
	MoveStake               string `json:"move_stake"`
	MintToken               string `json:"mint_token"`
	BurnToken               string `json:"burn_token"`
	VoteCommission          string `json:"vote_commission"`
	VoteUpdate              string `json:"vote_update"`
	FailedTx                string `json:"failed_tx"`
	AddLimitOrder           string `json:"add_limit_order"`
	RemoveLimitOrder        string `json:"remove_limit_order"`
}

type VoteUpdateData struct {
	PubKey  string `json:"pub_key"`
	Height  uint64 `json:"height,string"`
	Version string `json:"version"`
}

type AddLimitOrderData struct {
	CoinToSell  *Coin  `json:"coin_to_sell"`
	ValueToSell string `json:"value_to_sell"`
	CoinToBuy   *Coin  `json:"coin_to_buy"`
	ValueToBuy  string `json:"value_to_buy"`
}

type RemoveLimitOrderData struct {
	Id uint64 `json:"id,string"`
}

type LockData struct {
	DueBlock uint64 `json:"due_block,string"`
	Coin     *Coin  `json:"coin"`
	Value    string `json:"value"`
}

type LockStakeData struct {
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
	case "type.googleapis.com/api_pb.LockStakeData":
		return &LockStakeData{}
	case "type.googleapis.com/api_pb.LockData":
		return &LockData{}
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
