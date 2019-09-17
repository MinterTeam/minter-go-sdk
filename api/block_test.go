// +build integration

package api

import "testing"

func TestApi_Block(t *testing.T) {
	response, err := testApi.Block(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", response)
}
