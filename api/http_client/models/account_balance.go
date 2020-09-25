// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountBalance account balance
//
// swagger:model AccountBalance
type AccountBalance struct {

	// coin
	Coin string `json:"coin,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this account balance
func (m *AccountBalance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountBalance) UnmarshalBinary(b []byte) error {
	var res AccountBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
