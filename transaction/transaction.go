package transaction

import (
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

type Data interface {
	encode() ([]byte, error)
}

type Interface interface {
	SetNonce(nonce uint64)
	SetChainID(chain ChainID)
	SetGasCoin(name string)
	SetGasPrice(price *big.Int)
	SetData(data Data)
	Sign(privateKey []byte) error
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
	SignatureData *Signature
}

type Signature struct {
	V *big.Int
	R *big.Int
	S *big.Int
}

func (tx *Transaction) SetNonce(nonce uint64) {
	tx.Nonce = nonce
}

func (tx *Transaction) SetChainID(chain ChainID) {
	tx.ChainID = chain
}

func (tx *Transaction) SetGasCoin(name string) {
	copy(tx.GasCoin[:], name)
}

func (tx *Transaction) SetGasPrice(price uint8) {
	tx.GasPrice = price
}

func (tx *Transaction) Sign(privateKey []byte) error {
	hw := sha3.NewLegacyKeccak256()
	err := rlp.Encode(hw, tx) //todo не все поля?
	if err != nil {
		panic(err)
	}
	hw.Sum(nil)
	h := make([]byte, 32)
	hw.Write(h)
	sig, err := secp256k1.Sign(h, privateKey)
	if err != nil {
		return err
	}

	tx.SignatureData = &Signature{
		R: new(big.Int).SetBytes(sig[:32]),
		S: new(big.Int).SetBytes(sig[32:64]),
		V: new(big.Int).SetBytes([]byte{sig[64] + 27}),
	}

	tx.Payload, err = rlp.EncodeToBytes(tx.SignatureData)
	if err != nil {
		return err
	}

	return nil
}
