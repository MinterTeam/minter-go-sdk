// +build integration

package api

import (
	"testing"
)

func TestApi_MissedBlocks(t *testing.T) {
	response, err := testApi.MissedBlocks("Mp1ada5ac409b965623bf6a4320260190038ae27230abfb5ebc9158280cdffffff", 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
