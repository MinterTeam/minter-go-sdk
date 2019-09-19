// +build integration

package api

import (
	"testing"
)

func TestApi_Address(t *testing.T) {
	response, err := testApi.Address("Mxeeee1973381ab793719fff497b9a516719fcd5a2")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
