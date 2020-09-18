/*
	This package is used to call gRPC methods here https://pkg.go.dev/github.com/MinterTeam/node-grpc-gateway/api_pb?tab=doc.

	Example:

		client, _ := grpc_client.New("node.chilinet.minter.network:28842")
		coinSymbol := "SUPERTEST9"
		dataCreateCoin := transaction.NewCreateCoinData().
			SetSymbol(coinSymbol)...
		transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
		tx, _ := transactionsBuilder.NewTransaction(dataCreateCoin)
		sign, _ := tx.SetNonce(1).SetGasPrice(1).Sign(privateKey)
		hash, _ := sign.Hash()
		encode, _ := sign.Encode()

		subscribeClient, _ := client.Subscribe("tm.event = 'Tx'")
		defer subscribeClient.CloseSend()

		res, err := client.SendTransaction(encode)
		if err != nil {
			log.Fatal(client.HttpError(err))
		}

		marshal, _ := client.Marshal(res)

		for {
			recv, err := subscribeClient.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				if statusError, ok := status.FromError(err); ok {
					if statusError.Code() == codes.DeadlineExceeded {
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
			log.Fatal(client.HttpError(err))
		}

		var coinID transaction.CoinID
		coin, _ := txRes.Tags["tx.coin_id"]
		id, _ := strconv.Atoi(coin)
		coinID = transaction.CoinID(id)

		dataSend := transaction.NewSendData().
			SetCoin(coinID)...

		tx, _ = transactionsBuilder.NewTransaction(dataSend)
		sign, _ = tx.SetNonce(2).SetGasPrice(1).Sign(privateKey)
		encode, _ = sign.Encode()

		res, err = client.SendTransaction(encode)
		if err != nil {
			log.Fatal(client.HttpError(err))
		}

*/
package grpc_client
