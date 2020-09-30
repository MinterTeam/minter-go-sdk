// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SendTransactionResponse send transaction response
//
// swagger:model SendTransactionResponse
type SendTransactionResponse struct {

	// code
	Code uint64 `json:"code,omitempty,string"`

	// hash
	Hash string `json:"hash,omitempty"`

	// log
	Log string `json:"log,omitempty"`
}

// Validate validates this send transaction response
func (m *SendTransactionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SendTransactionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendTransactionResponse) UnmarshalBinary(b []byte) error {
	var res SendTransactionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
