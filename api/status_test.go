package api

import (
	"testing"
)

func TestApi_Status(t *testing.T) {
	response, err := NewApi("https://minter-node-1.testnet.minter.network:8841").Status()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", response)
}
