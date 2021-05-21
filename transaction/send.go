package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"math/big"
)

// SendData is a Data of Transaction for sending arbitrary coin.
type SendData struct {
	Coin  CoinID   // ID of a coin
	To    Address  // Recipient address
	Value *big.Int // Amount of coin to send
}

// NewSendData returns new SendData of Transaction for sending arbitrary coin.
func NewSendData() *SendData {
	return &SendData{}
}

// SetCoin sets ID of a coin.
func (d *SendData) SetCoin(id uint64) *SendData {
	d.Coin = CoinID(id)
	return d
}

// SetTo sets recipient address.
func (d *SendData) SetTo(address string) (*SendData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.To[:], bytes)
	return d, nil
}

// MustSetTo tries to set recipient address and panics on error
func (d *SendData) MustSetTo(address string) *SendData {
	_, err := d.SetTo(address)
	if err != nil {
		panic(err)
	}
	return d
}

// SetValue sets amount of coin to send.
func (d *SendData) SetValue(value *big.Int) *SendData {
	d.Value = value
	return d
}

// Type returns Data type of the transaction.
func (d *SendData) Type() Type {
	return TypeSend
}

// Encode returns the byte representation of a transaction Data.
func (d *SendData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
