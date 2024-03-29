// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommissionVotesResponse commission votes response
//
// swagger:model CommissionVotesResponse
type CommissionVotesResponse struct {

	// votes
	Votes []*CommissionVotesResponseVote `json:"votes"`
}

// Validate validates this commission votes response
func (m *CommissionVotesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVotes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommissionVotesResponse) validateVotes(formats strfmt.Registry) error {
	if swag.IsZero(m.Votes) { // not required
		return nil
	}

	for i := 0; i < len(m.Votes); i++ {
		if swag.IsZero(m.Votes[i]) { // not required
			continue
		}

		if m.Votes[i] != nil {
			if err := m.Votes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("votes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this commission votes response based on the context it is used
func (m *CommissionVotesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVotes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommissionVotesResponse) contextValidateVotes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Votes); i++ {

		if m.Votes[i] != nil {
			if err := m.Votes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("votes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommissionVotesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommissionVotesResponse) UnmarshalBinary(b []byte) error {
	var res CommissionVotesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
