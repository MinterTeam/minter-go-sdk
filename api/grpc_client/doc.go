/*Package grpc_client is the interface for uses gRPC methods.

Example:

	client, _ := grpc_client.New("localhost:8842")
	w, _ := wallet.Create("1 2 3 4 5 6 7 8 9 10 11 12", "")
	data := transaction.NewSendData().SetCoin(0).SetValue(big.NewInt(1)).MustSetTo(w.Address)
	transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
	tx, _ := transactionsBuilder.NewTransaction(data)
	sign, _ := tx.SetNonce(1).SetGasPrice(1).Sign(w.PrivateKey)
	encode, _ := sign.Encode()
	hash, _ := sign.Hash()
	subscribeClient, _ := client.Subscribe(fmt.Sprintf("tx.hash = '%s'", strings.ToUpper(hash[2:])))
	defer subscribeClient.CloseSend()

	res, _ := client.SendTransaction(encode)
	if res.Code != 0 {
		panic(res.Log)
	}

	for {
		recv, err := subscribeClient.Recv()
		if err == io.EOF {
			break
		}
		if code := status.Code(err); code != codes.OK {
			if code == codes.DeadlineExceeded || code == codes.Canceled {
				break
			}
			panic(err)
		}
		marshal, _ := client.Marshal(recv)
		log.Println("OK", marshal)
		break
	}

*/
package grpc_client
