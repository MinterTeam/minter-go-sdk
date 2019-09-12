package transaction

import (
	"encoding/hex"
	"errors"
	"strings"
)

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

func AddressToHex(address string) ([]byte, error) {
	if len(address) != 42 {
		return nil, errors.New("len < 42") //todo
	}
	if !strings.HasPrefix(address, "Mx") {
		return nil, errors.New("don't has prefix 'Mx'") //todo
	}
	bytes, err := hex.DecodeString(address[2:])
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
