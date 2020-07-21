package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strconv"
	"strings"
)

type Type byte

const (
	_ Type = iota
	TypeSend
	TypeSellCoin
	TypeSellAllCoin
	TypeBuyCoin
	TypeCreateCoin
	TypeDeclareCandidacy
	TypeDelegate
	TypeUnbond
	TypeRedeemCheck
	TypeSetCandidateOnline
	TypeSetCandidateOffline
	TypeCreateMultisig
	TypeMultisend
	TypeEditCandidate
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
	// feeMultisend Fee =  10+(n-1)*5
	feeTypeEditCandidate Fee = 100000
)

type SignatureType byte

const (
	_ SignatureType = iota
	SignatureTypeSingle
	SignatureTypeMulti
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
		SignatureType: SignatureTypeSingle,
		Data:          dataBytes,
	}

	object := &object{
		Transaction: transaction,
		data:        data,
	}

	return object.setType(data.Type()), nil
}

type Data interface {
	encode() ([]byte, error)

	// Get transaction data type
	Type() Type

	// Get transaction data fee
	Fee() Fee
}

type CoinID uint32

func (c CoinID) String() string { return strconv.Itoa(int(c)) }

type Coin [10]byte

func (c Coin) String() string { return string(bytes.Trim(c[:], "\x00")) }

type EncodeInterface interface {
	Encode() (string, error)
}

type SignedTransaction interface {
	EncodeInterface
	GetTransaction() *Transaction
	Fee() *big.Int
	Hash() (string, error)
	Data() Data
	Signature() (signature, error)
	AddSignature(signatures ...[]byte) (SignedTransaction, error)
	SignatureData() []byte
	SingleSignatureData() ([]byte, error)
	SenderAddress() (string, error)
	Sign(prKey string, multisigPrKeys ...string) (SignedTransaction, error)
}

type Interface interface {
	EncodeInterface
	setType(t Type) Interface
	SetSignatureType(signatureType SignatureType) Interface
	SetMultiSignatureType() Interface
	setSignature(signature signature) (SignedTransaction, error)
	SetNonce(nonce uint64) Interface
	SetGasCoin(id CoinID) Interface
	SetGasPrice(price uint8) Interface
	SetPayload(payload []byte) Interface
	SetServiceData(serviceData []byte) Interface
	Sign(key string, multisigPrKeys ...string) (SignedTransaction, error)
	Clone() Interface
}

type object struct {
	*Transaction
	data Data
}

// Get Fee of transaction in PIP
func (o *object) Fee() *big.Int {
	gasPrice := big.NewInt(0).Mul(big.NewInt(int64(o.data.Fee())), big.NewInt(1000000000000000))
	commission := big.NewInt(0).Add(big.NewInt(0).Mul(big.NewInt(int64(len(o.Payload))*2), big.NewInt(1000000000000000)), big.NewInt(0).Mul(big.NewInt(int64(len(o.ServiceData))*2), big.NewInt(1000000000000000)))
	return big.NewInt(0).Add(gasPrice, commission)
}

// Get data of the transaction
func (o *object) Data() Data {
	return o.data
}

// Make a copy of the transaction
func (o *object) Clone() Interface {
	tx := *o.Transaction
	return &object{Transaction: &tx, data: o.data}
}

// Get Transaction
func (o *object) GetTransaction() *Transaction {
	return o.Transaction
}

// Get Signature
func (o *object) SignatureData() []byte {
	return o.Transaction.SignatureData
}

// Get single SignatureData
func (o *object) SingleSignatureData() ([]byte, error) {
	s, err := o.Signature()
	if err != nil {
		return nil, err
	}
	return s.firstSig()
}

