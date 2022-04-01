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

// AddressesResponseResult addresses response result
//
// swagger:model AddressesResponseResult
type AddressesResponseResult struct {

	// balance
	Balance []*AddressBalance `json:"balance"`

	// bip value
	BipValue string `json:"bip_value,omitempty"`

	// Filled in when request delegated
	Delegated []*AddressDelegatedBalance `json:"delegated"`

	// locked stake until block
	LockedStakeUntilBlock uint64 `json:"locked_stake_until_block,omitempty,string"`

	// multisig
	Multisig *Multisig `json:"multisig,omitempty"`

	// Sum of balance and delegated by coins. Filled in when request delegated
	Total []*AddressBalance `json:"total"`

	// transaction count
	TransactionCount uint64 `json:"transaction_count,omitempty,string"`
}

// Validate validates this addresses response result
func (m *AddressesResponseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMultisig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressesResponseResult) validateBalance(formats strfmt.Registry) error {
	if swag.IsZero(m.Balance) { // not required
		return nil
	}

	for i := 0; i < len(m.Balance); i++ {
		if swag.IsZero(m.Balance[i]) { // not required
			continue
		}

		if m.Balance[i] != nil {
			if err := m.Balance[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("balance" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddressesResponseResult) validateDelegated(formats strfmt.Registry) error {
	if swag.IsZero(m.Delegated) { // not required
		return nil
	}

	for i := 0; i < len(m.Delegated); i++ {
		if swag.IsZero(m.Delegated[i]) { // not required
			continue
		}

		if m.Delegated[i] != nil {
			if err := m.Delegated[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegated" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddressesResponseResult) validateMultisig(formats strfmt.Registry) error {
	if swag.IsZero(m.Multisig) { // not required
		return nil
	}

	if m.Multisig != nil {
		if err := m.Multisig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multisig")
			}
			return err
		}
	}

	return nil
}

func (m *AddressesResponseResult) validateTotal(formats strfmt.Registry) error {
	if swag.IsZero(m.Total) { // not required
		return nil
	}

	for i := 0; i < len(m.Total); i++ {
		if swag.IsZero(m.Total[i]) { // not required
			continue
		}

		if m.Total[i] != nil {
			if err := m.Total[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("total" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this addresses response result based on the context it is used
func (m *AddressesResponseResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDelegated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMultisig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressesResponseResult) contextValidateBalance(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Balance); i++ {

		if m.Balance[i] != nil {
			if err := m.Balance[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("balance" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddressesResponseResult) contextValidateDelegated(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Delegated); i++ {

		if m.Delegated[i] != nil {
			if err := m.Delegated[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegated" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddressesResponseResult) contextValidateMultisig(ctx context.Context, formats strfmt.Registry) error {

	if m.Multisig != nil {
		if err := m.Multisig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multisig")
			}
			return err
		}
	}

	return nil
}

func (m *AddressesResponseResult) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Total); i++ {

		if m.Total[i] != nil {
			if err := m.Total[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("total" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddressesResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressesResponseResult) UnmarshalBinary(b []byte) error {
	var res AddressesResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
