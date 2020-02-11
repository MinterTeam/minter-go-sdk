package api

import (
	"encoding/json"
	"strconv"
)

type EstimateCoinSellAllResponse struct {
	Jsonrpc string                     `json:"jsonrpc"`
	ID      string                     `json:"id,omitempty"`
	Result  *EstimateCoinSellAllResult `json:"result,omitempty"`
	Error   *Error                     `json:"error,omitempty"`
}
type EstimateCoinSellAllResult struct {
	WillGet string `json:"will_get"`
}

func (a *Api) EstimateCoinSellAll(coinToSell string, coinToBuy string, valueToSell string, gasPrice int) (*EstimateCoinSellAllResult, error) {
	return a.EstimateCoinSellAllAtHeight(coinToSell, coinToBuy, valueToSell, gasPrice, LatestBlockHeight)
}

func (a *Api) EstimateCoinSellAllAtHeight(coinToSell string, coinToBuy string, valueToSell string, gasPrice int, height int) (*EstimateCoinSellAllResult, error) {
	params := make(map[string]string)
	params["height"] = strconv.Itoa(height)
	params["coin_to_sell"] = coinToSell
	params["coin_to_buy"] = coinToBuy
	params["value_to_sell"] = valueToSell
	params["gas_price"] = strconv.Itoa(gasPrice)

	res, err := a.client.R().SetQueryParams(params).Get("/estimate_coin_sell_all")
	if err != nil {
		return nil, err
	}
	if err := hasError(res); err != nil {
		return nil, err
	}

	response := new(EstimateCoinSellAllResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
