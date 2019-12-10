package transaction

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"strconv"
)

// Issue a check that will later be redeemed by the person of your choice.
type IssueCheckData struct {
	Nonce    []byte
	ChainID  ChainID
	DueBlock uint64
	Coin     [10]byte
	Value    *big.Int
	GasCoin  [10]byte
	Lock     *big.Int
	V        *big.Int
	R        *big.Int
	S        *big.Int
}

type Signed interface {
	Encode() (string, error)
}

type IssueCheckInterface interface {
	SetPassphrase(passphrase string) IssueCheckInterface
	Sign(prKey string) (Signed, error)
}

type IssueCheck struct {
	*IssueCheckData
	passphrase string
}

// Create Issue Check
// Nonce - unique "id" of the check. Coin Symbol - symbol of coin. Value - amount of coins.
// Due Block - defines last block height in which the check can be used.
func NewIssueCheck(nonce uint64, chainID ChainID, dueBlock uint64, coin string, value *big.Int, gasCoin string) IssueCheckInterface {
	check := &IssueCheck{
		IssueCheckData: &IssueCheckData{
			Nonce:    []byte(strconv.Itoa(int(nonce))),
			ChainID:  chainID,
			DueBlock: dueBlock,
			Value:    value,
		},
	}
	copy(check.Coin[:], coin)
	copy(check.GasCoin[:], gasCoin)
	return check
}

// Prepare check string and convert to data
func DecodeIssueCheck(check string) (*IssueCheckData, error) {
	src := []byte(check)[2:]
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}

	res := new(IssueCheckData)
	err = rlp.DecodeBytes(dst, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Set secret phrase which you will pass to receiver of the check
func (check *IssueCheck) SetPassphrase(passphrase string) IssueCheckInterface {
	check.passphrase = passphrase
	return check
}

//
func (check *IssueCheck) Encode() (string, error) {
	src, err := rlp.EncodeToBytes(check.IssueCheckData)
	if err != nil {
		return "", err
	}

	return "Mc" + hex.EncodeToString(src), err
}

// Sign Issue Check
func (check *IssueCheck) Sign(prKey string) (Signed, error) {
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
	bytes, err := addressToHex(address)
	if err != nil {
		return nil, err
	}

	check := &CheckAddress{passphrase: passphrase}
	copy(check.address[:], bytes)

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
