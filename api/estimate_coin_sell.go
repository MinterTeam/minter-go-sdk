package api

import (
	"encoding/json"
	"strconv"
)

type EstimateCoinSellResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		WillGet    string `json:"will_get"`
		Commission string `json:"commission"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}

func (a *Api) EstimateCoinSell(coinToSell string, valueToSell string, coinToBuy string, height int) (*EstimateCoinSellResponse, error) {

	params := make(map[string]string)
	params["coin_to_sell"] = coinToSell
	params["value_to_sell"] = valueToSell
	params["coin_to_buy"] = coinToBuy
	if height > 0 {
		params["height"] = strconv.Itoa(height)
	}

	res, err := a.client.R().SetQueryParams(params).Get("/estimate_coin_sell")
	if err != nil {
		return nil, err
	}

	result := new(EstimateCoinSellResponse)
	err = json.Unmarshal(res.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
