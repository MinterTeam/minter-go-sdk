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

// EstimateCoinSellAllResponse estimate coin sell all response
//
// swagger:model EstimateCoinSellAllResponse
type EstimateCoinSellAllResponse struct {

	// swap from
	SwapFrom *SwapFrom `json:"swap_from,omitempty"`

	// will get
	WillGet string `json:"will_get,omitempty"`
}

// Validate validates this estimate coin sell all response
func (m *EstimateCoinSellAllResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSwapFrom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EstimateCoinSellAllResponse) validateSwapFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.SwapFrom) { // not required
		return nil
	}

	if m.SwapFrom != nil {
		if err := m.SwapFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swap_from")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this estimate coin sell all response based on the context it is used
func (m *EstimateCoinSellAllResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSwapFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EstimateCoinSellAllResponse) contextValidateSwapFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.SwapFrom != nil {
		if err := m.SwapFrom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swap_from")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EstimateCoinSellAllResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EstimateCoinSellAllResponse) UnmarshalBinary(b []byte) error {
	var res EstimateCoinSellAllResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
