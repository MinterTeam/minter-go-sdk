package v2_test

import (
	v2 "github.com/MinterTeam/minter-go-sdk/api/v2"
	"github.com/klim0v/node-api-v2-client-go/client/api_service"
)

func ExampleNew() {
	apiv2, _ := v2.New("https://minter-node-1.testnet.minter.network/v2/")

	status, _ := apiv2.APIServiceStatus(api_service.NewAPIServiceStatusParams())

	status.GetPayload()
}
