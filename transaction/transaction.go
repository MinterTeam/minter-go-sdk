package transaction

import (
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

type Type byte

const (
	TypeSend Type = iota + 1
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
	signatureTypeSingle
	signatureTypeMulti
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
		return object.setType(TypeCreateMultisig).setSignatureType(signatureTypeMulti), nil
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
	fee() Fee
}

type SignedTransaction interface {
	Encode() (string, error)
	Fee() *big.Int
	Hash() (string, error)
	Data() DataInterface
	Signature() (*Signature, error)
	SenderAddress() (string, error)
	PublicKey() (string, error)
}

type Interface interface {
	setType(t Type) Interface
	setSignatureType(signatureType SignatureType) Interface
	SetNonce(nonce uint64) Interface
	SetGasCoin(name string) Interface
	SetGasPrice(price uint8) Interface
	SetPayload(payload []byte) Interface
	SetServiceData(serviceData []byte) Interface
	Sign(prKey string) (SignedTransaction, error)
}

type object struct {
	*Transaction
	data DataInterface
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

func (o *object) Signature() (*Signature, error) {
	signature := new(Signature)
	err := rlp.DecodeBytes(o.SignatureData, signature)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// Decode transaction
func Decode(tx string) (SignedTransaction, error) {
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
		data = &SellCoinData{}
	case TypeCreateCoin:
		data = &BuyCoinData{}
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
		data:        data.(DataInterface),
	}
	return result, nil
}

// Get sender address
func (o *object) SenderAddress() (string, error) {
	publicKey, err := o.PublicKey()
	if err != nil {
		return "", err
	}

	address, err := wallet.AddressByPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	return address, nil
}

// Recover public key
func (o *object) PublicKey() (string, error) {
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

	sig := make([]byte, 65)

	copy(sig[:32], signature.R.Bytes())
	copy(sig[32:64], signature.S.Bytes())
	sig[64] = signature.V.Bytes()[0] - 27

	ecrecover, err := crypto.Ecrecover(hash[:], sig)
	if err != nil {
		return "", err
	}

	return "Mp" + hex.EncodeToString(ecrecover)[2:], nil
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

func (o *object) setType(t Type) Interface {
	o.Type = t
	return o
}

func (o *object) setSignatureType(signatureType SignatureType) Interface {
	o.SignatureType = signatureType
	return o
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
	bytes := make([]byte, hex.DecodedLen(len(encode)-2))
	_, err = hex.Decode(bytes, []byte(encode)[2:])
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(bytes)

	return "Mt" + hex.EncodeToString(hash[:])[:40], nil
}

// Sign transaction
func (o *object) Sign(prKey string) (SignedTransaction, error) {
	privateKey, err := crypto.HexToECDSA(prKey)
	if err != nil {
		return nil, err
	}

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

	sig, err := crypto.Sign(h[:], privateKey)
	if err != nil {
		return nil, err
	}

	o.Transaction.SignatureData, err = rlp.EncodeToBytes(&Signature{
		R: new(big.Int).SetBytes(sig[:32]),
		S: new(big.Int).SetBytes(sig[32:64]),
		V: new(big.Int).SetBytes([]byte{sig[64] + 27}),
	})
	if err != nil {
		return nil, err
	}

	return o, nil
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

func addressToHex(address string) ([]byte, error) {
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
