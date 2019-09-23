package transaction

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
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

	object := object{
		Transaction: transaction,
		fee:         uint(data.fee()),
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
	//case *CreateMultisigData:
	//	return transaction.setType(TypeCreateMultisig), nil
	case *MultiMultisendDataItem:
		return object.setType(TypeMultisend), nil
	case *EditCandidateData:
		return object.setType(TypeEditCandidate), nil

	default:
		return nil, errors.New("") //todo
	}
}

type DataInterface interface {
	encode() ([]byte, error)
	fee() Fee
}

type SignedTransaction interface {
	Encode() ([]byte, error)
	Fee() *big.Int
	Hash() (string, error)
}

type Interface interface {
	setType(t Type) Interface
	setFee(commission Fee) Interface
	SetNonce(nonce uint64) Interface
	SetGasCoin(name string) Interface
	SetGasPrice(price uint8) Interface
	//todo Decode(string) Interface
	Sign(prKey string) (SignedTransaction, error)
}

type object struct {
	*Transaction
	fee uint
}

func (o *object) Fee() *big.Int {
	gasPrice := big.NewInt(0).Mul(big.NewInt(int64(o.fee)), big.NewInt(1000000000000000))
	commission := big.NewInt(0).Add(big.NewInt(0).Mul(big.NewInt(int64(len(o.Payload))*2), big.NewInt(1000000000000000)), big.NewInt(0).Mul(big.NewInt(int64(len(o.ServiceData))*2), big.NewInt(1000000000000000)))
	//todo: testing
	return big.NewInt(0).Mul(gasPrice, commission)
}

func (o *object) setFee(commission Fee) Interface {
	return o
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

func (o *object) Hash() (string, error) {
	encode, err := o.Transaction.Encode()
	if err != nil {
		return "", err
	}
	bytes := make([]byte, hex.DecodedLen(len(encode)-2))
	_, err = hex.Decode(bytes, encode[2:])
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(bytes)

	return "Mt" + hex.EncodeToString(hash[:])[:40], nil
}

func (o *object) Sign(prKey string) (SignedTransaction, error) {
	privateKey, err := crypto.HexToECDSA(prKey)
	if err != nil {
		return nil, err
	}

	x := []interface{}{
		o.Transaction.Nonce,
		o.Transaction.ChainID,
		o.Transaction.GasPrice,
		o.Transaction.GasCoin,
		o.Transaction.Type,
		o.Transaction.Data,
		o.Transaction.Payload,
		o.Transaction.ServiceData,
		o.Transaction.SignatureType,
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
