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

// FrozenResponseFrozen frozen response frozen
//
// swagger:model FrozenResponseFrozen
type FrozenResponseFrozen struct {

	// address
	Address string `json:"address,omitempty"`

	// candidate key
	CandidateKey string `json:"candidate_key,omitempty"`

	// coin
	Coin *Coin `json:"coin,omitempty"`

	// height
	Height uint64 `json:"height,omitempty,string"`

	// move to candidate key
	MoveToCandidateKey string `json:"move_to_candidate_key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this frozen response frozen
func (m *FrozenResponseFrozen) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrozenResponseFrozen) validateCoin(formats strfmt.Registry) error {
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

// ContextValidate validate this frozen response frozen based on the context it is used
func (m *FrozenResponseFrozen) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCoin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrozenResponseFrozen) contextValidateCoin(ctx context.Context, formats strfmt.Registry) error {

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
func (m *FrozenResponseFrozen) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrozenResponseFrozen) UnmarshalBinary(b []byte) error {
	var res FrozenResponseFrozen
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
