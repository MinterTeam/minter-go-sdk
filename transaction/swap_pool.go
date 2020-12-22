package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type AddSwapPoolData struct {
	Coin0          CoinID
	Coin1          CoinID
	Volume0        *big.Int
	MaximumVolume1 *big.Int
}

// Type returns Data type of the transaction.
func (d *AddSwapPoolData) Type() Type {
	return TypeAddSwapPool
}

// Fee returns commission of transaction Data
func (d *AddSwapPoolData) Fee() Fee {
	return feeTypeAddSwapPoolData
}

// Encode returns the byte representation of a transaction Data.
func (d *AddSwapPoolData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

type RemoveSwapPoolData struct {
	Coin0          CoinID
	Coin1          CoinID
	Liquidity      *big.Int
	MinimumVolume0 *big.Int
	MinimumVolume1 *big.Int
}

// Type returns Data type of the transaction.
func (d *RemoveSwapPoolData) Type() Type {
	return TypeRemoveSwapPool
}

// Fee returns commission of transaction Data
func (d *RemoveSwapPoolData) Fee() Fee {
	return feeTypeRemoveSwapPoolData
}

// Encode returns the byte representation of a transaction Data.
func (d *RemoveSwapPoolData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

type BuySwapPoolData struct {
	CoinToBuy          CoinID
	ValueToBuy         *big.Int
	CoinToSell         CoinID
	MaximumValueToSell *big.Int
}

// Type returns Data type of the transaction.
func (d *BuySwapPoolData) Type() Type {
	return TypeBuySwapPool
}

// Fee returns commission of transaction Data
func (d *BuySwapPoolData) Fee() Fee {
	return feeTypeBuyCoin
}

// Encode returns the byte representation of a transaction Data.
func (d *BuySwapPoolData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

type SellAllSwapPoolData struct {
	CoinToSell        CoinID
	CoinToBuy         CoinID
	MinimumValueToBuy *big.Int
}

// Type returns Data type of the transaction.
func (d *SellAllSwapPoolData) Type() Type {
	return TypeSellAllSwapPool
}

// Fee returns commission of transaction Data
func (d *SellAllSwapPoolData) Fee() Fee {
	return feeTypeSellAllCoin
}

// Encode returns the byte representation of a transaction Data.
func (d *SellAllSwapPoolData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

type SellSwapPoolData struct {
	CoinToSell        CoinID
	ValueToSell       *big.Int
	CoinToBuy         CoinID
	MinimumValueToBuy *big.Int
}

// Type returns Data type of the transaction.
func (d *SellSwapPoolData) Type() Type {
	return TypeSellSwapPool
}

// Fee returns commission of transaction Data
func (d *SellSwapPoolData) Fee() Fee {
	return feeTypeSellCoin
}

// Encode returns the byte representation of a transaction Data.
func (d *SellSwapPoolData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
