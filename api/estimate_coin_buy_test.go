// +build integration

package api

import (
	"testing"
)

func TestApi_EstimateCoinBuy(t *testing.T) {
	response, err := testApi.EstimateCoinBuyAtHeight("BIP", "1", "MNT", LatestBlockHeight)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
