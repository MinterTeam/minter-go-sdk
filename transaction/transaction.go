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

// Type of transaction is determined by a single byte.
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
	TypeSetHaltBlock
	TypeRecreateCoin
	TypeChangeOwner
)

// For each transaction sender should pay fee. Fees are measured in "units".
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
	feeTypeSetHaltBlock  Fee = 1000
	feeTypeRecreateCoin  Fee = 10000000
	feeTypeChangeOwner   Fee = 10000000
)

type SignatureType byte

const (
	_ SignatureType = iota
	SignatureTypeSingle
	SignatureTypeMulti
)

// id of the network (1 - mainnet, 2 - testnet)
type ChainID byte

const (
	_ ChainID = iota
	MainNetChainID
	TestNetChainID
)

type Builder struct {
	ChainID ChainID
}

// New builder for transactions
func NewBuilder(chainID ChainID) *Builder {
	return &Builder{ChainID: chainID}
}

// New transaction from data
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

	object := &Object{
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

// ID of a coin
type CoinID uint32

func (c CoinID) String() string { return strconv.Itoa(int(c)) }

// Symbol of a coin
type CoinSymbol [10]byte

func (c CoinSymbol) String() string { return string(bytes.Trim(c[:], "\x00")) }

type EncodeInterface interface {
	// Get string representation of a transaction.
	Encode() (string, error)
}

type SignedTransaction interface {
	EncodeInterface
	// Get Transaction
	GetTransaction() *Transaction
	// Get fee of transaction in PIP. Also sender should pay extra 2 units per byte in Payload and Service Data fields.
	Fee() *big.Int
	// Get hash of transaction
	Hash() (string, error)
	// Get data of the transaction
	Data() Data
	// Get signature interface
	Signature() (SignatureInterface, error)
	// Add signature from bytes
	AddSignature(signatures ...[]byte) (SignedTransaction, error)
	// Get bytes of Signature
	SignatureData() []byte
	// Get single SignatureData
	SingleSignatureData() ([]byte, error)
	// Get sender address
	SenderAddress() (string, error)
	// Get set of signers
	Signers() ([]string, error)
	// Sign transaction
	Sign(prKey string, multisigPrKeys ...string) (SignedTransaction, error)
	// Make a copy of the transaction
	Clone() Interface
}

type Interface interface {
	EncodeInterface
	// Set signature type
	SetSignatureType(signatureType SignatureType) Interface
	// Set signature type to SignatureTypeMulti
	SetMultiSignatureType() Interface
	// Set nonce of transaction
	SetNonce(nonce uint64) Interface
	// Set ID of a coin to pay fee
	SetGasCoin(id CoinID) Interface
	// Set fee multiplier
	SetGasPrice(price uint8) Interface
	// Set arbitrary user-defined bytes
	SetPayload(payload []byte) Interface
	// Set ServiceData field
	SetServiceData(serviceData []byte) Interface
	// Sign transaction
	Sign(key string, multisigPrKeys ...string) (SignedTransaction, error)
	// Make a copy of the transaction
	Clone() Interface

	setType(t Type) Interface
	setSignature(signature SignatureInterface) (SignedTransaction, error)
}

type Object struct {
	*Transaction
	data Data
}

// Get fee of transaction in PIP. Also sender should pay extra 2 units per byte in Payload and Service Data fields.
func (o *Object) Fee() *big.Int {
	gasPrice := big.NewInt(0).Mul(big.NewInt(int64(o.data.Fee())), big.NewInt(1000000000000000))
	commission := big.NewInt(0).Add(big.NewInt(0).Mul(big.NewInt(int64(len(o.Payload))*2), big.NewInt(1000000000000000)), big.NewInt(0).Mul(big.NewInt(int64(len(o.ServiceData))*2), big.NewInt(1000000000000000)))
	return big.NewInt(0).Add(gasPrice, commission)
}

// Get data of the transaction
func (o *Object) Data() Data {
	return o.data
}

// Make a copy of the transaction
func (o *Object) Clone() Interface {
	tx := *o.Transaction
	return &Object{Transaction: &tx, data: o.data}
}

// Get Transaction
func (o *Object) GetTransaction() *Transaction {
	return o.Transaction
}

// Get bytes of Signature
func (o *Object) SignatureData() []byte {
	return o.Transaction.SignatureData
}

// Get first SignatureData
func (o *Object) SingleSignatureData() ([]byte, error) {
	s, err := o.Signature()
	if err != nil {
		return nil, err
	}
	return s.Single()
}

// Get signature interface
func (o *Object) Signature() (SignatureInterface, error) {
	var signature SignatureInterface
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
func Decode(tx string) (*Object, error) {
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
	case TypeSetHaltBlock:
		data = &SetHaltBlockData{}
	case TypeRecreateCoin:
		data = &RecreateCoinData{}
	case TypeChangeOwner:
		data = &ChangeOwnerData{}
	default:
		return nil, errors.New("unknown transaction type")
	}

	err = rlp.DecodeBytes(transaction.Data, data)
	if err != nil {
		return nil, err
	}

	result := &Object{
		Transaction: transaction,
		data:        data,
	}
	return result, nil
}

// Get sender address
func (o *Object) SenderAddress() (string, error) {
	signature, err := o.Signature()
	if err != nil {
		return "", err
	}

	if single, ok := signature.(*Signature); o.SignatureType == SignatureTypeSingle && ok {
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

		return single.signer(hash, SignatureTypeSingle)
	}

	if multi, ok := signature.(*SignatureMulti); o.SignatureType == SignatureTypeMulti && ok {
		return wallet.BytesToAddress(multi.Multisig), nil
	}

	return "", errors.New("signature is invalid")
}

type Transaction struct {
	Nonce         uint64        // used for prevent transaction reply
	ChainID       ChainID       // id of the network (1 - mainnet, 2 - testnet)
	GasPrice      uint8         // fee multiplier, should be equal or greater than current mempool min gas price.
	GasCoin       CoinID        // ID of a coin to pay fee, right padded with zeros
	Type          Type          // type of transaction
	Data          []byte        // data of transaction (depends on transaction type)
	Payload       []byte        // arbitrary user-defined bytes
	ServiceData   []byte        // reserved field
	SignatureType SignatureType // single or multisig transaction
	SignatureData []byte        // digital signature of transaction
}

type SignatureInterface interface {
	// Get digital signature of transaction
	Encode() ([]byte, error)
	// Get SingleSignature
	Single() ([]byte, error)
	// Get signature type
	Type() SignatureType
}

// Simple Signature
type Signature struct {
	V *big.Int
	R *big.Int
	S *big.Int
}

// Get signature type
func (s *Signature) Type() SignatureType {
	return SignatureTypeSingle
}

// Get digital signature of transaction
func (s *Signature) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(s)
}

