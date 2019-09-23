package wallet

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/foxnut/go-hdwallet"
)

type Wallet struct {
	wallet hdwallet.Wallet
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
	wallet, err := NewWallet(seed)
	if err != nil {
		return nil, err
	}
	return &Data{
		Mnemonic:   mnemonic,
		Seed:       hex.EncodeToString(seed),
		PrivateKey: wallet.PrivateKey(),
		PublicKey:  wallet.PublicKey(),
		Address:    wallet.Address(),
	}, nil
}

func NewWallet(seed []byte) (*Wallet, error) {
	masterKey, err := hdwallet.NewKey(hdwallet.Seed(seed))
	if err != nil {
		return nil, err
	}
	wallet, err := masterKey.GetWallet(hdwallet.CoinType(hdwallet.ETH))
	if err != nil {
		return nil, err
	}
	return &Wallet{wallet: wallet}, nil
}

func (w *Wallet) Address() string {
	return addressToLowerPrefix0xToMx(crypto.PubkeyToAddress(*w.wallet.GetKey().PublicECDSA).Hex())
}

func (w *Wallet) PrivateKey() string {
	return w.wallet.GetKey().PrivateHex()
}

func (w *Wallet) PublicKey() string {
	return pubPrefix04ToMp(w.wallet.GetKey().PublicHex(false))
}
