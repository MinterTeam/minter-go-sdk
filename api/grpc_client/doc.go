package grpc_client

/*
This package is used to call gRPC methods here https://pkg.go.dev/github.com/MinterTeam/node-grpc-gateway@v1.1.1/api_pb?tab=doc


Examples:

		client := grpc_client.New(grpcAddress)
		statusResponse, _ := client.Status()
		fmt.Printf("%#v", statusResponse)
		// Result: &api_pb.StatusResponse{Version:"1.1.6-testnet", LatestBlockHash:"A02D691E2AC87EF1847B1D89D0283D44AE83654A9A1643B9EC9551D5E2D0D647", LatestAppHash:"C5E19AD5E7BC3A77EFB7385CAF2FEE78917628A35104CBC32723D7B2E666C087", LatestBlockHeight:"16491", LatestBlockTime:"2020-04-09T11:27:24.530403396+03:00", KeepLastStates:"120", CatchingUp:false, PublicKey:"Mp0d29a83e54653a1d5f34e561e0135f1e81cbcae152f1f327ab36857a7e32de4c", NodeId:"4735e67924e611b89fbd3f951441b5e912e226d3", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}


		f := func(context.Context) func() context.Context {
			return func() context.Context {
				ctx, _ := context.WithTimeout(context.Context, time.Second)
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
		//&api_pb.SubscribeResponse{Query:"tags.tx.type='01'", Data:(*structpb.Struct)(0xc00016eed0), Events:[]*api_pb.SubscribeResponse_Event{(*api_pb.SubscribeResponse_Event)(0xc00007ec30), (*api_pb.SubscribeResponse_Event)(0xc00007ee10), (*api_pb.SubscribeResponse_Event)(0xc00007ee60), (*api_pb.SubscribeResponse_Event)(0xc00007eeb0), (*api_pb.SubscribeResponse_Event)(0xc00007ef00), (*api_pb.SubscribeResponse_Event)(0xc00007ef50), (*api_pb.SubscribeResponse_Event)(0xc00007efa0)}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//tx -> &structpb.Value{Kind:(*structpb.Value_StringValue)(0xc000061d90), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//result -> &structpb.Value{Kind:(*structpb.Value_StructValue)(0xc000010378), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//height -> &structpb.Value{Kind:(*structpb.Value_NumberValue)(0xc000027690), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//index -> &structpb.Value{Kind:(*structpb.Value_NumberValue)(0xc0000276a0), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tags.tx.type", Events:[]string{"01"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tags.tx.from", Events:[]string{"eb11b60871b2b81c610ee5d28db8b1ca10d898c1"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tags.tx.to", Events:[]string{"eb11b60871b2b81c610ee5d28db8b1ca10d898c1"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tags.tx.coin", Events:[]string{"MNT"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tm.event", Events:[]string{"Tx"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tx.hash", Events:[]string{"8428E3AF832BBD808AED1EDECA072A89B192678C8578089BD27573C1870DD6A5"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tx.height", Events:[]string{"18418"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//
		//&api_pb.SubscribeResponse{Query:"tags.tx.type='01'", Data:(*structpb.Struct)(0xc000226450), Events:[]*api_pb.SubscribeResponse_Event{(*api_pb.SubscribeResponse_Event)(0xc00022c2d0), (*api_pb.SubscribeResponse_Event)(0xc00022c320), (*api_pb.SubscribeResponse_Event)(0xc00022c370), (*api_pb.SubscribeResponse_Event)(0xc00022c3c0), (*api_pb.SubscribeResponse_Event)(0xc00022c410), (*api_pb.SubscribeResponse_Event)(0xc00022c460), (*api_pb.SubscribeResponse_Event)(0xc00022c4b0)}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//height -> &structpb.Value{Kind:(*structpb.Value_NumberValue)(0xc0002240f0), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//index -> &structpb.Value{Kind:(*structpb.Value_NumberValue)(0xc000224100), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//tx -> &structpb.Value{Kind:(*structpb.Value_StringValue)(0xc0002531c0), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//result -> &structpb.Value{Kind:(*structpb.Value_StructValue)(0xc00022e050), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tx.height", Events:[]string{"18434"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tags.tx.type", Events:[]string{"01"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tags.tx.from", Events:[]string{"eb11b60871b2b81c610ee5d28db8b1ca10d898c1"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tags.tx.to", Events:[]string{"a7bd17b15f341ebd38a300a469744f1541f6ffcb"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tags.tx.coin", Events:[]string{"CAPITAL"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tm.event", Events:[]string{"Tx"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
		//&api_pb.SubscribeResponse_Event{Key:"tx.hash", Events:[]string{"002290B942B5685462689DB838F9EEA6AE07C062B48BBC369E4EE0D6E0964A06"}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}

*/
