// +build integration

package api

import "testing"

func TestApi_Transactions(t *testing.T) {
	response, err := testApi.Transactions("tags.tx.coin='CAPITAL'", 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