// Get digital signature of transaction
func (s *Signature) Single() ([]byte, error) {
	return s.Encode()
}

// Get signer address
func (s *Signature) signer(hash [32]byte, t SignatureType) (string, error) {
	publicKey, err := crypto.Ecrecover(hash[:], s.toBytes())
	if err != nil {
		return "", err
	}

	address, err := wallet.AddressByPublicKey("Mp" + hex.EncodeToString(publicKey[t-1:]))
	if err != nil {
		return "", err
	}

	return address, nil
}

func decodeSignature(b []byte) (*Signature, error) {
	s := &Signature{}
	err := rlp.DecodeBytes(b, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Signature) toBytes() []byte {
	R, S := s.R.Bytes(), s.S.Bytes()
	sig := make([]byte, 65)
	copy(sig[32-len(R):32], R)
	copy(sig[64-len(S):64], S)
	sig[64] = byte(s.V.Uint64() - 27)

	return sig
}

// Signature from Multisig address
type SignatureMulti struct {
	Multisig   [20]byte
	Signatures []*Signature
}

// Get signature type
func (s *SignatureMulti) Type() SignatureType {
	return SignatureTypeMulti
}

// Get digital signature of transaction
func (s *SignatureMulti) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(s)
}

// Get set of signers
func (o *Object) Signers() ([]string, error) {
	signatures, err := o.Signature()
	if err != nil {
		return nil, err
	}

	if multi, ok := signatures.(*SignatureMulti); o.SignatureType == SignatureTypeMulti && ok {
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
			return nil, err
		}

		signers := make([]string, 0, len(multi.Signatures))
		for _, signature := range multi.Signatures {
			address, err := signature.signer(hash, SignatureTypeMulti)
			if err != nil {
				return nil, err
			}

			signers = append(signers, address)
		}
		return signers, nil
	}

	address, err := o.SenderAddress()
	if err != nil {
		return nil, err
	}

	return []string{address}, nil
}

// Get first SingleSignature
func (s *SignatureMulti) Single() ([]byte, error) {
	if len(s.Signatures) == 0 {
		return nil, errors.New("signature not set")
	}
	return s.Signatures[0].Encode()
}

func (o *Object) setType(t Type) Interface {
	o.Type = t
	return o
}

// Set signature type
func (o *Object) SetSignatureType(signatureType SignatureType) Interface {
	o.SignatureType = signatureType
	return o
}

// Set signature type to SignatureTypeMulti
func (o *Object) SetMultiSignatureType() Interface {
	o.SignatureType = SignatureTypeMulti
	return o
}

func (o *Object) setSignature(signature SignatureInterface) (SignedTransaction, error) {
	var err error
	o.Transaction.SignatureData, err = signature.Encode()
	if err != nil {
		return nil, err
	}

	return o, nil
}

// Set nonce of transaction
func (o *Object) SetNonce(nonce uint64) Interface {
	o.Nonce = nonce
	return o
}

// Set ID of a coin to pay fee
func (o *Object) SetGasCoin(id CoinID) Interface {
	o.GasCoin = id
	return o
}

// Set fee multiplier
func (o *Object) SetGasPrice(price uint8) Interface {
	o.GasPrice = price
	return o
}

// Set arbitrary user-defined bytes
func (o *Object) SetPayload(payload []byte) Interface {
	o.Payload = payload
	return o
}

// Set ServiceData field
func (o *Object) SetServiceData(serviceData []byte) Interface {
	o.ServiceData = serviceData
	return o
}

// Get string representation of a transaction. RLP-encoded structure in hex format.
func (tx *Transaction) Encode() (string, error) {
	src, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "", err
	}

	return "0x" + hex.EncodeToString(src), err
}

// Get hash of transaction
func (o *Object) Hash() (string, error) {
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

func (o *Object) addSignature(signatures ...*Signature) (SignedTransaction, error) {
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

// Add signature from bytes
func (o *Object) AddSignature(signatures ...[]byte) (SignedTransaction, error) {
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

// Sign transaction
func (o *Object) Sign(key string, multisigPrKeys ...string) (SignedTransaction, error) {
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
