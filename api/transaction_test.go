// +build integration

package api

import "testing"

func TestApi_Transaction(t *testing.T) {
	response, err := testApi.Transaction("Mt708c2019938339aba4bf6c2f771373bc43e0efa7df65c187950964321734cd82")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
