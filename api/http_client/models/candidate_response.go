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

// CandidateResponse candidate response
//
// swagger:model CandidateResponse
type CandidateResponse struct {

	// Commission (from 0 to 100) from rewards which delegators will pay to validator
	Commission uint64 `json:"commission,omitempty,string"`

	// Address that allows one to start the candidate by sending the SetCandidateOnline transaction or stop it by sending the SetCandidateOffline transaction.
	ControlAddress string `json:"control_address,omitempty"`

	// id
	ID uint64 `json:"id,omitempty,string"`

	// jailed until
	JailedUntil uint64 `json:"jailed_until,omitempty,string"`

	// Smallest steak size. Note: filled in when request includes_stakes
	MinStake string `json:"min_stake,omitempty"`

	// Address that allows one to start the candidate by sending the SetCandidateOnline transaction or stop it by sending the SetCandidateOffline transaction. It also enables the owner to edit the node by sending EditCandidate.
	OwnerAddress string `json:"owner_address,omitempty"`

	// Public key of a candidate
	PublicKey string `json:"public_key,omitempty"`

	// Address where validator’s rewards go to.
	RewardAddress string `json:"reward_address,omitempty"`

	// List of stakes. Note: filled in when request includes_stakes
	Stakes []*CandidateResponseStake `json:"stakes"`

	// Candidate status. Available values: offline = 1, online = 2
	Status uint64 `json:"status,omitempty,string"`

	// Total stake of a candidate
	TotalStake string `json:"total_stake,omitempty"`

	// Number of unique wallets in steaks. Note: filled in when request includes_stakes
	UniqUsers uint64 `json:"uniq_users,omitempty,string"`

	// Number of occupied steak slots. Note: filled in when request includes_stakes
	UsedSlots uint64 `json:"used_slots,omitempty,string"`

	// Is a validator at the current height
	Validator bool `json:"validator,omitempty"`
}

// Validate validates this candidate response
func (m *CandidateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStakes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CandidateResponse) validateStakes(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CandidateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CandidateResponse) UnmarshalBinary(b []byte) error {
	var res CandidateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
