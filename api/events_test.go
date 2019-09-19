// +build integration

package api

import "testing"

func TestApi_Events(t *testing.T) {
	response, err := testApi.Events(12)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
