package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type AddLiquidityData struct {
	Coin0          CoinID
	Coin1          CoinID
	Volume0        *big.Int
	MaximumVolume1 *big.Int
}

// Type returns Data type of the transaction.
func (d *AddLiquidityData) Type() Type {
	return TypeAddLiquidity
}

// Fee returns commission of transaction Data
func (d *AddLiquidityData) Fee() Fee {
	return feeTypeAddLiquidityData
}

// Encode returns the byte representation of a transaction Data.
func (d *AddLiquidityData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

type RemoveLiquidityData struct {
	Coin0          CoinID
	Coin1          CoinID
	Liquidity      *big.Int
	MinimumVolume0 *big.Int
	MinimumVolume1 *big.Int
}

// Type returns Data type of the transaction.
func (d *RemoveLiquidityData) Type() Type {
	return TypeRemoveLiquidity
}

// Fee returns commission of transaction Data
func (d *RemoveLiquidityData) Fee() Fee {
	return feeTypeRemoveLiquidity
}

// Encode returns the byte representation of a transaction Data.
func (d *RemoveLiquidityData) Encode() ([]byte, error) {
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
