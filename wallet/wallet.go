package wallet

import (
	"encoding/hex"
	"errors"
)

// Wallet of minter
type Wallet struct {
	*Data
}

// Data of wallet
type Data struct {
	Mnemonic   string
	Seed       string
	PrivateKey string
	PublicKey  string
	Address    string
}

// New return new wallet. This method returns generated seed, private key, public key, mnemonic and Minter address.
func New() (*Data, error) {
	mnemonic, err := NewMnemonic()
	if err != nil {
		return nil, err
	}

	data, err := Create(mnemonic, nil)
	if err != nil {
		return nil, err
	}

	data.Mnemonic = mnemonic

	return data, nil
}

// Create returns wallet by exists seed or mnemonic. Note: pass only one value.
func Create(mnemonic string, seed []byte) (*Data, error) {
	if mnemonic != "" {
		if len(seed) != 0 {
			return nil, errors.New("pass only one value (seed or mnemonic")
		}
		var err error
		seed, err = Seed(mnemonic)
		if err != nil {
			return nil, err
		}
	}

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

	return &Data{
		Mnemonic:   mnemonic,
		Seed:       hex.EncodeToString(seed),
		PrivateKey: prKey,
		PublicKey:  pubKey,
		Address:    address,
	}, nil
}
