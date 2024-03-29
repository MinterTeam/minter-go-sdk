// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EstimateTxCommissionResponse estimate tx commission response
//
// swagger:model EstimateTxCommissionResponse
type EstimateTxCommissionResponse struct {

	// commission
	Commission string `json:"commission,omitempty"`
}

// Validate validates this estimate tx commission response
func (m *EstimateTxCommissionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this estimate tx commission response based on context it is used
func (m *EstimateTxCommissionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EstimateTxCommissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EstimateTxCommissionResponse) UnmarshalBinary(b []byte) error {
	var res EstimateTxCommissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
