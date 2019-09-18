// +build integration

package api

import (
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"strconv"
	"testing"
)

func TestApi_Block(t *testing.T) {
	response, err := testApi.Block(19)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range response.Result.Transactions {
		t.Run(strconv.Itoa(v.Type), func(t *testing.T) {
			data, err := v.DataStruct()
			if err != nil {
				t.Error(err)
			}
			var ok bool
			switch transaction.Type(v.Type) {
			case transaction.TypeSend:
				_, ok = data.(*SendData)
			case transaction.TypeSellCoin:
				_, ok = data.(*SellCoinData)
			case transaction.TypeSellAllCoin:
				_, ok = data.(*SellAllCoinData)
			case transaction.TypeBuyCoin:
				_, ok = data.(*SellCoinData)
			case transaction.TypeCreateCoin:
				_, ok = data.(*BuyCoinData)
			case transaction.TypeDeclareCandidacy:
				_, ok = data.(*DeclareCandidacyData)
			case transaction.TypeDelegate:
				_, ok = data.(*DelegateData)
			case transaction.TypeUnbond:
				_, ok = data.(*UnbondData)
			case transaction.TypeRedeemCheck:
				_, ok = data.(*RedeemCheckData)
			case transaction.TypeSetCandidateOnline:
				_, ok = data.(*SetCandidateOnData)
			case transaction.TypeSetCandidateOffline:
				_, ok = data.(*SetCandidateOffData)
			case transaction.TypeCreateMultisig:
				_, ok = data.(*CreateMultisigData)
			case transaction.TypeMultisend:
				_, ok = data.(*MultisendData)
			case transaction.TypeEditCandidate:
				_, ok = data.(*EditCandidateData)
			default:
				t.Fatal("not found interface by type")
			}
			if !ok {
				t.Fatalf("interface conversion: interface {} is %T", data)
			}
		})
	}
}
