package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// swagger:model protobufAny
type ProtobufAny map[string]interface{}

// Validate validates this protobuf any
func (m *ProtobufAny) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtobufAny) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtobufAny) UnmarshalBinary(b []byte) error {
	var res ProtobufAny
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res

	return nil
}

func (m *ProtobufAny) UnmarshalTo(i Data) error {
	binary, err := m.MarshalBinary()
	if err != nil {
		return err
	}

	if err := swag.ReadJSON(binary, i); err != nil {
		return err
	}

	return nil
}

func (m *ProtobufAny) UnmarshalNew() (Data, error) {
	return convertToData(m)
}
