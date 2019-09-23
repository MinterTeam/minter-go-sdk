package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
	"strings"
)

// Generate mnemonic.
func NewMnemonic() (string, error) {
	entropy := make([]byte, 16)
	_, err := rand.Read(entropy)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

// Get seed from mnemonic.
func Seed(mnemonic string) ([]byte, error) {
	return bip39.NewSeedWithErrorChecking(mnemonic, "")
}

// Get private key from seed.
func PrivateKeyBySeed(seed []byte) (string, error) {
	wallet, err := NewWallet(seed)
	if err != nil {
		return "", err
	}
	return wallet.PrivateKey(), nil
}

// Get Minter address from public key.
func AddressByPublicKey(publicKey string) (string, error) {
	bytes, err := hex.DecodeString(publicKey[2:])
	if err != nil {
		return "", err
	}
	return addressToLowerPrefix0xToMx(common.BytesToAddress(crypto.Keccak256(bytes)[12:]).String()), nil
}

// Get public key from private key.
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
