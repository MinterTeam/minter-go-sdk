package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/MinterTeam/minter-go-sdk/v2/internal/common"
	"github.com/MinterTeam/minter-go-sdk/v2/internal/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// NewMnemonic generates mnemonic.
func NewMnemonic() (string, error) {
	entropy := make([]byte, 16)
	_, err := rand.Read(entropy)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

// Seed returns seed from mnemonic.
func Seed(mnemonic string) (string, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(seed), nil
}

// PrivateKeyBySeed returns private key from seed.
func PrivateKeyBySeed(seed string) (string, error) {
	bytes, err := hex.DecodeString(seed)
	if err != nil {
		return "", err
	}

	// Generate a new master node using the seed.
	masterKey, err := bip32.NewMasterKey(bytes)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H
	acc44H, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H
	acc44H60H, err := acc44H.NewChildKey(bip32.FirstHardenedChild + 60)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H/0H
	acc44H60H0H, err := acc44H60H.NewChildKey(bip32.FirstHardenedChild + 0)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H/0H/0
	acc44H60H0H0, err := acc44H60H0H.NewChildKey(0)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H/0H/0/0
	acc44H60H0H00, err := acc44H60H0H0.NewChildKey(0)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(acc44H60H0H00.Key), nil
}

// AddressByPublicKey returns Minter address from public key.
func AddressByPublicKey(publicKey string) (string, error) {
	publicKey = strings.Title(strings.ToLower(publicKey))
	if !strings.HasPrefix(publicKey, "Mp") {
		return "", errors.New("public key don't has prefix 'Mp'")
	}

	decodeString, err := hex.DecodeString(publicKey[2:])
	if err != nil {
		return "", err
	}

	return addressToLowerPrefix0xToMx(common.BytesToAddress(crypto.Keccak256(decodeString)[12:]).String()), nil
}

// BytesToAddress converts hex value to [20]byte of address
func BytesToAddress(address [20]byte) string {
	return addressToLowerPrefix0xToMx(common.BytesToAddress(address[:]).String())
}

// PublicKeyByPrivateKey returns public key from private key.
func PublicKeyByPrivateKey(privateKey string) (string, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	return PubPrefix04ToMp(hex.EncodeToString(crypto.FromECDSAPub(key.Public().(*ecdsa.PublicKey)))), nil
}

func addressToLowerPrefix0xToMx(key string) string {
	return strings.Replace(strings.ToLower(key), "0x", "Mx", 1)
}

// PubPrefix04ToMp replace "04" to "Mp"
func PubPrefix04ToMp(key string) string {
	return strings.Replace(key, "04", "Mp", 1)
}

// AddressToHex converts string address to hex bytes
func AddressToHex(address string) ([]byte, error) {
	address = strings.Title(strings.ToLower(address))
	if !strings.HasPrefix(address, "Mx") {
		return nil, errors.New("address don't has prefix 'Mx'")
	}

	if len(address) != 42 {
		return nil, errors.New("address length less than 42 characters")
	}

	bytes, err := hex.DecodeString(address[2:])
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PublicKeyToHex converts string public key to hex bytes
func PublicKeyToHex(key string) ([]byte, error) {
	key = strings.Title(strings.ToLower(key))
	if !strings.HasPrefix(key, "Mp") {
		return nil, errors.New("public key don't has prefix 'Mp'")
	}

	hexKey, err := hex.DecodeString(key[2:])
	if err != nil {
		return hexKey, err
	}
	return hexKey, nil
}

// IsValidAddress checks address for correctness
func IsValidAddress(address string) bool {
	_, err := AddressToHex(address)
	return err == nil
}
