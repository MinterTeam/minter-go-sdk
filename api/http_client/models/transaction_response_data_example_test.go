package models_test

import (
	"encoding/json"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
)

func ExampleConvertToData() {
	transactionResponse := models.TransactionResponse{
		Type: 18,
		Data: &models.ProtobufAny{
			"threshold": "5",
			"weights":   []string{"1", "2", "3"},
			"addresses": []string{"Mx0", "Mx1", "Mx2"},
		},
	}
	data, _ := models.ConvertToData(transactionResponse.Type, transactionResponse.Data)
	editMultisigData := data.(*models.EditMultisigData)

	fmt.Printf("%T %[1]v\n", editMultisigData.Threshold)
	fmt.Printf("%T %[1]v\n", []uint64(editMultisigData.Weights))
	fmt.Printf("%T %[1]v\n", editMultisigData.Addresses)

	marshal, _ := json.Marshal(editMultisigData)
	fmt.Printf("%s", marshal)

	// Output:
	// uint64 5
	// []uint64 [1 2 3]
	// []string [Mx0 Mx1 Mx2]
	// {"threshold":"5","weights":["1","2","3"],"addresses":["Mx0","Mx1","Mx2"]}

}
