package models_test

import (
	"encoding/json"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
)

func ExampleProtobufAny_UnmarshalNew() {
	transactionResponse := models.TransactionResponse{
		TypeHex: "0x12",
		Type:    18,
		Data: &models.ProtobufAny{
			"@type":     "type.googleapis.com/api_pb.EditMultisigData",
			"threshold": "5",
			"weights":   []string{"1", "2", "3"},
			"addresses": []string{"Mx0", "Mx1", "Mx2"},
		},
	}
	data, _ := transactionResponse.Data.UnmarshalNew()
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
