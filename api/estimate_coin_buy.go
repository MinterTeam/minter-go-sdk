package api

import (
	"encoding/json"
	"strconv"
)

type EstimateCoinBuyResponse struct {
	Jsonrpc string                 `json:"jsonrpc"`
	ID      string                 `json:"id,omitempty"`
	Result  *EstimateCoinBuyResult `json:"result,omitempty"`
	Error   *Error                 `json:"error,omitempty"`
}

type EstimateCoinBuyResult struct {
	WillPay    string `json:"will_pay"`
	Commission string `json:"commission"`
}

// Return estimate of buy coin transaction.
func (a *Api) EstimateCoinBuy(coinToSell string, valueToBuy string, coinToBuy string, height int) (*EstimateCoinBuyResult, error) {

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
	if res.IsError() {
		return nil, NewResponseError(res)
	}

	response := new(EstimateCoinBuyResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
