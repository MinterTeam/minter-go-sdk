// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddressBalance address balance
//
// swagger:model AddressBalance
type AddressBalance struct {

	// bip value
	BipValue string `json:"bip_value,omitempty"`

	// coin
	Coin *Coin `json:"coin,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this address balance
func (m *AddressBalance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressBalance) validateCoin(formats strfmt.Registry) error {
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

// ContextValidate validate this address balance based on the context it is used
func (m *AddressBalance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCoin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressBalance) contextValidateCoin(ctx context.Context, formats strfmt.Registry) error {

	if m.Coin != nil {
		if err := m.Coin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddressBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressBalance) UnmarshalBinary(b []byte) error {
	var res AddressBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
