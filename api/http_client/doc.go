/*Package http_client is the interface for uses API v2 methods.

Example:

	client, _ := http_client.New("http://localhost:8843/v2")
	coinSymbol := "SUPERTEST9"
	dataCreateCoin := transaction.NewCreateCoinData().
		SetSymbol(coinSymbol) // ...
	transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
	tx, _ := transactionsBuilder.NewTransaction(dataCreateCoin)
	sign, _ := tx.SetNonce(1).SetGasPrice(1).Sign(privateKey)
	hash, _ := sign.Hash()
	encode, _ := sign.Encode()

	res, err := client.SendTransaction(api_service.NewSendTransactionParams().WithTx(encode))
	if err != nil {
		log.Fatal(http_client.ErrorBody(err))
	}

	subscribe, err := client.Subscribe(context.Background(), "tm.event = 'Tx'")
	if err != nil {
		panic(err)
	}

	for {
		body, err := subscribe.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			subscribe.Close()
			return
		}
		data, _ := json.Marshal(body.Result.Data)
		if strings.Contains(fmt.Sprintf("%s", data), strings.ToUpper(hash[2:])) {
			break
		}
	}

	txRes, err := client.Transaction(api_service.NewTransactionParams().WithHash(res.GetPayload().Hash))
	if err != nil {
		log.Fatal(http_client.ErrorBody(err))
	}

	coin, _ := txRes.GetPayload().Tags["tx.coin_id"]
	newCoinID, _ := strconv.Atoi(coin)
	mntID, _ := client.CoinID("MNT")

	dataSell := transaction.NewSellAllCoinData().
		SetCoinToBuy(uint64(newCoinID)).
		SetCoinToSell(mntID) // ...

	tx, _ = transactionsBuilder.NewTransaction(dataSell)
	sign, _ = tx.SetNonce(2).SetGasPrice(1).Sign(privateKey)
	encode, _ = sign.Encode()

	res, err = client.SendTransaction(api_service.NewSendTransactionParams().WithTx(encode))
	if err != nil {
		log.Fatal(http_client.ErrorBody(err))
	}

*/
package http_client
