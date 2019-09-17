// +build integration

package api

import "testing"

func TestApi_Transaction(t *testing.T) {
	response, err := testApi.Transaction("")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", response)
}
