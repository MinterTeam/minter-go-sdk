// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PriceCommissionResponse price commission response
//
// swagger:model PriceCommissionResponse
type PriceCommissionResponse struct {

	// add limit order
	AddLimitOrder string `json:"add_limit_order,omitempty"`

	// add liquidity
	AddLiquidity string `json:"add_liquidity,omitempty"`

	// burn token
	BurnToken string `json:"burn_token,omitempty"`

	// buy bancor
	BuyBancor string `json:"buy_bancor,omitempty"`

	// buy pool base
	BuyPoolBase string `json:"buy_pool_base,omitempty"`

	// buy pool delta
	BuyPoolDelta string `json:"buy_pool_delta,omitempty"`

	// coin
	Coin *Coin `json:"coin,omitempty"`

	// create coin
	CreateCoin string `json:"create_coin,omitempty"`

	// create multisig
	CreateMultisig string `json:"create_multisig,omitempty"`

	// create swap pool
	CreateSwapPool string `json:"create_swap_pool,omitempty"`

	// create ticker3
	CreateTicker3 string `json:"create_ticker3,omitempty"`

	// create ticker4
	CreateTicker4 string `json:"create_ticker4,omitempty"`

	// create ticker5
	CreateTicker5 string `json:"create_ticker5,omitempty"`

	// create ticker6
	CreateTicker6 string `json:"create_ticker6,omitempty"`

	// create ticker7 10
	CreateTicker710 string `json:"create_ticker7_10,omitempty"`

	// create token
	CreateToken string `json:"create_token,omitempty"`

	// declare candidacy
	DeclareCandidacy string `json:"declare_candidacy,omitempty"`

	// delegate
	Delegate string `json:"delegate,omitempty"`

	// edit candidate
	EditCandidate string `json:"edit_candidate,omitempty"`

	// edit candidate commission
	EditCandidateCommission string `json:"edit_candidate_commission,omitempty"`

	// edit candidate public key
	EditCandidatePublicKey string `json:"edit_candidate_public_key,omitempty"`

	// edit multisig
	EditMultisig string `json:"edit_multisig,omitempty"`

	// edit ticker owner
	EditTickerOwner string `json:"edit_ticker_owner,omitempty"`

	// failed tx
	FailedTx string `json:"failed_tx,omitempty"`

	// mint token
	MintToken string `json:"mint_token,omitempty"`

	// multisend base
	MultisendBase string `json:"multisend_base,omitempty"`

	// multisend delta
	MultisendDelta string `json:"multisend_delta,omitempty"`

	// payload byte
	PayloadByte string `json:"payload_byte,omitempty"`

	// recreate coin
	RecreateCoin string `json:"recreate_coin,omitempty"`

	// recreate token
	RecreateToken string `json:"recreate_token,omitempty"`

	// redeem check
	RedeemCheck string `json:"redeem_check,omitempty"`

	// remove limit order
	RemoveLimitOrder string `json:"remove_limit_order,omitempty"`

	// remove liquidity
	RemoveLiquidity string `json:"remove_liquidity,omitempty"`

	// sell all bancor
	SellAllBancor string `json:"sell_all_bancor,omitempty"`

	// sell all pool base
	SellAllPoolBase string `json:"sell_all_pool_base,omitempty"`

	// sell all pool delta
	SellAllPoolDelta string `json:"sell_all_pool_delta,omitempty"`

	// sell bancor
	SellBancor string `json:"sell_bancor,omitempty"`

	// sell pool base
	SellPoolBase string `json:"sell_pool_base,omitempty"`

	// sell pool delta
	SellPoolDelta string `json:"sell_pool_delta,omitempty"`

	// send
	Send string `json:"send,omitempty"`

	// set candidate off
	SetCandidateOff string `json:"set_candidate_off,omitempty"`

	// set candidate on
	SetCandidateOn string `json:"set_candidate_on,omitempty"`

	// set halt block
	SetHaltBlock string `json:"set_halt_block,omitempty"`

	// unbond
	Unbond string `json:"unbond,omitempty"`

	// vote commission
	VoteCommission string `json:"vote_commission,omitempty"`

	// vote update
	VoteUpdate string `json:"vote_update,omitempty"`
}

// Validate validates this price commission response
func (m *PriceCommissionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PriceCommissionResponse) validateCoin(formats strfmt.Registry) error {

	if swag.IsZero(m.Coin) { // not required
		return nil
	}

	if m.Coin != nil {
		if err := m.Coin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PriceCommissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PriceCommissionResponse) UnmarshalBinary(b []byte) error {
	var res PriceCommissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
