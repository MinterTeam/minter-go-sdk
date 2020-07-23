package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

type ChangeOwnerData struct {
	Symbol   CoinSymbol
	NewOwner [20]byte
}

func NewChangeOwnerData() *ChangeOwnerData {
	return &ChangeOwnerData{}
}

func (d *ChangeOwnerData) SetSymbol(symbol string) *ChangeOwnerData {
	copy(d.Symbol[:], symbol)
	return d
}

func (d *ChangeOwnerData) SetNewOwner(address string) (*ChangeOwnerData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.NewOwner[:], bytes)
	return d, nil
}

// Tries to set address of candidate and panics on error.
func (d *ChangeOwnerData) MustSetNewOwner(address string) *ChangeOwnerData {
	_, err := d.SetNewOwner(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *ChangeOwnerData) Type() Type {
	return TypeChangeOwner
}

func (d *ChangeOwnerData) Fee() Fee {
	return feeTypeChangeOwner
}

func (d *ChangeOwnerData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
