package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"math/big"
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
	TypeEditCoinOwner
	TypeEditMultisig
	TypePriceVote
	TypeEditCandidatePublicKey
)

// For each transaction sender should pay fee. Fees are measured in "units".
type fee uint

const (
	feeTypeSend                fee = 10
	feeTypeSellCoin            fee = 100
	feeTypeSellAllCoin         fee = 100
	feeTypeBuyCoin             fee = 100
	feeTypeCreateCoin          fee = 1000
	feeTypeDeclareCandidacy    fee = 10000
	feeTypeDelegate            fee = 200
	feeTypeUnbond              fee = 200
	feeTypeRedeemCheck         fee = 30
	feeTypeSetCandidateOnline  fee = 100
	feeTypeSetCandidateOffline fee = 100
	feeTypeCreateMultisig      fee = 100
	// feeMultisend fee =  10+(n-1)*5
	feeTypeEditCandidate          fee = 100000
	feeTypeEditCandidatePublicKey fee = 100000000
	feeTypeSetHaltBlock           fee = 1000
	feeTypeRecreateCoin           fee = 10000000
	feeTypeEditCoinOwner          fee = 10000000
	feeEditMultisig               fee = 1000
	feePriceVote                  fee = 10
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

func (b *Builder) NewTransaction(data DataInterface) (Interface, error) {
	dataBytes, err := data.encode()
	if err != nil {
		return nil, err
	}

	transaction := &Transaction{
		ChainID:       b.ChainID,
		SignatureType: SignatureTypeSingle,
		Data:          dataBytes,
	}

	object := object{
		Transaction: transaction,
		data:        data,
	}

	switch data.(type) {
	case *SendData:
		return object.setType(TypeSend), nil
	case *SellCoinData:
		return object.setType(TypeSellCoin), nil
	case *SellAllCoinData:
		return object.setType(TypeSellAllCoin), nil
	case *BuyCoinData:
		return object.setType(TypeBuyCoin), nil
	case *CreateCoinData:
		return object.setType(TypeCreateCoin), nil
	case *DeclareCandidacyData:
		return object.setType(TypeDeclareCandidacy), nil
	case *DelegateData:
		return object.setType(TypeDelegate), nil
	case *UnbondData:
		return object.setType(TypeUnbond), nil
	case *RedeemCheckData:
		return object.setType(TypeRedeemCheck), nil
	case *SetCandidateOnData:
		return object.setType(TypeSetCandidateOnline), nil
	case *SetCandidateOffData:
		return object.setType(TypeSetCandidateOffline), nil
	case *CreateMultisigData:
		return object.setType(TypeCreateMultisig), nil
	case *MultisendData:
		return object.setType(TypeMultisend), nil
	case *EditCandidateData:
		return object.setType(TypeEditCandidate), nil

	default:
		return nil, errors.New("unknown transaction type")
	}
}

type DataInterface interface {
	encode() ([]byte, error)
	fee() fee
}

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
	Data() DataInterface
	Signature() (signatureInterface, error)
	AddSignature(signatures ...[]byte) (SignedTransaction, error)
	SignatureData() []byte
	SimpleSignatureData() ([]byte, error)
	SenderAddress() (string, error)
	Sign(prKey string, multisigPrKeys ...string) (SignedTransaction, error)
}

type Interface interface {
	EncodeInterface
	setType(t Type) Interface
	SetSignatureType(signatureType SignatureType) Interface
	SetMultiSignatureType() Interface
	setSignature(signature signatureInterface) (SignedTransaction, error)
	SetNonce(nonce uint64) Interface
	SetGasCoin(name string) Interface
	SetGasPrice(price uint8) Interface
	SetPayload(payload []byte) Interface
	SetServiceData(serviceData []byte) Interface
	Sign(key string, multisigPrKeys ...string) (SignedTransaction, error)
	Clone() Interface
}

type object struct {
	*Transaction
	data   DataInterface
	encode string
}

// Get fee of transaction in PIP
func (o *object) Fee() *big.Int {
	gasPrice := big.NewInt(0).Mul(big.NewInt(int64(o.data.fee())), big.NewInt(1000000000000000))
	commission := big.NewInt(0).Add(big.NewInt(0).Mul(big.NewInt(int64(len(o.Payload))*2), big.NewInt(1000000000000000)), big.NewInt(0).Mul(big.NewInt(int64(len(o.ServiceData))*2), big.NewInt(1000000000000000)))
	return big.NewInt(0).Add(gasPrice, commission)
}

func (o *object) Data() DataInterface {
	return o.data
}

func (o *object) Clone() Interface {
	tx := *o.Transaction
	return &object{Transaction: &tx, data: o.data}
}

func (o *object) GetTransaction() *Transaction {
	return o.Transaction
}

func (o *object) SignatureData() []byte {
	return o.Transaction.SignatureData
}

func (o *object) SimpleSignatureData() ([]byte, error) {
	s, err := o.Signature()
	if err != nil {
		return nil, err
	}
	return s.firstSig()
}

func (o *object) Signature() (signatureInterface, error) {
	var signature signatureInterface
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

	var data interface{}
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
	case TypeEditCoinOwner:
		data = &EditCoinOwnerData{}
	case TypeEditCandidatePublicKey:
		data = &EditCandidatePublicKeyData{}
	default:
		return nil, errors.New("unknown transaction type")
	}

	err = rlp.DecodeBytes(transaction.Data, data)
	if err != nil {
		return nil, err
	}

	result := &object{
		Transaction: transaction,
		data:        data.(DataInterface),
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
	GasCoin       Coin
	Type          Type
	Data          []byte
	Payload       []byte
	ServiceData   []byte
	SignatureType SignatureType
	SignatureData []byte
}

type signatureInterface interface {
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

func (o *object) setSignature(signature signatureInterface) (SignedTransaction, error) {
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

func (o *object) SetGasCoin(name string) Interface {
	copy(o.GasCoin[:], name)
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

func (o *object) Encode() (string, error) {
	if o.encode != "" {
		return o.encode, nil
	}

	src, err := rlp.EncodeToBytes(o.Transaction)
	if err != nil {
		return "", err
	}

	return "0x" + hex.EncodeToString(src), err
}

// Get hash of transaction
func (o *object) Hash() (string, error) {
	encode, err := o.Encode()
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
	signedTx, err := o.sign(key, multisigPrKeys...)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func signature(prKey string, h [32]byte) (*Signature, error) {
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
