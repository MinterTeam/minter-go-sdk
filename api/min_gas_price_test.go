// +build integration

package api

import (
	"testing"
)

func TestApi_MinGasPrice(t *testing.T) {
	response, err := testApi.MinGasPrice()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", response)
}
