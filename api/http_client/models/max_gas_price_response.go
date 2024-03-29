// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MaxGasPriceResponse max gas price response
//
// swagger:model MaxGasPriceResponse
type MaxGasPriceResponse struct {

	// max gas price
	MaxGasPrice uint64 `json:"max_gas_price,omitempty,string"`
}

// Validate validates this max gas price response
func (m *MaxGasPriceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this max gas price response based on context it is used
func (m *MaxGasPriceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MaxGasPriceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaxGasPriceResponse) UnmarshalBinary(b []byte) error {
	var res MaxGasPriceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
