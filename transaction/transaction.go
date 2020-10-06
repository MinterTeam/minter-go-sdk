package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strconv"
	"strings"
)

// Type of transaction is determined by a single byte.
type Type byte

// Types of Data
const (
	_                          Type = iota
	TypeSend                        // 0x01
	TypeSellCoin                    // 0x02
	TypeSellAllCoin                 // 0x03
	TypeBuyCoin                     // 0x04
	TypeCreateCoin                  // 0x05
	TypeDeclareCandidacy            // 0x06
	TypeDelegate                    // 0x07
	TypeUnbond                      // 0x08
	TypeRedeemCheck                 // 0x09
	TypeSetCandidateOnline          // 0x0A
	TypeSetCandidateOffline         // 0x0B
	TypeCreateMultisig              // 0x0C
	TypeMultisend                   // 0x0D
	TypeEditCandidate               // 0x0E
	TypeSetHaltBlock                // 0x0F
	TypeRecreateCoin                // 0x10
	TypeEditCoinOwner               // 0x11
	TypeEditMultisig                // 0x12
	TypePriceVote                   // 0x13
	TypeEditCandidatePublicKey      // 0x14
)

// Fee is the commission that the sender must pay for sending the transaction. Fees are measured in "units". Also sender should pay extra 2 units per byte in Payload and ServiceData fields.
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
	feeTypeEditCandidate          Fee = 100000
	feeTypeEditCandidatePublicKey Fee = 100000000
	feeTypeSetHaltBlock           Fee = 1000
	feeTypeRecreateCoin           Fee = 10000000
	feeTypeEditCoinOwner          Fee = 10000000
	feeEditMultisig               Fee = 1000
	feePriceVote                  Fee = 10
)

// SignatureType is type of signature (1 - SignatureTypeSingle, 2 - SignatureTypeMulti)
type SignatureType byte

// Types of Signature
const (
	_                   SignatureType = iota
	SignatureTypeSingle               // 0x01
	SignatureTypeMulti                // 0x02
)

// ChainID is network identifier  (1 - MainNetChainID, 2 - TestNetChainID)
type ChainID byte

// Variants of ChainID
const (
	_              ChainID = iota
	MainNetChainID         // 0x01
	TestNetChainID         // 0x02
)

// CoinID ID of a coin.
type CoinID uint32

// String return CoinID as string.
func (c *CoinID) String() string { return strconv.Itoa(int(*c)) }

// CoinSymbol is symbol of a coin.
type CoinSymbol [10]byte

// String returns CoinSymbol as string.
func (s *CoinSymbol) String() string { return string(bytes.Trim(s[:], "\x00")) }

// Address is address.
type Address [20]byte

// String returns Address as string.
func (a *Address) String() string { return "Mx" + hex.EncodeToString(a[:]) }

// PublicKey is public key.
type PublicKey [32]byte

// String returns PublicKey as string.
func (p *PublicKey) String() string { return "Mp" + hex.EncodeToString(p[:]) }

// BipToPip converts BIP to PIP (multiplies input by 1e18)
func BipToPip(bip *big.Int) *big.Int {
	return big.NewInt(0).Mul(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil), bip)
}

// Builder is creator of Transaction.
type Builder struct {
	ChainID ChainID
}

// NewBuilder returns new Builder for creating Transaction.
func NewBuilder(chainID ChainID) *Builder {
	return &Builder{ChainID: chainID}
}

