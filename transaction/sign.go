package transaction

import (
	//"github.com/edwardstock/bip39go"
	//"github.com/tyler-smith/go-bip39"
	"strconv"
	"strings"
)

//
//type TransactionInterface interface {
//	Sign(privateKey []byte) (SignedTransaction, error)
//}

func MnemonicToSeed(mnemonic string) []byte {
	return nil
	//return bip39.NewSeed(mnemonic, "")
	//return bip39go.WordsToSeed(mnemonic,,)
}

func SeedToPrivateKey(seed []byte) []byte {
	return nil
}

func Mnemonic() []byte {
	return nil
}

func AddressFromPublicKey(publicKey []byte) string {
	return ""
}

func PrivateToPublicKey(publicKey []byte) string {
	return ""
}

func ValidateAddress(address string) bool {
	if len(address) != 42 {
		return false
	}
	if !strings.HasPrefix(address, "Mx") {
		return false
	}
	if _, err := strconv.ParseUint(address[:2], 16, 64); err != nil {
		return false
	}
	return true
}
