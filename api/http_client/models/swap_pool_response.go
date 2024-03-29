// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SwapPoolResponse swap pool response
//
// swagger:model SwapPoolResponse
type SwapPoolResponse struct {

	// amount0
	Amount0 string `json:"amount0,omitempty"`

	// amount1
	Amount1 string `json:"amount1,omitempty"`

	// id
	ID uint64 `json:"id,omitempty,string"`

	// liquidity
	Liquidity string `json:"liquidity,omitempty"`

	// price
	Price string `json:"price,omitempty"`
}

// Validate validates this swap pool response
func (m *SwapPoolResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this swap pool response based on context it is used
func (m *SwapPoolResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SwapPoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwapPoolResponse) UnmarshalBinary(b []byte) error {
	var res SwapPoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
