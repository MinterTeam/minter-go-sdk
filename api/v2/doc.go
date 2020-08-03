/*
	ClientService is the interface for uses API v2 methods.

	Example:

		import (
			v2 "github.com/MinterTeam/minter-go-sdk/v2/api/v2"
			"github.com/MinterTeam/minter-go-sdk/v2/api/v2/client/api_service"
			"github.com/MinterTeam/minter-go-sdk/v2/transaction"
			"github.com/MinterTeam/minter-go-sdk/v2/wallet"
			"log"
			"math/big"
		)

		func Example() {
			apiv2, err := v2.New("http://68.183.211.176:8843/")
			if err != nil {
				panic(err)
			}

			data := transaction.NewSendData().MustSetTo("Mxd82558ea00eb81d35f2654953598f5d51737d31d").SetCoin(0).SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))
			newTransaction, err := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
			if err != nil {
				panic(err)
			}

			seed, err := wallet.Seed("a b c d e f g h 1 2 3 4")
			if err != nil {
				panic(err)
			}

			prKey, err := wallet.PrivateKeyBySeed(seed)
			if err != nil {
				panic(err)
			}

			sign, err := newTransaction.SetGasCoin(0).SetGasPrice(1).SetNonce(156).Sign(prKey)
			if err != nil {
				panic(err)
			}

			encode, err := sign.Encode()
			if err != nil {
				panic(err)
			}

			log.Println(encode)

			res, err := apiv2.APIServiceSendTransaction(api_service.NewAPIServiceSendTransactionParams().WithTx(encode))
			if err != nil {
				runtimeError, ok := err.(*api_service.APIServiceSendTransactionDefault)
				if !ok {
					panic(err)
				}

				log.Println(runtimeError.Payload.Message)
				log.Println(runtimeError.Payload.Error)
				log.Println(runtimeError.Payload.Code)
				binary, err := runtimeError.GetPayload().MarshalBinary()
				if err != nil {
					panic(err)
				}
				log.Println(string(binary))
				return
			}
		}

*/
package v2
