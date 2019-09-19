// +build integration

package api

import "testing"

func TestApi_Candidates(t *testing.T) {
	response, err := testApi.Candidates(0, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
