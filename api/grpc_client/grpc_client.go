package grpc_client

import (
	"github.com/MinterTeam/minter-go-sdk/v2/api/grpc_client"
)

func New(address string) *grpc_client.Client {
	client, err := grpc_client.New(address)
	if err != nil {
		panic(err)
	}
	return client
}