// NewTransaction returns new transaction from data.
func (b *Builder) NewTransaction(data Data) (Interface, error) {
	dataBytes, err := data.Encode()
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

// Data is data of transaction
type Data interface {
	// Encode returns the byte representation of a transaction Data.
	Encode() ([]byte, error)
	// Type returns Data type of the transaction.
	Type() Type
	// Fee returns commission of transaction Data
	Fee() Fee
}

type encodeInterface interface {
	// Encode returns string representation of a transaction.
	Encode() (string, error)
}

// Signed is interface of signed Transaction.
type Signed interface {
	encodeInterface
	// GetTransaction returns Transaction struct
	GetTransaction() *Transaction
	// Fee returns fee of transaction in PIP. Also sender should pay extra 2 units per byte in Payload and ServiceData fields.
	Fee() *big.Int
	// Hash returns hash of Transaction.
	Hash() (string, error)
	// Data returns Data of the Transaction.
	Data() Data
	// Signature returns Signature interface.
	Signature() (Signature, error)
	// AddSignature adds signature from hex strings.
	AddSignature(signatures ...string) (Signed, error)
	// SignatureData returns bytes of Signature.
	SignatureData() []byte
	// SingleSignatureData returns SignatureData.
	SingleSignatureData() (string, error)
	// SenderAddress returns sender addresses.
	SenderAddress() (string, error)
	// Signers returns set of signers.
	Signers() ([]string, error)
	// Sign signs transaction with a private key.
	Sign(prKey string, multisigPrKeys ...string) (Signed, error)
	// Clone returns copy of the transaction.
	Clone() Interface
}

// Interface is Transaction data installer.
type Interface interface {
	encodeInterface
	// SetSignatureType sets signature type.
	SetSignatureType(signatureType SignatureType) Interface
	// SetMultiSignatureType sets signature type to SignatureTypeMulti.
	SetMultiSignatureType() Interface
	// SetNonce sets nonce of transaction.
	SetNonce(nonce uint64) Interface
	// SetGasCoin sets ID of a coin to pay fee.
	SetGasCoin(id uint64) Interface
	// SetGasPrice sets fee multiplier.
	SetGasPrice(price uint8) Interface
	// SetPayload sets arbitrary user-defined bytes.
	SetPayload(payload []byte) Interface
	// SetServiceData sets ServiceData field.
	SetServiceData(serviceData []byte) Interface
	// Sign signs transaction with a private key.
	Sign(key string, prKeys ...string) (Signed, error)
	// Clone returns copy of the transaction.
	Clone() Interface

	setType(t Type) Interface
	setSignature(signature Signature) (Signed, error)
}

type object struct {
	*Transaction
	data Data
}

// Fee returns fee of transaction in PIP. Also sender should pay extra 2 units per byte in Payload and ServiceData fields.
func (o *object) Fee() *big.Int {
	gasPrice := big.NewInt(0).Mul(big.NewInt(int64(o.data.Fee())), big.NewInt(1000000000000000))
	commission := big.NewInt(0).Add(big.NewInt(0).Mul(big.NewInt(int64(len(o.Payload))*2), big.NewInt(1000000000000000)), big.NewInt(0).Mul(big.NewInt(int64(len(o.ServiceData))*2), big.NewInt(1000000000000000)))
	return big.NewInt(0).Add(gasPrice, commission)
}

// Data returns Data of Transaction.
func (o *object) Data() Data {
	return o.data
}

// Clone returns copy of the transaction.
func (o *object) Clone() Interface {
	tx := *o.Transaction
	return &object{Transaction: &tx, data: o.data}
}

// GetTransaction returns Transaction struct.
func (o *object) GetTransaction() *Transaction {
	return o.Transaction
}

// SignatureData returns bytes of Signature.
func (o *object) SignatureData() []byte {
	return o.Transaction.SignatureData
}

// SingleSignatureData returns SignatureData.
func (o *object) SingleSignatureData() (string, error) {
	s, err := o.Signature()
	if err != nil {
		return "", err
	}

	single, err := s.Single()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(single), nil
}

// Signature returns Signature interface.
func (o *object) Signature() (Signature, error) {
	var signature Signature
	switch o.SignatureType {
	case SignatureTypeSingle:
		signature = &SignatureSingle{}
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

// Decode returns Signed model with Transaction and Data.
func Decode(tx string) (Signed, error) {
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
		data:        data,
	}
	return result, nil
}

// SenderAddress returns sender addresses.
func (o *object) SenderAddress() (string, error) {
	signature, err := o.Signature()
	if err != nil {
		return "", err
	}

	if single, ok := signature.(*SignatureSingle); o.SignatureType == SignatureTypeSingle && ok {
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

// Transaction is transaction model.
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

// Signature is interface of signatures
type Signature interface {
	// Encode returns digital signature of transaction.
	Encode() ([]byte, error)
	// Single returns SignatureSingle.
	Single() ([]byte, error)
	// Type returns signature type.
	Type() SignatureType
}

// SignatureSingle is single signature.
type SignatureSingle struct {
	V *big.Int
	R *big.Int
	S *big.Int
}

// Type returns SignatureType of Signature.
func (s *SignatureSingle) Type() SignatureType {
	return SignatureTypeSingle
}

// Encode returns digital signature of transaction.
func (s *SignatureSingle) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(s)
}

// Single returns SignatureSingle.
func (s *SignatureSingle) Single() ([]byte, error) {
	return s.Encode()
}

func (s *SignatureSingle) signer(hash PublicKey, t SignatureType) (string, error) {
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

func decodeSignature(b []byte) (*SignatureSingle, error) {
	s := &SignatureSingle{}
	err := rlp.DecodeBytes(b, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SignatureSingle) toBytes() []byte {
	R, S := s.R.Bytes(), s.S.Bytes()
	sig := make([]byte, 65)
	copy(sig[32-len(R):32], R)
	copy(sig[64-len(S):64], S)
	sig[64] = byte(s.V.Uint64() - 27)

	return sig
}

// MultisigAddress returns multisig address created by the sender at the block height.
func MultisigAddress(owner string, nonce uint64) string {
	o, err := wallet.AddressToHex(owner)
	if err != nil {
		panic(err)
	}

	var ownerAddress Address
	copy(ownerAddress[:], o)

	b, err := rlp.EncodeToBytes(&struct {
		Owner Address
		Nonce uint64
	}{Owner: ownerAddress, Nonce: nonce})
	if err != nil {
		panic(err)
	}

	var addr Address
	copy(addr[:], crypto.Keccak256(b)[12:])

	return wallet.BytesToAddress(addr)
}

// SignatureMulti is signature of multisig address
type SignatureMulti struct {
	Multisig   Address
	Signatures []*SignatureSingle
}

// Type returns SignatureType.
func (s *SignatureMulti) Type() SignatureType {
	return SignatureTypeMulti
}

// Encode returns digital signature of transaction.
func (s *SignatureMulti) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(s)
}

// Signers returns set of signers.
func (o *object) Signers() ([]string, error) {
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

// Single returns SingleSignature
func (s *SignatureMulti) Single() ([]byte, error) {
	if len(s.Signatures) == 0 {
		return nil, errors.New("signature not set")
	}

	return s.Signatures[0].Encode()
}

func (o *object) setType(t Type) Interface {
	o.Type = t
	return o
}

// SetSignatureType sets SignatureType of Transaction Data.
func (o *object) SetSignatureType(signatureType SignatureType) Interface {
	o.SignatureType = signatureType
	return o
}

// SetMultiSignatureType sets signature type to SignatureTypeMulti.
func (o *object) SetMultiSignatureType() Interface {
	o.SignatureType = SignatureTypeMulti
	return o
}

func (o *object) setSignature(signature Signature) (Signed, error) {
	var err error
	o.Transaction.SignatureData, err = signature.Encode()
	if err != nil {
		return nil, err
	}

	return o, nil
}

// SetNonce sets nonce of transaction.
func (o *object) SetNonce(nonce uint64) Interface {
	o.Nonce = nonce
	return o
}

// SetGasCoin sets CoinID of a coin to pay fee.
func (o *object) SetGasCoin(id uint64) Interface {
	o.GasCoin = CoinID(id)
	return o
}

// SetGasPrice sets fee multiplier.
func (o *object) SetGasPrice(price uint8) Interface {
	o.GasPrice = price
	return o
}

// SetPayload sets Payload field.
func (o *object) SetPayload(payload []byte) Interface {
	o.Payload = payload
	return o
}

// SetServiceData sets ServiceData field.
func (o *object) SetServiceData(serviceData []byte) Interface {
	o.ServiceData = serviceData
	return o
}

// Encode returns string representation of a transaction. RLP-encoded structure in hex format.
func (tx *Transaction) Encode() (string, error) {
	src, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "", err
	}

	return "0x" + hex.EncodeToString(src), nil
}

// Hash returns hash of Transaction.
func (o *object) Hash() (string, error) {
	encode, err := o.Transaction.Encode()
	if err != nil {
		return "", err
	}

	decode, err := hex.DecodeString(encode[2:])
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(decode)
	return "Mt" + hex.EncodeToString(hash[:]), nil
}

func (o *object) addSignature(signatures ...*SignatureSingle) (Signed, error) {
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

// AddSignature adds signature from hex strings.
func (o *object) AddSignature(signatures ...string) (Signed, error) {
	signature, err := o.Signature()
	if err != nil {
		return nil, err
	}

	if len(signatures) == 0 {
		return nil, errors.New("number of signatures must be greater than 0")
	}

	if o.SignatureType == SignatureTypeSingle {
		decode, err := hex.DecodeString(signatures[0])
		if err != nil {
			return nil, err
		}

		sign, err := decodeSignature(decode)
		if err != nil {
			return nil, err
		}

		return o.setSignature(sign)
	}
	if len(o.SignatureData()) == 0 {
		return nil, errors.New("multisig address not set")
	}
	signatureMulti := signature.(*SignatureMulti)
	signs := make([]*SignatureSingle, 0, len(signatures))
	for _, signature := range signatures {
		decode, err := hex.DecodeString(signature)
		if err != nil {
			return nil, err
		}

		sign, err := decodeSignature(decode)
		if err != nil {
			return nil, err
		}

		signs = append(signs, sign)
	}
	signatureMulti.Signatures = append(signatureMulti.Signatures, signs...)
	return o.setSignature(signatureMulti)
}

// Sign signs transaction with a private key.
func (o *object) Sign(key string, multisigPrKeys ...string) (Signed, error) {
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
				Multisig:   Address{},
				Signatures: make([]*SignatureSingle, 0, len(multisigPrKeys)),
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
		signatures := make([]*SignatureSingle, 0, len(multisigPrKeys))
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

func newSignature(prKey string, h PublicKey) (*SignatureSingle, error) {
	sig, err := sign(prKey, h)
	if err != nil {
		return nil, err
	}
	return &SignatureSingle{
		R: new(big.Int).SetBytes(sig[:32]),
		S: new(big.Int).SetBytes(sig[32:64]),
		V: new(big.Int).SetBytes([]byte{sig[64] + 27}),
	}, nil
}

func sign(prKey string, h PublicKey) ([]byte, error) {
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

func rlpHash(x interface{}) (h PublicKey, err error) {
	hw := sha3.NewLegacyKeccak256()
	err = rlp.Encode(hw, x)
	if err != nil {
		return h, err
	}
	hw.Sum(h[:0])
	return h, nil
}
