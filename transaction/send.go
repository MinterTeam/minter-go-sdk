package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction data for sending arbitrary coin.
type SendData struct {
	Coin  CoinID   // ID of a coin
	To    [20]byte // Recipient address
	Value *big.Int // Amount of coin to send
}

// New data of transaction data for sending arbitrary coin.
func NewSendData() *SendData {
	return &SendData{}
}

// Set ID of a coin.
func (d *SendData) SetCoin(id CoinID) *SendData {
	d.Coin = id
	return d
}

// Set recipient address.
func (d *SendData) SetTo(address string) (*SendData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.To[:], bytes)
	return d, nil
}

// Tries to set recipient address and panics on error
func (d *SendData) MustSetTo(address string) *SendData {
	_, err := d.SetTo(address)
	if err != nil {
		panic(err)
	}
	return d
}

// Set amount of coin to send.
func (d *SendData) SetValue(value *big.Int) *SendData {
	d.Value = value
	return d
}

func (d *SendData) Type() Type {
	return TypeSend
}

func (d *SendData) Fee() Fee {
	return feeTypeSend
}

func (d *SendData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
