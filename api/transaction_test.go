// +build integration

package api

import "testing"

func TestApi_Transaction(t *testing.T) {
	response, err := testApi.Transaction("Mtdecf9373bf68b0052fddee003509b591897c9539f4ee9cc6a32f4ab05c3629cc")
	if err != nil {
		t.Fatal(err)
	}

	dataStruct, err := response.DataStruct()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", dataStruct)
}
