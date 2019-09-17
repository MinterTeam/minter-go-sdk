// +build integration

package api

import (
	"testing"
)

func TestApi_MissedBlocks(t *testing.T) {
	response, err := testApi.MissedBlocks("", 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", response)
}
