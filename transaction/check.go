package transaction

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"strings"
)

// Minter Check is like an ordinary bank check.
// Each user of network can issue check with any amount of coins and pass it to another person.
// Receiver will be able to cash a check from arbitrary account.
type CheckData struct {
	Nonce    []byte   // Unique ID of the check
	ChainID  ChainID  // ID of the network
	DueBlock uint64   // Defines last block height in which the check can be used
	Coin     CoinID   // ID of coin
	Value    *big.Int // Amount of coins
	GasCoin  CoinID   // ID of a coin to pay fee
	Lock     *big.Int // Secret to prevent hijacking
	V        *big.Int // V signature of issuer
	R        *big.Int // R signature of issuer
	S        *big.Int // S signature of issuer
}

// Tries returns sender of transaction and panics on error.
func (check *CheckData) MustSender() string {
	sender, err := check.Sender()
	if err != nil {
		panic(err)
	}

	return sender
}

// Return sender of transactions.
func (check *CheckData) Sender() (string, error) {
	pub, err := check.PublicKey()
	if err != nil {
		return "", err
	}

	return wallet.AddressByPublicKey(pub)
}

func (check *CheckData) String() string {
	sender, err := check.Sender()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Check sender: %s nonce: %x, dueBlock: %d, value: %s, coinID: %s", sender, check.Nonce,
		check.DueBlock, check.Value.String(), check.Coin)
}

// Returns public key from the transaction signature.
func (check *CheckData) PublicKey() (string, error) {

	if check.V.BitLen() > 8 {
		return "", errors.New("invalid transaction v, r, s values")
	}

	v := byte(check.V.Uint64() - 27)
	if !crypto.ValidateSignatureValues(v, check.R, check.S, true) {
		return "", errors.New("invalid transaction v, r, s values")
	}

	r := check.R.Bytes()
	s := check.S.Bytes()

	sig := make([]byte, 65)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = v

	hash, err := rlpHash([]interface{}{
		check.Nonce,
		check.ChainID,
		check.DueBlock,
		check.Coin,
		check.Value,
		check.GasCoin,
		check.Lock,
	})
	if err != nil {
		return "", err
	}

	pub, err := secp256k1.RecoverPubkey(hash[:], sig)
	if err != nil {
		return "", err
	}

	if len(pub) == 0 || pub[0] != 4 {
		return "", errors.New("invalid public key")
	}

	return wallet.PubPrefix04ToMp(hex.EncodeToString(pub)), nil
}

type CheckInterface interface {
	SetPassphrase(passphrase string) CheckInterface
	Sign(prKey string) (EncodeInterface, error)
}

type Check struct {
	*CheckData
	passphrase string
}

// Issue a check that will later be redeemed by the person of your choice.
func NewCheck(nonce string, chainID ChainID, dueBlock uint64, coin CoinID, value *big.Int, gasCoin CoinID) CheckInterface {
	check := &Check{
		CheckData: &CheckData{
			Nonce:    []byte(nonce),
			ChainID:  chainID,
			DueBlock: dueBlock,
			Coin:     coin,
			Value:    value,
			GasCoin:  gasCoin,
		},
	}
	return check
}

// Get CheckData from RLP-encoded structure in base64 format.
func DecodeCheckBase64(rawCheck string) (*CheckData, error) {
	decode, err := base64.StdEncoding.DecodeString(rawCheck)
	if err != nil {
		panic(err)
	}

	res := new(CheckData)
	if err := rlp.DecodeBytes(decode, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Get CheckData from RLP-encoded structure in hex format.
func DecodeCheck(check string) (*CheckData, error) {
	check = strings.Title(strings.ToLower(check))
	if !strings.HasPrefix(check, "Mc") {
		return nil, errors.New("check don't has prefix 'Mc'")
	}

	decode, err := hex.DecodeString(check[2:])
	if err != nil {
		panic(err)
	}

	res := new(CheckData)
	if err := rlp.DecodeBytes(decode, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Set secret phrase which you will pass to receiver of the check
func (check *Check) SetPassphrase(passphrase string) CheckInterface {
	check.passphrase = passphrase
	return check
}

// Get string representation of a check. Checks are prefixed with "Mc". RLP-encoded structure in hex format.
func (check *Check) Encode() (string, error) {
	src, err := rlp.EncodeToBytes(check.CheckData)
	if err != nil {
		return "", err
	}

	return "Mc" + hex.EncodeToString(src), nil
}

// Get string representation of a check. RLP-encoded structure in base64 format.
func (check *Check) EncodeBase64() (string, error) {
	src, err := rlp.EncodeToBytes(check.CheckData)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(src), nil
}

// Sign Check with private key
func (check *Check) Sign(prKey string) (EncodeInterface, error) {
	msgHash, err := rlpHash([]interface{}{
		check.Nonce,
		check.ChainID,
		check.DueBlock,
		check.Coin,
		check.Value,
		check.GasCoin,
	})
	if err != nil {
		return nil, err
	}

	passphraseSum256 := sha256.Sum256([]byte(check.passphrase))

	lock, err := secp256k1.Sign(msgHash[:], passphraseSum256[:])
	if err != nil {
		return nil, err
	}

	check.Lock = big.NewInt(0).SetBytes(lock)

	msgHashWithLock, err := rlpHash([]interface{}{
		check.Nonce,
		check.ChainID,
		check.DueBlock,
		check.Coin,
		check.Value,
		check.GasCoin,
		check.Lock,
	})
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(prKey)
	if err != nil {
		return nil, err
	}

	sig, err := crypto.Sign(msgHashWithLock[:], privateKey)
	if err != nil {
		return nil, err
	}

	check.R = new(big.Int).SetBytes(sig[:32])
	check.S = new(big.Int).SetBytes(sig[32:64])
	check.V = new(big.Int).SetBytes([]byte{sig[64] + 27})

	return check, nil
}

type CheckAddress struct {
	address    [20]byte
	passphrase string
}

func NewCheckAddress(address string, passphrase string) (*CheckAddress, error) {
	toHex, err := wallet.AddressToHex(address)
	if err != nil {
		return nil, err
	}

	check := &CheckAddress{passphrase: passphrase}
	copy(check.address[:], toHex)

	return check, nil
}

// Proof Check
func (check *CheckAddress) Proof() (string, error) {

	passphraseSum256 := sha256.Sum256([]byte(check.passphrase))

	addressHash, err := rlpHash([]interface{}{
		check.address[:],
	})
	if err != nil {
		return "", err
	}

	lock, err := secp256k1.Sign(addressHash[:], passphraseSum256[:])
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(lock), nil
}
