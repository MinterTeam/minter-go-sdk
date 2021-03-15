// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// BlockField block field
//
// swagger:model BlockField
type BlockField string

const (

	// BlockFieldTransactions captures enum value "transactions"
	BlockFieldTransactions BlockField = "transactions"

	// BlockFieldMissed captures enum value "missed"
	BlockFieldMissed BlockField = "missed"

	// BlockFieldBlockReward captures enum value "block_reward"
	BlockFieldBlockReward BlockField = "block_reward"

	// BlockFieldSize captures enum value "size"
	BlockFieldSize BlockField = "size"

	// BlockFieldProposer captures enum value "proposer"
	BlockFieldProposer BlockField = "proposer"

	// BlockFieldValidators captures enum value "validators"
	BlockFieldValidators BlockField = "validators"

	// BlockFieldEvidence captures enum value "evidence"
	BlockFieldEvidence BlockField = "evidence"
)

// for schema
var blockFieldEnum []interface{}

func init() {
	var res []BlockField
	if err := json.Unmarshal([]byte(`["transactions","missed","block_reward","size","proposer","validators","evidence"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blockFieldEnum = append(blockFieldEnum, v)
	}
}

func (m BlockField) validateBlockFieldEnum(path, location string, value BlockField) error {
	if err := validate.EnumCase(path, location, value, blockFieldEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this block field
func (m BlockField) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBlockFieldEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}