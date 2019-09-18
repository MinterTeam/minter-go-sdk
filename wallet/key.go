package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
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
	return strings.Replace(strings.ToLower(crypto.PubkeyToAddress(*w.wallet.GetKey().PublicECDSA).Hex()), "0x", "Mx", 1)
}

func (w *Wallet) PrivateKey() string {
	return w.wallet.GetKey().PrivateHex()
}

func (w *Wallet) PublicKey() string {
	return w.wallet.GetKey().PublicHex(true)
}

func AddressByPublicKey(publicKey string) (string, error) {
	decodeString, err := hex.DecodeString(publicKey)
	if err != nil {
		return "", err
	}
	key, err := btcec.ParsePubKey(decodeString, btcec.S256())
	if err != nil {
		return "", err
	}
	return strings.Replace(strings.ToLower(crypto.PubkeyToAddress(*key.ToECDSA()).Hex()), "0x", "Mx", 1), nil
}
func PublicKeyByPrivateKey(privateKey string) (string, error) {
	key, err := ethcrypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	return strings.Replace(strings.ToLower(hex.EncodeToString(ethcrypto.FromECDSAPub(key.Public().(*ecdsa.PublicKey)))), "04", "Mp", 1), nil
}
