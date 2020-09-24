// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	Address(params *AddressParams) (*AddressOK, error)

	Addresses(params *AddressesParams) (*AddressesOK, error)

	Block(params *BlockParams) (*BlockOK, error)

	Candidate(params *CandidateParams) (*CandidateOK, error)

	Candidates(params *CandidatesParams) (*CandidatesOK, error)

	CoinInfo(params *CoinInfoParams) (*CoinInfoOK, error)

	CoinInfoByID(params *CoinInfoByIDParams) (*CoinInfoByIDOK, error)

	EstimateCoinBuy(params *EstimateCoinBuyParams) (*EstimateCoinBuyOK, error)

	EstimateCoinSell(params *EstimateCoinSellParams) (*EstimateCoinSellOK, error)

	EstimateCoinSellAll(params *EstimateCoinSellAllParams) (*EstimateCoinSellAllOK, error)

	EstimateTxCommission(params *EstimateTxCommissionParams) (*EstimateTxCommissionOK, error)

	Events(params *EventsParams) (*EventsOK, error)

	Frozen(params *FrozenParams) (*FrozenOK, error)

	Genesis(params *GenesisParams) (*GenesisOK, error)

	Halts(params *HaltsParams) (*HaltsOK, error)

	MaxGasPrice(params *MaxGasPriceParams) (*MaxGasPriceOK, error)

	MinGasPrice(params *MinGasPriceParams) (*MinGasPriceOK, error)

	MissedBlocks(params *MissedBlocksParams) (*MissedBlocksOK, error)

	NetInfo(params *NetInfoParams) (*NetInfoOK, error)

	SendTransaction(params *SendTransactionParams) (*SendTransactionOK, error)

	SendTransaction2(params *SendTransaction2Params) (*SendTransaction2OK, error)

	Status(params *StatusParams) (*StatusOK, error)

	Subscribe(params *SubscribeParams) (*SubscribeOK, error)

	Transaction(params *TransactionParams) (*TransactionOK, error)

	Transactions(params *TransactionsParams) (*TransactionsOK, error)

	UnconfirmedTxs(params *UnconfirmedTxsParams) (*UnconfirmedTxsOK, error)

	Validators(params *ValidatorsParams) (*ValidatorsOK, error)

	WaitList(params *WaitListParams) (*WaitListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Address returns coins list balance and transaction count of an address
*/
func (a *Client) Address(params *AddressParams) (*AddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Address",
		Method:             "GET",
		PathPattern:        "/address/{address}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddressReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddressDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Addresses returns list of addresses
*/
func (a *Client) Addresses(params *AddressesParams) (*AddressesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddressesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Addresses",
		Method:             "GET",
		PathPattern:        "/addresses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddressesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddressesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddressesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Block returns block data at given height
*/
func (a *Client) Block(params *BlockParams) (*BlockOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBlockParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Block",
		Method:             "GET",
		PathPattern:        "/block/{height}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BlockReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BlockOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BlockDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Candidate returns candidate s info by provided public key
*/
func (a *Client) Candidate(params *CandidateParams) (*CandidateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCandidateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Candidate",
		Method:             "GET",
		PathPattern:        "/candidate/{public_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CandidateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CandidateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CandidateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Candidates returns list of candidates
*/
func (a *Client) Candidates(params *CandidatesParams) (*CandidatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCandidatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Candidates",
		Method:             "GET",
		PathPattern:        "/candidates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CandidatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CandidatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CandidatesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CoinInfo returns information about coin symbol
*/
func (a *Client) CoinInfo(params *CoinInfoParams) (*CoinInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCoinInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CoinInfo",
		Method:             "GET",
		PathPattern:        "/coin_info/{symbol}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CoinInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CoinInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CoinInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CoinInfoByID returns information about coin ID
*/
func (a *Client) CoinInfoByID(params *CoinInfoByIDParams) (*CoinInfoByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCoinInfoByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CoinInfoById",
		Method:             "GET",
		PathPattern:        "/coin_info_by_id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CoinInfoByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CoinInfoByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CoinInfoByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EstimateCoinBuy returns estimate of buy coin transaction
*/
func (a *Client) EstimateCoinBuy(params *EstimateCoinBuyParams) (*EstimateCoinBuyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEstimateCoinBuyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EstimateCoinBuy",
		Method:             "GET",
		PathPattern:        "/estimate_coin_buy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EstimateCoinBuyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EstimateCoinBuyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EstimateCoinBuyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EstimateCoinSell returns estimate of sell coin transaction
*/
func (a *Client) EstimateCoinSell(params *EstimateCoinSellParams) (*EstimateCoinSellOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEstimateCoinSellParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EstimateCoinSell",
		Method:             "GET",
		PathPattern:        "/estimate_coin_sell",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EstimateCoinSellReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EstimateCoinSellOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EstimateCoinSellDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EstimateCoinSellAll returns estimate of sell all coin transaction
*/
func (a *Client) EstimateCoinSellAll(params *EstimateCoinSellAllParams) (*EstimateCoinSellAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEstimateCoinSellAllParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EstimateCoinSellAll",
		Method:             "GET",
		PathPattern:        "/estimate_coin_sell_all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EstimateCoinSellAllReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EstimateCoinSellAllOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EstimateCoinSellAllDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EstimateTxCommission returns estimate of transaction
*/
func (a *Client) EstimateTxCommission(params *EstimateTxCommissionParams) (*EstimateTxCommissionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEstimateTxCommissionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EstimateTxCommission",
		Method:             "GET",
		PathPattern:        "/estimate_tx_commission/{tx}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EstimateTxCommissionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EstimateTxCommissionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EstimateTxCommissionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Events returns events at given height
*/
func (a *Client) Events(params *EventsParams) (*EventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Events",
		Method:             "GET",
		PathPattern:        "/events/{height}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EventsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EventsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Frozen returns frozen balance
*/
func (a *Client) Frozen(params *FrozenParams) (*FrozenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrozenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Frozen",
		Method:             "GET",
		PathPattern:        "/frozen/{address}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FrozenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FrozenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FrozenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Genesis returns genesis file
*/
func (a *Client) Genesis(params *GenesisParams) (*GenesisOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenesisParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Genesis",
		Method:             "GET",
		PathPattern:        "/genesis",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GenesisReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GenesisOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GenesisDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Halts returns the candidate votes for stopping the network at block
*/
func (a *Client) Halts(params *HaltsParams) (*HaltsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHaltsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Halts",
		Method:             "GET",
		PathPattern:        "/halts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HaltsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HaltsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HaltsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MaxGasPrice returns current max gas
*/
func (a *Client) MaxGasPrice(params *MaxGasPriceParams) (*MaxGasPriceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMaxGasPriceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "MaxGasPrice",
		Method:             "GET",
		PathPattern:        "/max_gas_price",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MaxGasPriceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MaxGasPriceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MaxGasPriceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MinGasPrice returns current min gas price
*/
func (a *Client) MinGasPrice(params *MinGasPriceParams) (*MinGasPriceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMinGasPriceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "MinGasPrice",
		Method:             "GET",
		PathPattern:        "/min_gas_price",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MinGasPriceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MinGasPriceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MinGasPriceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MissedBlocks returns missed blocks by validator public key
*/
func (a *Client) MissedBlocks(params *MissedBlocksParams) (*MissedBlocksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMissedBlocksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "MissedBlocks",
		Method:             "GET",
		PathPattern:        "/missed_blocks/{public_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MissedBlocksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MissedBlocksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MissedBlocksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NetInfo returns network info
*/
func (a *Client) NetInfo(params *NetInfoParams) (*NetInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetInfo",
		Method:             "GET",
		PathPattern:        "/net_info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NetInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NetInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NetInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SendTransaction returns the result of sending signed tx to ensure that transaction was successfully committed to the blockchain you need to find the transaction by the hash and ensure that the status code equals to 0
*/
func (a *Client) SendTransaction(params *SendTransactionParams) (*SendTransactionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendTransactionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SendTransaction",
		Method:             "GET",
		PathPattern:        "/send_transaction/{tx}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SendTransactionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SendTransactionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SendTransactionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SendTransaction2 returns the result of sending signed tx to ensure that transaction was successfully committed to the blockchain you need to find the transaction by the hash and ensure that the status code equals to 0
*/
func (a *Client) SendTransaction2(params *SendTransaction2Params) (*SendTransaction2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendTransaction2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SendTransaction2",
		Method:             "POST",
		PathPattern:        "/send_transaction",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SendTransaction2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SendTransaction2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SendTransaction2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Status returns node status including pubkey latest block
*/
func (a *Client) Status(params *StatusParams) (*StatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Status",
		Method:             "GET",
		PathPattern:        "/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Subscribe returns a subscription for events by query

  Only supported in WS and gRPC methods
*/
func (a *Client) Subscribe(params *SubscribeParams) (*SubscribeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscribeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Subscribe",
		Method:             "GET",
		PathPattern:        "/subscribe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SubscribeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscribeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SubscribeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Transaction returns transaction info
*/
func (a *Client) Transaction(params *TransactionParams) (*TransactionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransactionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Transaction",
		Method:             "GET",
		PathPattern:        "/transaction/{hash}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TransactionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TransactionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TransactionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Transactions returns transactions by query
*/
func (a *Client) Transactions(params *TransactionsParams) (*TransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransactionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Transactions",
		Method:             "GET",
		PathPattern:        "/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TransactionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TransactionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UnconfirmedTxs returns unconfirmed transactions
*/
func (a *Client) UnconfirmedTxs(params *UnconfirmedTxsParams) (*UnconfirmedTxsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnconfirmedTxsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UnconfirmedTxs",
		Method:             "GET",
		PathPattern:        "/unconfirmed_txs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UnconfirmedTxsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnconfirmedTxsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UnconfirmedTxsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Validators returns list of active validators
*/
func (a *Client) Validators(params *ValidatorsParams) (*ValidatorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidatorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Validators",
		Method:             "GET",
		PathPattern:        "/validators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidatorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidatorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ValidatorsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  WaitList returns the list of address stakes in waitlist
*/
func (a *Client) WaitList(params *WaitListParams) (*WaitListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWaitListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "WaitList",
		Method:             "GET",
		PathPattern:        "/waitlist/{public_key}/{address}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WaitListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WaitListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WaitListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
