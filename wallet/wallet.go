package wallet

import (
	"errors"
)

// Wallet of minter
type Wallet struct {
	Mnemonic   string
	Seed       string
	PrivateKey string
	PublicKey  string
	Address    string
}

// New return new wallet. This method returns generated seed, private key, public key, mnemonic and Minter address.
func New() (*Wallet, error) {
	mnemonic, err := NewMnemonic()
	if err != nil {
		return nil, err
	}

	data, err := Create(mnemonic, "")
	if err != nil {
		return nil, err
	}

	data.Mnemonic = mnemonic

	return data, nil
}

// Create returns wallet by exists seed or mnemonic.
func Create(mnemonic string, seed string) (*Wallet, error) {
	if mnemonic != "" {
		seed1, err := Seed(mnemonic)
		if err != nil {
			return nil, err
		}

		if len(seed) != 0 && seed1 != seed {
			return nil, errors.New("seed and mnemonic are not empty, but not equal")
		}

		seed = seed1
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

	return &Wallet{
		Mnemonic:   mnemonic,
		Seed:       seed,
		PrivateKey: prKey,
		PublicKey:  pubKey,
		Address:    address,
	}, nil
}