func (o *object) Signature() (signature, error) {
	var signature signature
	switch o.SignatureType {
	case SignatureTypeSingle:
		signature = &Signature{}
	case SignatureTypeMulti:
		signature = &SignatureMulti{}
	default:
		return nil, errors.New("not set signature type")
	}

	if len(o.SignatureData()) == 0 {
		return signature, nil
	}

	err := rlp.DecodeBytes(o.SignatureData(), signature)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// Decode transaction
func Decode(tx string) (SignedTransaction, error) {
	if !strings.HasPrefix(tx, "0x") {
		return nil, errors.New("transaction don't has prefix '0x'")
	}

	decodeString, err := hex.DecodeString(tx[2:])
	if err != nil {
		return nil, err
	}

	transaction := new(Transaction)
	err = rlp.DecodeBytes(decodeString, transaction)
	if err != nil {
		return nil, err
	}

	var data Data
	switch transaction.Type {
	case TypeSend:
		data = &SendData{}
	case TypeSellCoin:
		data = &SellCoinData{}
	case TypeSellAllCoin:
		data = &SellAllCoinData{}
	case TypeBuyCoin:
		data = &BuyCoinData{}
	case TypeCreateCoin:
		data = &CreateCoinData{}
	case TypeDeclareCandidacy:
		data = &DeclareCandidacyData{}
	case TypeDelegate:
		data = &DelegateData{}
	case TypeUnbond:
		data = &UnbondData{}
	case TypeRedeemCheck:
		data = &RedeemCheckData{}
	case TypeSetCandidateOnline:
		data = &SetCandidateOnData{}
	case TypeSetCandidateOffline:
		data = &SetCandidateOffData{}
	case TypeCreateMultisig:
		data = &CreateMultisigData{}
	case TypeMultisend:
		data = &MultisendData{}
	case TypeEditCandidate:
		data = &EditCandidateData{}
	default:
		return nil, errors.New("unknown transaction type")
	}

	err = rlp.DecodeBytes(transaction.Data, data)
	if err != nil {
		return nil, err
	}

	result := &object{
		Transaction: transaction,
		data:        data,
	}
	return result, nil
}

// Get sender address
func (o *object) SenderAddress() (string, error) {
	if o.SignatureType == SignatureTypeSingle {
		hash, err := rlpHash([]interface{}{
			o.Transaction.Nonce,
			o.Transaction.ChainID,
			o.Transaction.GasPrice,
			o.Transaction.GasCoin,
			o.Transaction.Type,
			o.Transaction.Data,
			o.Transaction.Payload,
			o.Transaction.ServiceData,
			o.Transaction.SignatureType,
		})
		if err != nil {
			return "", err
		}

		signature, err := o.Signature()
		if err != nil {
			return "", err
		}

		ecrecover, err := crypto.Ecrecover(hash[:], signature.(*Signature).toBytes())
		if err != nil {
			return "", err
		}

		address, err := wallet.AddressByPublicKey("Mp" + hex.EncodeToString(ecrecover))
		if err != nil {
			return "", err
		}

		return address, nil
	}

	signature, err := o.Signature()
	if err != nil {
		return "", err
	}

	return wallet.BytesToAddress(signature.(*SignatureMulti).Multisig), nil
}

type Transaction struct {
	Nonce         uint64
	ChainID       ChainID
	GasPrice      uint8
	GasCoin       CoinID
	Type          Type
	Data          []byte
	Payload       []byte
	ServiceData   []byte
	SignatureType SignatureType
	SignatureData []byte
}

type signature interface {
	encode() ([]byte, error)
	firstSig() ([]byte, error)
}

type Signature struct {
	V *big.Int
	R *big.Int
	S *big.Int
}

func (s *Signature) encode() ([]byte, error) {
	return rlp.EncodeToBytes(s)
}

func (s *Signature) firstSig() ([]byte, error) {
	return s.encode()
}

func decodeSignature(b []byte) (*Signature, error) {
	s := &Signature{}
	err := rlp.DecodeBytes(b, s)
	if err != nil {
		return nil, err
	}
	return s, err
}

func (s *Signature) toBytes() []byte {
	sig := make([]byte, 65)
	copy(sig[:32], s.R.Bytes())
	copy(sig[32:64], s.S.Bytes())
	sig[64] = s.V.Bytes()[0] - 27

	return sig
}

type SignatureMulti struct {
	Multisig   [20]byte
	Signatures []*Signature
}

func (s *SignatureMulti) encode() ([]byte, error) {
	return rlp.EncodeToBytes(s)
}

func (s *SignatureMulti) firstSig() ([]byte, error) {
	if len(s.Signatures) == 0 {
		return nil, errors.New("signature not set")
	}
	return s.Signatures[0].encode()
}

func (o *object) setType(t Type) Interface {
	o.Type = t
	return o
}

func (o *object) SetSignatureType(signatureType SignatureType) Interface {
	o.SignatureType = signatureType
	return o
}

func (o *object) SetMultiSignatureType() Interface {
	o.SignatureType = SignatureTypeMulti
	return o
}

func (o *object) setSignature(signature signature) (SignedTransaction, error) {
	var err error
	o.Transaction.SignatureData, err = signature.encode()
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (o *object) SetNonce(nonce uint64) Interface {
	o.Nonce = nonce
	return o
}

func (o *object) SetGasCoin(id CoinID) Interface {
	o.GasCoin = id
	return o
}

func (o *object) SetGasPrice(price uint8) Interface {
	o.GasPrice = price
	return o
}

func (o *object) SetPayload(payload []byte) Interface {
	o.Payload = payload
	return o
}

func (o *object) SetServiceData(serviceData []byte) Interface {
	o.ServiceData = serviceData
	return o
}

func (tx *Transaction) Encode() (string, error) {
	src, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "", err
	}

	return "0x" + hex.EncodeToString(src), err
}

// Get hash of transaction
func (o *object) Hash() (string, error) {
	encode, err := o.Transaction.Encode()
	if err != nil {
		return "", err
	}

	b := make([]byte, hex.DecodedLen(len(encode)-2))
	_, err = hex.Decode(b, []byte(encode)[2:])
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(b)
	return "Mt" + hex.EncodeToString(hash[:]), nil
}

func (o *object) addSignature(signatures ...*Signature) (SignedTransaction, error) {
	signature, err := o.Signature()
	if err != nil {
		return nil, err
	}
	if len(signatures) == 0 {
		return nil, errors.New("number of signatures must be greater than 0")
	}
	if o.SignatureType == SignatureTypeSingle {
		return o.setSignature(signatures[0])
	}
	if len(o.SignatureData()) == 0 {
		return nil, errors.New("multisig address not set")
	}
	signatureMulti := signature.(*SignatureMulti)
	signatureMulti.Signatures = append(signatureMulti.Signatures, signatures...)
	return o.setSignature(signatureMulti)
}

func (o *object) AddSignature(signatures ...[]byte) (SignedTransaction, error) {
	signature, err := o.Signature()
	if err != nil {
		return nil, err
	}
	if len(signatures) == 0 {
		return nil, errors.New("number of signatures must be greater than 0")
	}
	if o.SignatureType == SignatureTypeSingle {
		sig, err := decodeSignature(signatures[0])
		if err != nil {
			return nil, err
		}
		return o.setSignature(sig)
	}
	if len(o.SignatureData()) == 0 {
		return nil, errors.New("multisig address not set")
	}
	signatureMulti := signature.(*SignatureMulti)
	signs := make([]*Signature, 0, len(signatures))
	for _, signature := range signatures {
		sig, err := decodeSignature(signature)
		if err != nil {
			return nil, err
		}
		signs = append(signs, sig)
	}
	signatureMulti.Signatures = append(signatureMulti.Signatures, signs...)
	return o.setSignature(signatureMulti)
}

// sign transaction
func (o *object) Sign(key string, multisigPrKeys ...string) (SignedTransaction, error) {
	h, err := rlpHash([]interface{}{
		o.Transaction.Nonce,
		o.Transaction.ChainID,
		o.Transaction.GasPrice,
		o.Transaction.GasCoin,
		o.Transaction.Type,
		o.Transaction.Data,
		o.Transaction.Payload,
		o.Transaction.ServiceData,
		o.Transaction.SignatureType,
	})
	if err != nil {
		return nil, err
	}

	switch o.SignatureType {
	case SignatureTypeSingle:
		signature, err := newSignature(key, h)
		if err != nil {
			return nil, err
		}
		return o.addSignature(signature)
	case SignatureTypeMulti:
		if len(o.SignatureData()) == 0 {
			sig := &SignatureMulti{
				Multisig:   [20]byte{},
				Signatures: make([]*Signature, 0, len(multisigPrKeys)),
			}
			addressToHex, err := wallet.AddressToHex(key)
			if err != nil {
				return nil, err
			}
			copy(sig.Multisig[:], addressToHex)
			_, err = o.setSignature(sig)
			if err != nil {
				return nil, err
			}
		}
		_, err := o.Signature()
		if err != nil {
			return nil, err
		}

		if len(multisigPrKeys) == 0 {
			return o, nil
		}
		signatures := make([]*Signature, 0, len(multisigPrKeys))
		for _, prKey := range multisigPrKeys {
			signature, err := newSignature(prKey, h)
			if err != nil {
				return nil, err
			}

			signatures = append(signatures, signature)
		}

		return o.addSignature(signatures...)
	default:
		return nil, fmt.Errorf("undefined signature type: %d", o.SignatureType)
	}
}

func newSignature(prKey string, h [32]byte) (*Signature, error) {
	sig, err := sign(prKey, h)
	if err != nil {
		return nil, err
	}
	return &Signature{
		R: new(big.Int).SetBytes(sig[:32]),
		S: new(big.Int).SetBytes(sig[32:64]),
		V: new(big.Int).SetBytes([]byte{sig[64] + 27}),
	}, nil
}

func sign(prKey string, h [32]byte) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(prKey)
	if err != nil {
		return nil, err
	}

	sig, err := crypto.Sign(h[:], privateKey)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func rlpHash(x interface{}) (h [32]byte, err error) {
	hw := sha3.NewLegacyKeccak256()
	err = rlp.Encode(hw, x)
	if err != nil {
		return h, err
	}
	hw.Sum(h[:0])
	return h, nil
}
