// +build integration

package grpc_client_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/api/grpc_client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const grpcAddressTestnet = "minter-node-1.testnet.minter.network:8842"

func ExampleClient_Status() {
	statusResponse, err := grpc_client.New(grpcAddressTestnet).Status()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v", statusResponse)
	// Output: &api_pb.StatusResponse{Version:"1.1.6-testnet", LatestBlockHash:"A02D691E2AC87EF1847B1D89D0283D44AE83654A9A1643B9EC9551D5E2D0D647", LatestAppHash:"C5E19AD5E7BC3A77EFB7385CAF2FEE78917628A35104CBC32723D7B2E666C087", LatestBlockHeight:"16491", LatestBlockTime:"2020-04-09T11:27:24.530403396+03:00", KeepLastStates:"120", CatchingUp:false, PublicKey:"Mp0d29a83e54653a1d5f34e561e0135f1e81cbcae152f1f327ab36857a7e32de4c", NodeId:"4735e67924e611b89fbd3f951441b5e912e226d3", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
}

func ExampleClient_CoinInfo_exast_coin() {
	statusResponse, err := grpc_client.New(grpcAddressTestnet).CoinInfo("CAPITAL")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v", statusResponse)
	// Output: &api_pb.CoinInfoResponse{Name:"CAPITAL", Symbol:"CAPITAL", Volume:"10004406818749673124164086", Crr:"80", ReserveBalance:"100055088268419724781795", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
}

func ExampleClient_CoinInfo_coin_is_not_found() {
	coinInfoResponse, err := grpc_client.New(grpcAddressTestnet).CoinInfo("KLIM")
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
	// Output:
	//rpc error: code = FailedPrecondition desc = Coin not found
	//Coin not found
	//[]
	//code:9 message:"Coin not found"
}

func ExampleClient_Address_empty_balance_at_height() {
	addressResponse, err := grpc_client.New(grpcAddressTestnet).Address("Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c", 17262)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%#v", addressResponse)
	// Output: &api_pb.AddressResponse{Balance:map[string]string{"MNT":"0"}, TransactionsCount:"0", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
}

func ExampleClient_Address() {
	addressResponse, err := grpc_client.New(grpcAddressTestnet).Address("Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%#v", addressResponse)
	// Output: &api_pb.AddressResponse{Balance:map[string]string{"MNT":"15000000000000000000000"}, TransactionsCount:"0", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
}

func ExampleClient_Addresses() {
	addressResponse, err := grpc_client.New(grpcAddressTestnet).Addresses([]string{"Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c", "Mxa7bd17b15f341ebd38a300a469744f1541f6ffcb"}, 17520)
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, address := range addressResponse.Addresses {
		fmt.Printf("%#v\n", address)
	}
	// Output:
	//&api_pb.AddressesResponse_Result{Address:"Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c", Balance:map[string]string{"MNT":"15000000000000000000000"}, TransactionsCount:"0", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.AddressesResponse_Result{Address:"Mxa7bd17b15f341ebd38a300a469744f1541f6ffcb", Balance:map[string]string{"MNT":"5299412918045840693417561"}, TransactionsCount:"703", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
}

func ExampleClient_Block() {
	blockResponse, err := grpc_client.New(grpcAddressTestnet).Block(17500)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%#v\n", blockResponse)

	for _, validator := range blockResponse.Validators {
		fmt.Printf("%#v\n", validator)
	}
	for _, transaction := range blockResponse.Transactions {
		fmt.Printf("%#v\n", transaction)
	}
	// Output:
	//&api_pb.BlockResponse{Hash:"f3b9c0c2bd4f535b0077a46c66e05e0ec651bfd684722418db80c84a23a3a024", Height:"17500", Time:"2020-04-09T09:38:19.705407393Z", TransactionsCount:"1", Transactions:[]*api_pb.BlockResponse_Transaction{(*api_pb.BlockResponse_Transaction)(0xc000212100)}, BlockReward:"330000000000000000000", Size:"759", Proposer:"Mp0d29a83e54653a1d5f34e561e0135f1e81cbcae152f1f327ab36857a7e32de4c", Validators:[]*api_pb.BlockResponse_Validator{(*api_pb.BlockResponse_Validator)(0xc000020480)}, Evidence:(*api_pb.BlockResponse_Evidence)(0xc0000204c0), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.BlockResponse_Validator{PublicKey:"Mp0d29a83e54653a1d5f34e561e0135f1e81cbcae152f1f327ab36857a7e32de4c", Signed:true, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
	//&api_pb.BlockResponse_Transaction{Hash:"Mtd98a3b970cced396bad3c70ecc6cccb5881ae84ca503c53cda8b3093e8f7be9f", RawTx:"f8888202bf02018a4d4e540000000000000001aceb8a4d4e5400000000000000944157ff22db9fd8ea0dc955e4d78007ba9f2774b68a032d26d12e980b600000808001b845f8431ba01be01c94303de43562b85140f2659a7450be8fa256b37912cde4e07850547026a037edb83c3fe15b2cc45ae1cb2be6a133518f28589ac0583dde9871a412ae7af3", From:"Mxa7bd17b15f341ebd38a300a469744f1541f6ffcb", Nonce:"703", GasPrice:"1", Type:"1", Data:(*structpb.Struct)(0xc00020e3f0), Payload:[]uint8(nil), ServiceData:[]uint8(nil), Gas:"10", GasCoin:"MNT", Tags:map[string]string{"tx.coin":"MNT", "tx.from":"a7bd17b15f341ebd38a300a469744f1541f6ffcb", "tx.to":"4157ff22db9fd8ea0dc955e4d78007ba9f2774b6", "tx.type":"01"}, Code:"0", Log:"", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
}

func ExampleClient_Subscribe() {
	subscribeClient, err := grpc_client.New(grpcAddressTestnet).Subscribe("tags.tx.type='01'")
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
	// Output:
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
}
