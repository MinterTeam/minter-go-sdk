// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenesisResponseAppStateVersion genesis response app state version
//
// swagger:model GenesisResponseAppStateVersion
type GenesisResponseAppStateVersion struct {

	// height
	Height uint64 `json:"height,omitempty,string"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this genesis response app state version
func (m *GenesisResponseAppStateVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this genesis response app state version based on context it is used
func (m *GenesisResponseAppStateVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenesisResponseAppStateVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenesisResponseAppStateVersion) UnmarshalBinary(b []byte) error {
	var res GenesisResponseAppStateVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
