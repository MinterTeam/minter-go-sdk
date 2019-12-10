package api

type AddressesResponse struct {
	Jsonrpc string           `json:"jsonrpc"`
	ID      string           `json:"id,omitempty"`
	Result  *AddressesResult `json:"result,omitempty"`
	Error   *Error           `json:"error,omitempty"`
}

type AddressesResult struct {
}

func (a *Api) Addresses(address string, height int) (*AddressesResult, error) {
	//todo
	return nil, nil
}
