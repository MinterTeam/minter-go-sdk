package api

import (
	"encoding/json"
	"strconv"
)

type EstimateCoinSellResponse struct {
	Jsonrpc string                  `json:"jsonrpc"`
	ID      string                  `json:"id,omitempty"`
	Result  *EstimateCoinSellResult `json:"result,omitempty"`
	Error   *Error                  `json:"error,omitempty"`
}
type EstimateCoinSellResult struct {
	WillGet    string `json:"will_get"`
	Commission string `json:"commission"`
}

// Return estimate of sell coin transaction.
func (a *Api) EstimateCoinSell(coinToSell string, valueToSell string, coinToBuy string) (*EstimateCoinSellResult, error) {
	return a.EstimateCoinSellAtHeight(coinToSell, valueToSell, coinToBuy, LatestBlockHeight)
}

// Return estimate of sell coin transaction.
func (a *Api) EstimateCoinSellAtHeight(coinToSell string, valueToSell string, coinToBuy string, height int) (*EstimateCoinSellResult, error) {

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
	if err := hasError(res); err != nil {
		return nil, err
	}

	response := new(EstimateCoinSellResponse)
	err = json.Unmarshal(res.Body(), response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.Result, nil
}
