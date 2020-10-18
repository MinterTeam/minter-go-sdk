package models_test

import (
	"github.com/MinterTeam/minter-go-sdk/v2/api"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
)

func mark(address, publicKey string)  {}
func doSomething(*api.StakeKickEvent) {}

func ExampleConvertStructToEvent() {
	data := map[string]interface{}{
		"type": api.TypeStakeKickEvent,
		"value": map[string]interface{}{
			"address":           "Mx",
			"amount":            "1000000000000",
			"coin":              "1",
			"validator_pub_key": "Mp",
		},
	}
	event, _ := models.ConvertStructToEvent(data)
	mark(event.GetAddress(), event.GetValidatorPublicKey())
	doSomething(event.(*api.StakeKickEvent))
}
