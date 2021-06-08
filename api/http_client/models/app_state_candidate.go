// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppStateCandidate app state candidate
//
// swagger:model AppStateCandidate
type AppStateCandidate struct {

	// commission
	Commission uint64 `json:"commission,omitempty,string"`

	// control address
	ControlAddress string `json:"control_address,omitempty"`

	// id
	ID uint64 `json:"id,omitempty,string"`

	// jailed until
	JailedUntil int64 `json:"jailed_until,omitempty,string"`

	// last edit commission height
	LastEditCommissionHeight int64 `json:"last_edit_commission_height,omitempty,string"`

	// owner address
	OwnerAddress string `json:"owner_address,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// reward address
	RewardAddress string `json:"reward_address,omitempty"`

	// stakes
	Stakes []*AppStateCandidateStake `json:"stakes"`

	// status
	Status int64 `json:"status,omitempty,string"`

	// total bip stake
	TotalBipStake string `json:"total_bip_stake,omitempty"`

	// updates
	Updates []*AppStateCandidateStake `json:"updates"`
}

// Validate validates this app state candidate
func (m *AppStateCandidate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStakes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppStateCandidate) validateStakes(formats strfmt.Registry) error {

	if swag.IsZero(m.Stakes) { // not required
		return nil
	}

	for i := 0; i < len(m.Stakes); i++ {
		if swag.IsZero(m.Stakes[i]) { // not required
			continue
		}

		if m.Stakes[i] != nil {
			if err := m.Stakes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stakes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppStateCandidate) validateUpdates(formats strfmt.Registry) error {

	if swag.IsZero(m.Updates) { // not required
		return nil
	}

	for i := 0; i < len(m.Updates); i++ {
		if swag.IsZero(m.Updates[i]) { // not required
			continue
		}

		if m.Updates[i] != nil {
			if err := m.Updates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppStateCandidate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppStateCandidate) UnmarshalBinary(b []byte) error {
	var res AppStateCandidate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
