// +build integration

package api

import (
	"testing"
)

func TestApi_GetAddress(t *testing.T) {
	response, err := NewApi("https://minter-node-1.testnet.minter.network:8841").
		Address([]byte("Mxeeee1973381ab793719fff497b9a516719fcd5a2"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", response)
}
