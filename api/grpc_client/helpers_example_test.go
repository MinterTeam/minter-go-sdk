package grpc_client_test

import (
	"github.com/MinterTeam/minter-go-sdk/v2/api"
	"github.com/MinterTeam/minter-go-sdk/v2/api/grpc_client"
)

func mark(address, publicKey string) {}

func doSomething1(*api.RewardEvent)    {}
func doSomething2(*api.SlashEvent)     {}
func doSomething3(*api.UnbondEvent)    {}
func doSomething4(*api.StakeKickEvent) {}

func ExampleConvertStructToEvent() {
	client, _ := grpc_client.New("localhost:8842")
	events, _ := client.Events(2280)
	for _, data := range events.Events {
		event, _ := grpc_client.ConvertStructToEvent(data)

		mark(event.GetAddress(), event.GetValidatorPublicKey())

		switch e := event.(type) {
		case *api.RewardEvent:
			doSomething1(e)
		case *api.SlashEvent:
			doSomething2(e)
		case *api.UnbondEvent:
			doSomething3(e)
		case *api.StakeKickEvent:
			doSomething4(e)
		}

	}
}
