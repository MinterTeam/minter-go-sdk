package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

type EditCoinOwnerData struct {
	Symbol   CoinSymbol
	NewOwner [20]byte
}

func NewEditCoinOwnerData() *EditCoinOwnerData {
	return &EditCoinOwnerData{}
}

func (d *EditCoinOwnerData) SetSymbol(symbol string) *EditCoinOwnerData {
	copy(d.Symbol[:], symbol)
	return d
}

func (d *EditCoinOwnerData) SetNewOwner(address string) (*EditCoinOwnerData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.NewOwner[:], bytes)
	return d, nil
}

// Tries to set address of candidate and panics on error.
func (d *EditCoinOwnerData) MustSetNewOwner(address string) *EditCoinOwnerData {
	_, err := d.SetNewOwner(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *EditCoinOwnerData) Type() Type {
	return TypeEditCoinOwner
}

func (d *EditCoinOwnerData) Fee() Fee {
	return feeTypeEditCoinOwner
}

func (d *EditCoinOwnerData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
