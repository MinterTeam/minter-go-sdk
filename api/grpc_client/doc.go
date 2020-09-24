/*Package grpc_client is the interface for uses gRPC methods.

Example:

	client, _ := grpc_client.New("localhost:8842")
	coinSymbol := "SUPERTEST9"
	dataCreateCoin := transaction.NewCreateCoinData().
		SetSymbol(coinSymbol) // ...
	transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
	tx, _ := transactionsBuilder.NewTransaction(dataCreateCoin)
	sign, _ := tx.SetNonce(1).SetGasPrice(1).Sign(privateKey)
	hash, _ := sign.Hash()
	encode, _ := sign.Encode()

	subscribeClient, _ := client.Subscribe("tm.event = 'Tx'")
	defer subscribeClient.CloseSend()

	res, err := client.SendTransaction(encode)
		if err != nil {
			log.Fatal(client.HTTPError(err))
		}

	marshal, _ := client.Marshal(res)

	for {
		recv, err := subscribeClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			if statusError, ok := status.FromError(err); ok {
				if statusError.Code() == codes.DeadlineExceeded || statusError.Code() == codes.Canceled {
					break
				}
			}
			log.Fatal(err)
		}

		marshal, _ := client.Marshal(recv)

		if !strings.Contains(marshal, strings.ToUpper(hash[2:])) {
			continue
			}

		break
	}

	txRes, err := client.Transaction(res.Hash)
	if err != nil {
		log.Fatal(client.HTTPError(err))
	}

	coin, _ := txRes.Tags["tx.coin_id"]
	newCoinID, _ := strconv.Atoi(coin)
	mntID, _ := client.CoinID("MNT")
	dataSend := transaction.NewSellAllCoinData().
		SetCoinToSell(transaction.CoinID(newCoinID)).
		SetCoinToSell(transaction.CoinID(mntID)) // ...

	tx, _ = transactionsBuilder.NewTransaction(dataSend)
	sign, _ = tx.SetNonce(2).SetGasPrice(1).Sign(privateKey)
	encode, _ = sign.Encode()

	res, err = client.SendTransaction(encode)
	if err != nil {
		log.Fatal(client.HTTPError(err))
	}

*/
package grpc_client
