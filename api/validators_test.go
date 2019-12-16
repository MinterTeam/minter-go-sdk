// +build integration

package api

import (
	"testing"
)

func TestApi_Validators(t *testing.T) {
	response, err := testApi.ValidatorsAtHeight(0)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range response {
		t.Logf("%+v", v)
	}
}
