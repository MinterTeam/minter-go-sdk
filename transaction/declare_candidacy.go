package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type DeclareCandidacyData struct {
	Address    [20]byte
	PubKey     []byte
	Commission uint
	Coin       [10]byte
	Stake      *big.Int
}

func NewDeclareCandidacyData() *DeclareCandidacyData {
	return &DeclareCandidacyData{}
}

func (d *DeclareCandidacyData) SetAddress(address string) (*DeclareCandidacyData, error) {
	bytes, err := AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.Address[:], bytes)
	return d, nil
}

func (d *DeclareCandidacyData) MustSetAddress(address string) *DeclareCandidacyData {
	_, err := d.SetAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *DeclareCandidacyData) SetPubKey(key string) *DeclareCandidacyData {
	d.PubKey = []byte(key)
	return d
}

func (d *DeclareCandidacyData) SetCommission(value uint) *DeclareCandidacyData {
	d.Commission = value
	return d
}

func (d *DeclareCandidacyData) SetCoin(symbol string) *DeclareCandidacyData {
	copy(d.Coin[:], symbol)
	return d
}

func (d *DeclareCandidacyData) SetStake(value *big.Int) *DeclareCandidacyData {
	d.Stake = big.NewInt(0).Mul(value, expPip)
	return d
}

func (d *DeclareCandidacyData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
