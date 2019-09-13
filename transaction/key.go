package transaction

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
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

func ToECDSA(d []byte) (*ecdsa.PrivateKey, error) {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = secp256k1.S256()
	if 8*len(d) != priv.Params().BitSize {
		return nil, fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
	}
	priv.D = new(big.Int).SetBytes(d)

	// The priv.D must < N
	secp256k1N, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	if priv.D.Cmp(secp256k1N) >= 0 {
		return nil, fmt.Errorf("invalid private key, >=N")
	}
	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return nil, fmt.Errorf("invalid private key, zero or negative")
	}

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}
	return priv, nil
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
