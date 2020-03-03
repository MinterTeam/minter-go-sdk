package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"strconv"
)

// Issue a check that will later be redeemed by the person of your choice.
type CheckData struct {
	Nonce    []byte
	ChainID  ChainID
	DueBlock uint64
	Coin     [10]byte
	Value    *big.Int
	Lock     *big.Int
	V        *big.Int
	R        *big.Int
	S        *big.Int
}

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

	return fmt.Sprintf("Check sender: %s nonce: %x, dueBlock: %d, value: %s %s", sender, check.Nonce,
		check.DueBlock, check.Value.String(), string(bytes.Trim(check.Coin[:], "\x00")))
}

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

type Signed interface {
	Encode() (string, error)
}

type CheckInterface interface {
	SetPassphrase(passphrase string) CheckInterface
	Sign(prKey string) (Signed, error)
}

type Check struct {
	*CheckData
	passphrase string
}

// Create Check
// Nonce - unique "id" of the check. Coin Symbol - symbol of coin. Value - amount of coins.
// Due Block - defines last block height in which the check can be used.
func NewCheck(nonce uint64, chainID ChainID, dueBlock uint64, coin string, value *big.Int) CheckInterface {
	check := &Check{
		CheckData: &CheckData{
			Nonce:    []byte(strconv.Itoa(int(nonce))),
			ChainID:  chainID,
			DueBlock: dueBlock,
			Value:    value,
		},
	}
	copy(check.Coin[:], coin)
	return check
}

// Prepare check string and convert to data
func DecodeCheck(rawCheck string) (*CheckData, error) {
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

// Set secret phrase which you will pass to receiver of the check
func (check *Check) SetPassphrase(passphrase string) CheckInterface {
	check.passphrase = passphrase
	return check
}

//
func (check *Check) Encode() (string, error) {
	src, err := rlp.EncodeToBytes(check.CheckData)
	if err != nil {
		return "", err
	}

	return "Mc" + hex.EncodeToString(src), err
}

// Sign Check
func (check *Check) Sign(prKey string) (Signed, error) {
	msgHash, err := rlpHash([]interface{}{
		check.Nonce,
		check.ChainID,
		check.DueBlock,
		check.Coin,
		check.Value,
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
	toHex, err := addressToHex(address)
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
