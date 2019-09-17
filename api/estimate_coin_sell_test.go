// +build integration

package api

import "testing"

func TestApi_EstimateCoinSell(t *testing.T) {
	response, err := testApi.EstimateCoinSell("BIP", "1", "MNT", 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", response)
}
