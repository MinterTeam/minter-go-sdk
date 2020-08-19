package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

type ChangeCoinOwnerData struct {
	Symbol   CoinSymbol
	NewOwner [20]byte
}

func NewChangeCoinOwnerData() *ChangeCoinOwnerData {
	return &ChangeCoinOwnerData{}
}

func (d *ChangeCoinOwnerData) SetSymbol(symbol string) *ChangeCoinOwnerData {
	copy(d.Symbol[:], symbol)
	return d
}

func (d *ChangeCoinOwnerData) SetNewOwner(address string) (*ChangeCoinOwnerData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.NewOwner[:], bytes)
	return d, nil
}

// Tries to set address of candidate and panics on error.
func (d *ChangeCoinOwnerData) MustSetNewOwner(address string) *ChangeCoinOwnerData {
	_, err := d.SetNewOwner(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *ChangeCoinOwnerData) Type() Type {
	return TypeChangeCoinOwner
}

func (d *ChangeCoinOwnerData) Fee() Fee {
	return feeTypeChangeCoinOwner
}

func (d *ChangeCoinOwnerData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
