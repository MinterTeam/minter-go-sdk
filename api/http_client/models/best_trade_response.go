// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BestTradeResponse best trade response
//
// swagger:model BestTradeResponse
type BestTradeResponse struct {

	// path
	Path []uint64 `json:"path"`

	// result
	Result string `json:"result,omitempty"`
}

// Validate validates this best trade response
func (m *BestTradeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this best trade response based on context it is used
func (m *BestTradeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BestTradeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BestTradeResponse) UnmarshalBinary(b []byte) error {
	var res BestTradeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
