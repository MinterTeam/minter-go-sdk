// +build integration

package grpc_client

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

const address = "minter-node-1.testnet.minter.network:8842"

var testClient = New(address)

func TestClient_Status(t *testing.T) {
	statusResponse, err := testClient.Status()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", statusResponse)
}

func TestClient_CoinInfo(t *testing.T) {
	statusResponse, err := testClient.CoinInfo("CAPITAL")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", statusResponse)
}

func TestClient_CoinInfoNotFound(t *testing.T) {
	coinInfoResponse, err := testClient.CoinInfo("KLIM")
	if err == nil {
		t.Fatalf("want error: \"not found\" , got #%v", coinInfoResponse)
	}

	grpcErr, ok := status.FromError(err)
	if !ok {
		t.Fatalf("want error type: \"GRPC Status\"")
	}

	if grpcErr.Code() != codes.FailedPrecondition {
		t.Fatalf("want grpc code: \"FailedPrecondition\", got %s", grpcErr.Code().String())
	}

	t.Log(grpcErr.Err())
	t.Log(grpcErr.Message())
	t.Log(grpcErr.Details())
	t.Log(grpcErr.Proto())
}
