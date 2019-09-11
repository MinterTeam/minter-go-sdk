package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type SendTransaction struct {
	*Transaction
}

func NewSendTransaction() *SendTransaction {
	return &SendTransaction{Transaction: &Transaction{SignatureType: signatureTypeSingle, ChainID: TestNetChainID}}
}

type SendData struct {
	Coin  [10]byte //todo
	To    [20]byte
	Value *big.Int
}

func NewSendData(coin string, to string, value *big.Int) *SendData {
	data := &SendData{Value: value}
	copy(data.Coin[:], coin)
	copy(data.To[:], to)
	return data
}

func (d *SendData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

func (tx *SendTransaction) SetData(data Data) {
	_, ok := data.(*SendData)
	if !ok {
		panic("") //todo
	}
	tx.Data, _ = data.encode()
}
