// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppStateCommissionVote app state commission vote
//
// swagger:model AppStateCommissionVote
type AppStateCommissionVote struct {

	// commission
	Commission *AppStateCommission `json:"commission,omitempty"`

	// height
	Height uint64 `json:"height,omitempty,string"`

	// votes
	Votes []string `json:"votes"`
}

// Validate validates this app state commission vote
func (m *AppStateCommissionVote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppStateCommissionVote) validateCommission(formats strfmt.Registry) error {
	if swag.IsZero(m.Commission) { // not required
		return nil
	}

	if m.Commission != nil {
		if err := m.Commission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commission")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app state commission vote based on the context it is used
func (m *AppStateCommissionVote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppStateCommissionVote) contextValidateCommission(ctx context.Context, formats strfmt.Registry) error {

	if m.Commission != nil {
		if err := m.Commission.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppStateCommissionVote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppStateCommissionVote) UnmarshalBinary(b []byte) error {
	var res AppStateCommissionVote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
