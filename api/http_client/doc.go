/*
	ClientService is the interface for uses API v2 methods.

	Example:

		client, _ := http_client.New("http://localhost:8843")
		coinSymbol := "SUPERTEST9"
		dataCreateCoin := transaction.NewCreateCoinData().
			SetSymbol(coinSymbol)...
		transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
		tx, _ := transactionsBuilder.NewTransaction(dataCreateCoin)
		sign, _ := tx.SetNonce(1).SetGasPrice(1).Sign(privateKey)
		encode, _ := sign.Encode()

		res, err := client.SendTransaction(api_service.NewSendTransactionParams().WithTx(encode))
		if err != nil {
			log.Fatal(http_client.ErrorBody(err))
		}

		txRes, err := client.Transaction(api_service.NewTransactionParams().WithHash(res.GetPayload().Hash))
		if err != nil {
			log.Fatal(http_client.ErrorBody(err))
		}

		var coinID transaction.CoinID
		coin, _ := txRes.GetPayload().Tags["tx.coin_id"]
		id, _ := strconv.Atoi(coin)
		coinID = transaction.CoinID(id)

		dataSend := transaction.NewSendData().
			SetCoin(coinID)...

		tx, _ = transactionsBuilder.NewTransaction(dataSend)
		sign, _ = tx.SetNonce(2).SetGasPrice(1).Sign(privateKey)
		encode, _ = sign.Encode()

		res, err = client.SendTransaction(api_service.NewSendTransactionParams().WithTx(encode))
		if err != nil {
			log.Fatal(http_client.ErrorBody(err))
		}

*/
package http_client
