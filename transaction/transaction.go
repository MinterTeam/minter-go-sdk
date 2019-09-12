package transaction

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"math/big"
)

type Type byte

const (
	typeSend Type = iota
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
	_ ChainID = iota
	MainNetChainID
	TestNetChainID
)

type Builder struct {
	ChainID ChainID
}

func NewBuilder(chainID ChainID) *Builder {
	return &Builder{ChainID: chainID}
}

func (b *Builder) NewTransaction(data Data) (Interface, error) {
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

	default:
		return nil, errors.New("") //todo
	}
}

type Data interface {
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

func (tx *Transaction) Sign(privateKey []byte) (SignedTransaction, error) {
	hw := sha3.NewLegacyKeccak256()
	err := rlp.Encode(hw, []interface{}{
		tx.Nonce,
		tx.ChainID,
		tx.GasPrice,
		tx.GasCoin,
		tx.Type,
		tx.Data,
		tx.Payload,
		tx.ServiceData,
		tx.SignatureType,
	})
	if err != nil {
		return nil, err
	}

	hw.Sum(nil)
	h := make([]byte, 32)
	hw.Write(h)

	sig, err := secp256k1.Sign(h, privateKey)
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
