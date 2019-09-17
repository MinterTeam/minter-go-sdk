// +build integration

package api

import (
	"testing"
)

func TestApi_EstimateCoinBuy(t *testing.T) {
	response, err := testApi.EstimateCoinBuy("BIP", "1", "MNT", 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", response)
}
