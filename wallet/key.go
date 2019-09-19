package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

func (w *Wallet) Address() string {
	return addressToLowerPrefix0xToMx(crypto.PubkeyToAddress(*w.wallet.GetKey().PublicECDSA).Hex())
}

func (w *Wallet) PrivateKey() string {
	return w.wallet.GetKey().PrivateHex()
}

func (w *Wallet) PublicKey() string {
	return pubPrefix04ToMp(w.wallet.GetKey().PublicHex(false))
}

func AddressByPublicKey(publicKey string) (string, error) {
	bytes, err := hex.DecodeString(publicKey[2:])
	if err != nil {
		return "", err
	}
	return addressToLowerPrefix0xToMx(common.BytesToAddress(crypto.Keccak256(bytes)[12:]).String()), nil
}

func PublicKeyByPrivateKey(privateKey string) (string, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	return pubPrefix04ToMp(hex.EncodeToString(crypto.FromECDSAPub(key.Public().(*ecdsa.PublicKey)))), nil
}

func addressToLowerPrefix0xToMx(key string) string {
	return strings.Replace(strings.ToLower(key), "0x", "Mx", 1)
}

func pubPrefix04ToMp(key string) string {
	return strings.Replace(key, "04", "Mp", 1)
}
