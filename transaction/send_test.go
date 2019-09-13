package transaction

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
	"strings"
	"testing"
)

func TestTransaction_Sign(t *testing.T) {
	value := big.NewInt(1)
	pipValue := value.String() + strings.Repeat("0", 18)
	address := "Mx1b685a7c1e78726c48f619c497a07ed75fe00483"
	symbolMNT := "MNT"
	data := NewSendData().
		SetCoin(symbolMNT).
		SetValue(value).
		MustSetTo(address)

	if string(data.Coin[:3]) != symbolMNT {
		t.Errorf("SendData.Coin got %s, want %s", data.Coin, symbolMNT)
	}

	addressBytes, err := hex.DecodeString(address[2:])
	if string(data.To[:]) != string(addressBytes) {
		t.Errorf("SendData.To got %s, want %s", string(data.To[:]), string(addressBytes))
	}

	if data.Value.String() != pipValue {
		t.Errorf("SendData.Value got %s, want %s", data.Value.String(), pipValue)
	}
	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(symbolMNT)
	transaction := tx.(*Transaction)

	if transaction.Nonce != nonce {
		t.Errorf("Nonce got %d, want %d", transaction.Nonce, nonce)
	}

	if transaction.ChainID != TestNetChainID {
		t.Errorf("ChainID got %d, want %d", transaction.ChainID, TestNetChainID)
	}

	if transaction.GasPrice != gasPrice {
		t.Errorf("GasPrice got %d, want %d", transaction.GasPrice, gasPrice)
	}

	if string(transaction.GasCoin[:3]) != symbolMNT {
		t.Errorf("GasCoin got %s, want %s", transaction.GasCoin, symbolMNT)
	}

	privateKey, err := hex.DecodeString("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	key, err := toECDSA(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	_, err = transaction.Sign(key)
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf8840102018a4d4e540000000000000001aae98a4d4e5400000000000000941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808001b845f8431ca01f36e51600baa1d89d2bee64def9ac5d88c518cdefe45e3de66a3cf9fe410de4a01bc2228dc419a97ded0efe6848de906fbe6c659092167ef0e7dcb8d15024123a"
	if string(transaction.SignatureData) == validSignature {
		t.Errorf("encode get %s, want %s", string(transaction.SignatureData), validSignature)
	}
}

func toECDSA(d []byte) (*ecdsa.PrivateKey, error) {
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
