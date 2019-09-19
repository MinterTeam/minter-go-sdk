// +build integration

package api

import (
	"testing"
)

func TestApi_Validators(t *testing.T) {
	response, err := testApi.Validators(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
