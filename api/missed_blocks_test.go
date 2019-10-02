// +build integration

package api

import (
	"testing"
)

func TestApi_MissedBlocks(t *testing.T) {
	responseValidators, err := testApi.Validators(0)
	if err != nil {
		t.Fatal(err)
	}
	if len(responseValidators) == 0 {
		t.Fatal("no candidates")
	}
	response, err := testApi.MissedBlocks(responseValidators[0].PubKey, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
