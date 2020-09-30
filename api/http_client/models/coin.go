// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Coin coin
//
// swagger:model Coin
type Coin struct {

	// id
	ID uint64 `json:"id,omitempty,string"`

	// symbol
	Symbol string `json:"symbol,omitempty"`
}

// Validate validates this coin
func (m *Coin) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Coin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Coin) UnmarshalBinary(b []byte) error {
	var res Coin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
