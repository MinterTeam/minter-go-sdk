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
		for s, value := range recv.Data.Fields {
			fmt.Printf("%s -> %#v\n", s, value)
		}
		for _, event := range recv.Events {
			fmt.Printf("%#v\n", event)
		}
		fmt.Println()
	}
	// Result:
	//&api_pb.SubscribeResponse{Query:"tags.tx.type='01'", Data:(*structpb.Struct)(0xc000202060), Events:[]*api_pb.SubscribeResponse_Event{(*api_pb.SubscribeResponse_Event)(0xc000208500), (*api_pb.SubscribeResponse_Event)(0xc0002086e0), (*api_pb.SubscribeResponse_Event)(0xc000208730), (*api_pb.SubscribeResponse_Event)(0xc000208780), (*api_pb.SubscribeResponse_Event)(0xc0002087d0), (*api_pb.SubscribeResponse_Event)(0xc000208820), (*api_pb.SubscribeResponse_Event)(0xc000208870)}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//height -> &structpb.Value{Kind:(*structpb.Value_NumberValue)(0xc0002060d8), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//index -> &structpb.Value{Kind:(*structpb.Value_NumberValue)(0xc0002060e8), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//tx -> &structpb.Value{Kind:(*structpb.Value_StringValue)(0xc00020e070), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//result -> &structpb.Value{Kind:(*structpb.Value_StructValue)(0xc000210048), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tags.tx.coin", Events:[]string{"MNT"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tm.event", Events:[]string{"Tx"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tx.hash", Events:[]string{"01EFD8EEF507A5BFC4A7D57ECA6F61B96B7CDFF559698639A6733D25E2553539"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tx.height", Events:[]string{"23326"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tags.tx.type", Events:[]string{"01"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tags.tx.from", Events:[]string{"eb11b60871b2b81c610ee5d28db8b1ca10d898c1"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tags.tx.to", Events:[]string{"eb11b60871b2b81c610ee5d28db8b1ca10d898c1"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//
	//&api_pb.SubscribeResponse{Query:"tags.tx.type='01'", Data:(*structpb.Struct)(0xc0002028d0), Events:[]*api_pb.SubscribeResponse_Event{(*api_pb.SubscribeResponse_Event)(0xc000208b90), (*api_pb.SubscribeResponse_Event)(0xc000208be0), (*api_pb.SubscribeResponse_Event)(0xc000208c30), (*api_pb.SubscribeResponse_Event)(0xc000208c80), (*api_pb.SubscribeResponse_Event)(0xc000208cd0), (*api_pb.SubscribeResponse_Event)(0xc000208d20), (*api_pb.SubscribeResponse_Event)(0xc000208d70)}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//tx -> &structpb.Value{Kind:(*structpb.Value_StringValue)(0xc00020e3c0), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//result -> &structpb.Value{Kind:(*structpb.Value_StructValue)(0xc000210118), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//height -> &structpb.Value{Kind:(*structpb.Value_NumberValue)(0xc000206550), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//index -> &structpb.Value{Kind:(*structpb.Value_NumberValue)(0xc000206560), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tags.tx.type", Events:[]string{"01"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tags.tx.from", Events:[]string{"eb11b60871b2b81c610ee5d28db8b1ca10d898c1"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tags.tx.to", Events:[]string{"eb11b60871b2b81c610ee5d28db8b1ca10d898c1"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tags.tx.coin", Events:[]string{"CAPITAL"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tm.event", Events:[]string{"Tx"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tx.hash", Events:[]string{"60780B4E54982C645B3D613FC83203C122CF3EF9FBD9394E564F8AD2E79469F7"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.SubscribeResponse_Event{Key:"tx.height", Events:[]string{"23328"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}

*/

package grpc_client
