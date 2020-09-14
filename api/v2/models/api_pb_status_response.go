// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPbStatusResponse api pb status response
//
// swagger:model api_pbStatusResponse
type APIPbStatusResponse struct {

	// catching up
	CatchingUp bool `json:"catching_up,omitempty"`

	// keep last states
	KeepLastStates string `json:"keep_last_states,omitempty"`

	// latest app hash
	LatestAppHash string `json:"latest_app_hash,omitempty"`

	// latest block hash
	LatestBlockHash string `json:"latest_block_hash,omitempty"`

	// latest block height
	LatestBlockHeight string `json:"latest_block_height,omitempty"`

	// latest block time
	LatestBlockTime string `json:"latest_block_time,omitempty"`

	// node id
	NodeID string `json:"node_id,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// total slashed
	TotalSlashed string `json:"total_slashed,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this api pb status response
func (m *APIPbStatusResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPbStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPbStatusResponse) UnmarshalBinary(b []byte) error {
	var res APIPbStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}