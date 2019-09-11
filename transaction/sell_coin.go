package transaction

import "math/big"

type SellCoinTransaction struct {
	Transaction
}

type SellCoinData struct {
	CoinToSell        [10]byte
	ValueToSell       *big.Int
	CoinToBuy         [10]byte
	MinimumValueToBuy *big.Int
}
