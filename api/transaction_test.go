// +build integration

package api

import "testing"

func TestApi_Transaction(t *testing.T) {
	response, err := testApi.Transaction("Mtdd7181de1397eed1513928bb4463cf43fc3cbb5d5056018c063b02113517eca0")
	if err != nil {
		t.Fatal(err)
	}

	_, err = response.DataStruct()
	if err != nil {
		t.Fatal(err)
	}

	var dataStruct SendData
	err = response.Data.FillStruct(&dataStruct)
	if err != nil {
		t.Fatal(err)
	}

}
