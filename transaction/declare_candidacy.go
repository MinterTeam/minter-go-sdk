package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

// Transaction data for declaring new validator candidacy.
type DeclareCandidacyData struct {
	Address    [20]byte // Address of candidate
	PubKey     [32]byte // Public key of a validator
	Commission uint     // Commission (from 0 to 100) from rewards which delegators will pay to validator
	Coin       CoinID   // ID of coin to stake
	Stake      *big.Int // Amount of coins to stake
}

// New data of transaction for declaring new validator candidacy.
func NewDeclareCandidacyData() *DeclareCandidacyData {
	return &DeclareCandidacyData{}
}

// Set address of candidate. This address would be able to control candidate. Also all rewards will be sent to this address
func (d *DeclareCandidacyData) SetAddress(address string) (*DeclareCandidacyData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.Address[:], bytes)
	return d, nil
}

// Tries to set address of candidate and panics on error.
func (d *DeclareCandidacyData) MustSetAddress(address string) *DeclareCandidacyData {
	_, err := d.SetAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

// Set public key of a validator.
func (d *DeclareCandidacyData) SetPubKey(key string) (*DeclareCandidacyData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// Tries to set public key of validator and panics on error.
func (d *DeclareCandidacyData) MustSetPubKey(key string) *DeclareCandidacyData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Set commission (from 0 to 100) from rewards which delegators will pay to validator.
func (d *DeclareCandidacyData) SetCommission(value uint) *DeclareCandidacyData {
	d.Commission = value
	return d
}

// Set ID of coin to stake.
func (d *DeclareCandidacyData) SetCoin(id CoinID) *DeclareCandidacyData {
	d.Coin = id
	return d
}

// Set amount of coins to stake.
func (d *DeclareCandidacyData) SetStake(value *big.Int) *DeclareCandidacyData {
	d.Stake = value
	return d
}

func (d *DeclareCandidacyData) Type() Type {
	return TypeDeclareCandidacy
}

func (d *DeclareCandidacyData) Fee() Fee {
	return feeTypeDeclareCandidacy
}

func (d *DeclareCandidacyData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
