// +build integration

package api

import (
	"testing"
)

func TestApi_MaxGas(t *testing.T) {
	response, err := testApi.MaxGas()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
