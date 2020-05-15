/*
	This package is used to call gRPC methods here https://pkg.go.dev/github.com/MinterTeam/node-grpc-gateway@v1.1.1/api_pb?tab=doc.


	Examples

		client := grpc_client.New(grpcAddress)
		statusResponse, _ := client.Status()
		fmt.Printf("%#v", statusResponse)
		// Result: &api_pb.StatusResponse{Version:"1.1.6-testnet", LatestBlockHash:"A02D691E2AC87EF1847B1D89D0283D44AE83654A9A1643B9EC9551D5E2D0D647", LatestAppHash:"C5E19AD5E7BC3A77EFB7385CAF2FEE78917628A35104CBC32723D7B2E666C087", LatestBlockHeight:"16491", LatestBlockTime:"2020-04-09T11:27:24.530403396+03:00", KeepLastStates:"120", CatchingUp:false, PublicKey:"Mp0d29a83e54653a1d5f34e561e0135f1e81cbcae152f1f327ab36857a7e32de4c", NodeId:"4735e67924e611b89fbd3f951441b5e912e226d3", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}


		f := func(c context.Context) func() context.Context {
			return func() context.Context {
				ctx, _ := context.WithTimeout(c, time.Second)
				return ctx
			}
		}
		_, err := client.WithContextFunc(f).Genesis()
		fmt.Println(err)
		// Result: rpc error: code = DeadlineExceeded desc = context deadline exceeded


		coinInfoResponse, err := client.CoinInfo("MNT")
		if err == nil {
			fmt.Printf("want error: \"not found\" , got #%v", coinInfoResponse)
			return
		}

		grpcErr, ok := status.FromError(err)
		if !ok {
			fmt.Printf("want error type: \"GRPC Status\"")
			return
		}

		if grpcErr.Code() != codes.FailedPrecondition {
			fmt.Printf("want grpc code: \"FailedPrecondition\", got %s", grpcErr.Code().String())
			return
		}

		fmt.Printf("%v\n", grpcErr.Err())
		fmt.Printf("%v\n", grpcErr.Message())
		fmt.Printf("%v\n", grpcErr.Details())
		fmt.Printf("%v\n", grpcErr.Proto())
		// Result:
		//rpc error: code = FailedPrecondition desc = Coin not found
		//Coin not found
		//[]
		//code:9 message:"Coin not found"


		subscribeClient, err := client.Subscribe("tags.tx.type='01'")
		if err != nil {
			fmt.Print(err)
			return
		}
		for i := 0; i < 2; i++ {
			recv, err := subscribeClient.Recv()
			if err != nil {
				fmt.Print(err)
				return
			}

			fmt.Printf("%#v\n", recv)
		}

*/

package grpc_client
