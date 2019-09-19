package transaction

import (
	"encoding/hex"
	"errors"
	//"github.com/MinterTeam/minter-go-sdk/wallet"
	"github.com/ethereum/go-ethereum/common/math"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
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

	switch data.(type) {
	case *SendData:
		return transaction.setType(TypeSend), nil
	case *SellCoinData:
		return transaction.setType(TypeSellCoin), nil
	case *SellAllCoinData:
		return transaction.setType(TypeSellAllCoin), nil
	case *BuyCoinData:
		return transaction.setType(TypeBuyCoin), nil
	case *CreateCoinData:
		return transaction.setType(TypeCreateCoin), nil
	case *DeclareCandidacyData:
		return transaction.setType(TypeDeclareCandidacy), nil
	case *DelegateData:
		return transaction.setType(TypeDelegate), nil
	case *UnbondData:
		return transaction.setType(TypeUnbond), nil
	case *RedeemCheckData:
		return transaction.setType(TypeRedeemCheck), nil
	case *SetCandidateOnData:
		return transaction.setType(TypeSetCandidateOnline), nil
	case *SetCandidateOffData:
		return transaction.setType(TypeSetCandidateOffline), nil
	//case *CreateMultisigData:
	//	return transaction.setType(TypeCreateMultisig), nil
	case *MultiMultisendDataItem:
		return transaction.setType(TypeMultisend), nil
	case *EditCandidateData:
		return transaction.setType(TypeEditCandidate), nil

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
	Sign(prKey string) (SignedTransaction, error)
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

func (tx *Transaction) Sign(prKey string) (SignedTransaction, error) {
	privateKey, err := ethcrypto.HexToECDSA(prKey)
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
