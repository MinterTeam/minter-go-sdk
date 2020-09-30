// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ValidatorsResponseResult validators response result
//
// swagger:model ValidatorsResponseResult
type ValidatorsResponseResult struct {

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// voting power
	VotingPower uint64 `json:"voting_power,omitempty,string"`
}

// Validate validates this validators response result
func (m *ValidatorsResponseResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ValidatorsResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValidatorsResponseResult) UnmarshalBinary(b []byte) error {
	var res ValidatorsResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
