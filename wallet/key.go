package wallet

import (
	"crypto/rand"
	"github.com/foxnut/go-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"strings"
)

func NewMnemonic() (string, error) {
	entropy := make([]byte, 16)
	_, err := rand.Read(entropy)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

func Seed(mnemonic string) ([]byte, error) {
	return bip39.NewSeedWithErrorChecking(mnemonic, "")
}

type Wallet struct {
	wallet hdwallet.Wallet
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

func (w *Wallet) Address() (string, error) {
	address, err := w.wallet.GetAddress()
	if err != nil {
		return "", err
	}
	return strings.Replace(strings.ToLower(address), "0x", "Mx", 1), nil
}

func (w *Wallet) PrivateKey() string {
	return w.wallet.GetKey().PrivateHex()
}

func (w *Wallet) PublicKey() string {
	return w.wallet.GetKey().PublicHex(false)
}
