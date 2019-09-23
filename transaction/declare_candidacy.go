package transaction

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction for declaring new validator candidacy.
// Address - Address of candidate in Minter Network. This address would be able to control candidate.
// Also all rewards will be sent to this address. PubKey - Public key of a validator.
// Commission - Commission (from 0 to 100) from rewards which delegators will pay to validator.
// Coin - Symbol of coin to stake. Stake - Amount of coins to stake.
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

func (d *DeclareCandidacyData) SetPubKey(key string) (*DeclareCandidacyData, error) {
	var err error
	d.PubKey, err = hex.DecodeString(key[2:])
	if err != nil {
		return d, err
	}
	return d, nil
}

func (d *DeclareCandidacyData) MustSetPubKey(key string) *DeclareCandidacyData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
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
	d.Stake = value
	return d
}

func (d *DeclareCandidacyData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (d *DeclareCandidacyData) fee() Fee {
	return feeTypeDeclareCandidacy
}
