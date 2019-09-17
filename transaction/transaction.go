package transaction

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strings"
)

var (
	expPip = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
)

type Type byte

const (
	typeSend Type = iota + 1
	typeSellCoin
	typeSellAllCoin
	typeBuyCoin
	typeCreateCoin
	typeDeclareCandidacy
	typeDelegate
	typeUnbond
	typeRedeemCheck
	typeSetCandidateOnline
	typeSetCandidateOffline
	typeCreateMultisig
	typeMultisend
	typeEditCandidate
)

type Fee uint

const (
	feeTypeSend                Fee = 10
	feeTypeSellCoin            Fee = 100
	feeTypeSellAllCoin         Fee = 100
	feeTypeBuyCoin             Fee = 100
	feeTypeCreateCoin          Fee = 1000
	feeTypeDeclareCandidacy    Fee = 10000
	feeTypeDelegate            Fee = 200
	feeTypeUnbond              Fee = 200
	feeTypeRedeemCheck         Fee = 30
	feeTypeSetCandidateOnline  Fee = 100
	feeTypeSetCandidateOffline Fee = 100
	feeTypeCreateMultisig      Fee = 100
	//feeMultisend Fee =  10+(n-1)*5
	feeTypeEditCandidate Fee = 100000
)

type SignatureType byte

const (
	_ SignatureType = iota
	signatureTypeSingle
	//signatureTypeMulti
)

type ChainID byte

const (
	MainNetChainID ChainID = iota + 1
	TestNetChainID
)

type Builder struct {
	ChainID ChainID
}

func NewBuilder(chainID ChainID) *Builder {
	return &Builder{ChainID: chainID}
}

func (b *Builder) NewTransaction(data DataInterface) (Interface, error) {
	dataBytes, err := data.encode()
	if err != nil {
		return nil, err
	}

	transaction := &Transaction{
		ChainID:       b.ChainID,
		SignatureType: signatureTypeSingle,
		Data:          dataBytes,
	}

	switch data.(type) {
	case *SendData:
		return transaction.setType(typeSend), nil
	case *SellCoinData:
		return transaction.setType(typeSellCoin), nil
	case *SellAllCoinData:
		return transaction.setType(typeSellAllCoin), nil
	case *BuyCoinData:
		return transaction.setType(typeBuyCoin), nil
	case *CreateCoinData:
		return transaction.setType(typeCreateCoin), nil
	case *DeclareCandidacyData:
		return transaction.setType(typeDeclareCandidacy), nil
	case *DelegateData:
		return transaction.setType(typeDelegate), nil
	case *UnbondData:
		return transaction.setType(typeUnbond), nil
	case *RedeemCheckData:
		return transaction.setType(typeRedeemCheck), nil
	case *SetCandidateOnData:
		return transaction.setType(typeSetCandidateOnline), nil
	case *SetCandidateOffData:
		return transaction.setType(typeSetCandidateOffline), nil

	case *EditCandidateData:
		return transaction.setType(typeEditCandidate), nil

	default:
		return nil, errors.New("") //todo
	}
}

type DataInterface interface {
	encode() ([]byte, error)
}

type SignedTransaction interface {
	Encode() ([]byte, error)
}

type Interface interface {
	setType(t Type) Interface
	SetNonce(nonce uint64) Interface
	SetGasCoin(name string) Interface
	SetGasPrice(price uint8) Interface
	Sign(prKey []byte) (SignedTransaction, error)
}

type Transaction struct {
	Nonce         uint64
	ChainID       ChainID
	GasPrice      uint8
	GasCoin       [10]byte
	Type          Type
	Data          []byte
	Payload       []byte
	ServiceData   []byte
	SignatureType SignatureType
	SignatureData []byte
}

type Signature struct {
	V *big.Int
	R *big.Int
	S *big.Int
}

func (tx *Transaction) setType(t Type) Interface {
	tx.Type = t
	return tx
}
func (tx *Transaction) SetNonce(nonce uint64) Interface {
	tx.Nonce = nonce
	return tx
}

func (tx *Transaction) SetGasCoin(name string) Interface {
	copy(tx.GasCoin[:], name)
	return tx
}

func (tx *Transaction) SetGasPrice(price uint8) Interface {
	tx.GasPrice = price
	return tx
}

func (tx *Transaction) Encode() ([]byte, error) {
	src, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, hex.EncodedLen(len(src))+2)
	dst[0], dst[1] = '0', 'x'
	hex.Encode(dst[2:], src)
	return dst, err
}

func (tx *Transaction) Sign(prKey []byte) (SignedTransaction, error) {
	privateKey, err := toECDSA(prKey)
	if err != nil {
		return nil, err
	}

	x := []interface{}{
		tx.Nonce,
		tx.ChainID,
		tx.GasPrice,
		tx.GasCoin,
		tx.Type,
		tx.Data,
		tx.Payload,
		tx.ServiceData,
		tx.SignatureType,
	}

	var h [32]byte
	hw := sha3.NewLegacyKeccak256()
	err = rlp.Encode(hw, x)
	if err != nil {
		return nil, err
	}
	hw.Sum(h[:0])

	seckey := math.PaddedBigBytes(privateKey.D, privateKey.Params().BitSize/8)

	sig, err := secp256k1.Sign(h[:], seckey)
	if err != nil {
		return nil, err
	}
	tx.SignatureData, err = rlp.EncodeToBytes(&Signature{
		R: new(big.Int).SetBytes(sig[:32]),
		S: new(big.Int).SetBytes(sig[32:64]),
		V: new(big.Int).SetBytes([]byte{sig[64] + 27}),
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func AddressToHex(address string) ([]byte, error) {
	if len(address) != 42 {
		return nil, errors.New("len < 42")
	}
	if !strings.HasPrefix(address, "Mx") {
		return nil, errors.New("don't has prefix 'Mx'")
	}
	bytes, err := hex.DecodeString(address[2:])
	if err != nil {
		return nil, err
	}
	return bytes, nil
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
