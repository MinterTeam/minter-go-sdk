// +build integration

package api

import (
	"testing"
)

func TestApi_Events(t *testing.T) {
	response, err := testApi.Events(12)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range response.Events {
		t.Run(v.Type, func(t *testing.T) {
			data, err := v.ValueStruct()
			if err != nil {
				t.Error(err)
			}
			var ok bool
			switch v.Type {
			case "minter/RewardEvent":
				_, ok = data.(*RewardEventValue)
			case "minter/SlashEvent":
				_, ok = data.(*SlashEventValue)
			case "minter/UnbondEvent":
				_, ok = data.(*UnbondEventValue)
			case "minter/CoinLiquidationEvent":
				_, ok = data.(*CoinLiquidationEventValue)
			default:
				t.Fatal("not found interface by type")
			}
			if !ok {
				t.Fatalf("interface conversion: interface {} is %T", data)
			}
		})
	}
	t.Logf("%+v", response)
}
