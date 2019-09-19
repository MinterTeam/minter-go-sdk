// +build integration

package api

import (
	"testing"
)

func TestApi_UnconfirmedTxs(t *testing.T) {
	response, err := testApi.UnconfirmedTxs(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
