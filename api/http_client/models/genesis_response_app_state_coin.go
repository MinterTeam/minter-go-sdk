// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenesisResponseAppStateCoin genesis response app state coin
//
// swagger:model GenesisResponseAppStateCoin
type GenesisResponseAppStateCoin struct {

	// burnable
	Burnable bool `json:"burnable,omitempty"`

	// crr
	Crr uint64 `json:"crr,omitempty,string"`

	// id
	ID uint64 `json:"id,omitempty,string"`

	// max supply
	MaxSupply string `json:"max_supply,omitempty"`

	// mintable
	Mintable bool `json:"mintable,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner address
	OwnerAddress string `json:"owner_address,omitempty"`

	// reserve
	Reserve string `json:"reserve,omitempty"`

	// symbol
	Symbol string `json:"symbol,omitempty"`

	// version
	Version uint64 `json:"version,omitempty,string"`

	// volume
	Volume string `json:"volume,omitempty"`
}

// Validate validates this genesis response app state coin
func (m *GenesisResponseAppStateCoin) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this genesis response app state coin based on context it is used
func (m *GenesisResponseAppStateCoin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenesisResponseAppStateCoin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenesisResponseAppStateCoin) UnmarshalBinary(b []byte) error {
	var res GenesisResponseAppStateCoin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
