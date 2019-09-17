package api

import (
	"encoding/json"
	"strconv"
)

type EstimateCoinBuyResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		WillGet    string `json:"will_get"`
		Commission string `json:"commission"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

func (a *Api) EstimateCoinBuy(coinToSell string, valueToBuy string, coinToBuy string, height int) (*EstimateCoinBuyResponse, error) {

	params := make(map[string]string)
	params["coin_to_sell"] = coinToSell
	params["value_to_buy"] = valueToBuy
	params["coin_to_buy"] = coinToBuy
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/estimate_coin_buy")
	if err != nil {
		return nil, err
	}

	result := new(EstimateCoinBuyResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
