// +build integration

package api

import (
	"testing"
)

func TestApi_CoinInfo(t *testing.T) {
	response, err := testApi.CoinInfoAtHeight("CAPITAL", LatestBlockHeight)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
