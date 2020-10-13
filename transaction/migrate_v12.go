package transaction

import (
	"encoding/json"
	"github.com/MinterTeam/minter-go-sdk/api"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-resty/resty/v2"
	"os"
	"reflect"
	"time"
)

var httpClient *api.Api

func coinID(symbol string) (transaction.CoinID, error) {
	coinID, err := httpClient.CoinInfo(symbol)
	if err != nil {
		return 0, err
	}
	return transaction.CoinID(coinID.ID), nil
}

func (o *object) sign(key string, multisigPrKeys ...string) (SignedTransaction, error) {
	tx := transaction.Object{
		Transaction: &transaction.Transaction{
			Nonce:         o.Nonce,
			ChainID:       transaction.ChainID(o.ChainID),
			GasPrice:      o.GasPrice,
			GasCoin:       0,
			Type:          transaction.Type(o.Type),
			Data:          nil,
			Payload:       o.Payload,
			ServiceData:   o.ServiceData,
			SignatureType: transaction.SignatureType(o.SignatureType),
			SignatureData: nil,
		},
	}

	var err error

	tx.Transaction.GasCoin, err = coinID(o.GasCoin.String())
	if err != nil {
		return nil, err
	}

	tx.Transaction.Data, err = o.EncodeDataReflect()
	if err != nil {
		return nil, err
	}

	signedTx, err := tx.Sign(key, multisigPrKeys...)
	if err != nil {
		return nil, err
	}

	o.encode, err = signedTx.Encode()
	if err != nil {
		return nil, err
	}

	return o, err
}

func (o *object) EncodeData() ([]byte, error) {
	marshal, err := json.Marshal(o.data)
	if err != nil {
		return nil, err
	}
	fields := map[string]interface{}{}
	err = json.Unmarshal(marshal, &fields)
	if err != nil {
		return nil, err
	}
	data := make([]interface{}, 0, len(fields))
	for _, field := range fields {
		if float, ok := field.(float64); ok {
			data = append(data, uint(float))
			continue
		}
		if coin, ok := field.(Coin); ok {
			coinID, err := coinID(coin.String())
			if err != nil {
				return nil, err
			}
			data = append(data, coinID)
			continue
		}
		data = append(data, field)
	}

	return rlp.EncodeToBytes(data)
}

func (o *object) EncodeDataReflect() (bytes []byte, err error) {
	valueOf := reflect.ValueOf(o.data).Elem()
	data := make([]interface{}, 0, valueOf.NumField())
	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i).Interface()
		if coin, ok := field.(Coin); ok {
			field, err = coinID(coin.String())
			if err != nil {
				return bytes, err
			}
		}
		data = append(data, field)
	}

	return rlp.EncodeToBytes(data)
}

func init() {
	nodeApiUrl := os.Getenv("NODE_API")
	if nodeApiUrl == "" {
		panic("set env 'NODE_API'")
	}
	httpClient = api.NewApiWithClient(nodeApiUrl, resty.New().SetTimeout(time.Minute).EnableTrace().SetDebug(true))
}
