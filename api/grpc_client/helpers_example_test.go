package grpc_client_test

import (
	"github.com/MinterTeam/minter-go-sdk/v2/api/grpc_client"
)

func ExampleConvertStructToEvent() {
	client, _ := grpc_client.New("localhost:8842")
	events, _ := client.Events(2280)
	for _, data := range events.Events {
		event, _ := grpc_client.ConvertStructToEvent(data)

		_ = event.GetAddress()
		_ = event.GetValidatorPublicKey()
	}
}
