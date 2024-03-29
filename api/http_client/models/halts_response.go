// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HaltsResponse halts response
//
// swagger:model HaltsResponse
type HaltsResponse struct {

	// public keys
	PublicKeys []string `json:"public_keys"`
}

// Validate validates this halts response
func (m *HaltsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this halts response based on context it is used
func (m *HaltsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HaltsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HaltsResponse) UnmarshalBinary(b []byte) error {
	var res HaltsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
