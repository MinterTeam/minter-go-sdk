package transaction

import "github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"

type EditCommissionData struct {
	PubKey     PublicKey
	Commission uint32
}

// Type returns Data type of the transaction.
func (d *EditCommissionData) Type() Type {
	return TypeEditCommission
}

// Fee returns commission of transaction Data
func (d *EditCommissionData) Fee() Fee {
	return feeTypeEditCommissionData
}

// Encode returns the byte representation of a transaction Data.
func (d *EditCommissionData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
