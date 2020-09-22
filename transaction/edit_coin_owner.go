package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// EditCoinOwnerData is a Data of Transaction for editing coin owner.
type EditCoinOwnerData struct {
	Symbol   Coin
	NewOwner [20]byte
}

// NewEditCoinOwnerData returns new EditCoinOwnerData of Transaction for editing coin owner.
func NewEditCoinOwnerData() *EditCoinOwnerData {
	return &EditCoinOwnerData{}
}

// SetSymbol sets symbol of a coin.
func (d *EditCoinOwnerData) SetSymbol(symbol string) *EditCoinOwnerData {
	copy(d.Symbol[:], symbol)
	return d
}

// SetNewOwner sets new owner address of a coin.
func (d *EditCoinOwnerData) SetNewOwner(address string) (*EditCoinOwnerData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.NewOwner[:], bytes)
	return d, nil
}

// MustSetNewOwner tries to set address of candidate and panics on error.
func (d *EditCoinOwnerData) MustSetNewOwner(address string) *EditCoinOwnerData {
	_, err := d.SetNewOwner(address)
	if err != nil {
		panic(err)
	}
	return d
}

// Type returns Data type of the transaction.
func (d *EditCoinOwnerData) Type() Type {
	return TypeEditCoinOwner
}

// Fee returns commission of transaction Data
func (d *EditCoinOwnerData) fee() fee {
	return feeTypeEditCoinOwner
}

// Encode returns the byte representation of a transaction Data.
func (d *EditCoinOwnerData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
