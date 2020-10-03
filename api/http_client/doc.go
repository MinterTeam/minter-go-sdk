/*Package http_client is the interface for uses API v2 methods.

Example:

	client, _ := http_client.New("http://localhost:8843/v2")
	w, _ := wallet.Create("1 2 3 4 5 6 7 8 9 10 11 12", "")
	data := transaction.NewSendData().SetCoin(0).SetValue(big.NewInt(1)).MustSetTo(w.Address)
	transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
	tx, _ := transactionsBuilder.NewTransaction(data)
	sign, _ := tx.SetNonce(4).SetGasPrice(1).Sign(w.PrivateKey)
	encode, _ := sign.Encode()
	hash, _ := sign.Hash()
	subscribeClient, _ := client.Subscribe(context.Background(), fmt.Sprintf("tx.hash = '%s'", strings.ToUpper(hash[2:])))
	defer subscribeClient.CloseSend()

	res, _ := client.SendTransaction(api_service.NewSendTransactionParams().WithTx(encode))
	if res.Payload.Code != 0 {
		panic(res.Payload.Log)
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
		log.Println("OK", recv.Result)
		break
	}

*/
package http_client
