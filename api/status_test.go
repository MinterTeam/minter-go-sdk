// +build integration

package api

import (
	"testing"
)

func TestApi_Status(t *testing.T) {
	response, err := testApi.Status()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
