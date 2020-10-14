package models

import (
	"github.com/go-openapi/swag"
)

// AsTransactionResponse returns TransactionResponse
func (m *BlockResponseTransaction) AsTransactionResponse() (*TransactionResponse, error) {
	binary, err := m.MarshalBinary()
	if err != nil {
		return nil, err
	}

	var res TransactionResponse
	if err := swag.ReadJSON(binary, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
