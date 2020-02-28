package wallet

import (
	"encoding/hex"
)

type Wallet struct {
	*Data
}

type Data struct {
	Mnemonic   string
	Seed       string
	PrivateKey string
	PublicKey  string
	Address    string
}

// Create wallet. This method returns generated seed, private key, public key, mnemonic and Minter address.
func Create() (*Data, error) {
	mnemonic, err := NewMnemonic()
	if err != nil {
		return nil, err
	}

	seed, err := Seed(mnemonic)
	if err != nil {
		return nil, err
	}
	data, err := NewWallet(seed)
	if err != nil {
		return nil, err
	}

	data.Mnemonic = mnemonic

	return data.Data, nil
}

func NewWallet(seed []byte) (*Wallet, error) {
	prKey, err := PrivateKeyBySeed(seed)
	if err != nil {
		return nil, err
	}

	pubKey, err := PublicKeyByPrivateKey(prKey)
	if err != nil {
		return nil, err
	}

	address, err := AddressByPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		Data: &Data{
			Mnemonic:   "",
			Seed:       hex.EncodeToString(seed),
			PrivateKey: prKey,
			PublicKey:  pubKey,
			Address:    address,
		},
	}, nil
}

func (w *Wallet) Address() string {
	return w.Data.Address
}

func (w *Wallet) PrivateKey() string {
	return w.Data.PrivateKey
}

func (w *Wallet) PublicKey() string {
	return w.Data.PublicKey
}
