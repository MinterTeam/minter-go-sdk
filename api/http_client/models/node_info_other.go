// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeInfoOther node info other
//
// swagger:model NodeInfoOther
type NodeInfoOther struct {

	// rpc address
	RPCAddress string `json:"rpc_address,omitempty"`

	// tx index
	TxIndex string `json:"tx_index,omitempty"`
}

// Validate validates this node info other
func (m *NodeInfoOther) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node info other based on context it is used
func (m *NodeInfoOther) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeInfoOther) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeInfoOther) UnmarshalBinary(b []byte) error {
	var res NodeInfoOther
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
