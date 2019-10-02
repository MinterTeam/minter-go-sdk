// +build integration

package api

import (
	"testing"
)

func TestApi_CoinInfo(t *testing.T) {
	response, err := testApi.CoinInfo("CAPITAL", 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
