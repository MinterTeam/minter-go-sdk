package api

type EstimateCoinSellAllResponse struct {
	Jsonrpc string                     `json:"jsonrpc"`
	ID      string                     `json:"id,omitempty"`
	Result  *EstimateCoinSellAllResult `json:"result,omitempty"`
	Error   *Error                     `json:"error,omitempty"`
}
type EstimateCoinSellAllResult struct {
}

func (a *Api) EstimateCoinSellAll(coinToSell string, coinToBuy string, valueToSell string, gasPrice string, height int) (*EstimateCoinSellAllResult, error) {
	//todo
	return nil, nil
}
