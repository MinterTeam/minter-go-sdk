// +build integration

package api

import "testing"

func TestApi_EstimateCoinSell(t *testing.T) {
	response, err := testApi.EstimateCoinSell("BIP", "1", "MNT", LatestBlockHeight)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
