// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransactionResponse transaction response
//
// swagger:model TransactionResponse
type TransactionResponse struct {

	// code
	Code uint64 `json:"code,omitempty,string"`

	// data
	Data *ProtobufAny `json:"data,omitempty"`

	// from
	From string `json:"from,omitempty"`

	// gas
	Gas uint64 `json:"gas,omitempty,string"`

	// gas coin
	GasCoin *Coin `json:"gas_coin,omitempty"`

	// gas price
	GasPrice uint64 `json:"gas_price,omitempty,string"`

	// hash
	Hash string `json:"hash,omitempty"`

	// height
	Height uint64 `json:"height,omitempty,string"`

	// index
	Index uint64 `json:"index,omitempty,string"`

	// log
	Log string `json:"log,omitempty"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty,string"`

	// payload
	// Format: byte
	Payload strfmt.Base64 `json:"payload,omitempty"`

	// raw tx
	RawTx string `json:"raw_tx,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`

	// type
	Type uint64 `json:"type,omitempty,string"`
}

// Validate validates this transaction response
func (m *TransactionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasCoin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *TransactionResponse) validateGasCoin(formats strfmt.Registry) error {

	if swag.IsZero(m.GasCoin) { // not required
		return nil
	}

	if m.GasCoin != nil {
		if err := m.GasCoin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gas_coin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionResponse) UnmarshalBinary(b []byte) error {
	var res TransactionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
