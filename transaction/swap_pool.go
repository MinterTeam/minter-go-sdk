package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type CreateSwapPoolData struct {
	Coin0   CoinID
	Coin1   CoinID
	Volume0 *big.Int
	Volume1 *big.Int
}

func NewCreateSwapPoolData() *CreateSwapPoolData {
	return &CreateSwapPoolData{}
}

func (d *CreateSwapPoolData) SetCoin0(id uint64) *CreateSwapPoolData {
	d.Coin0 = CoinID(id)
	return d
}
func (d *CreateSwapPoolData) SetCoin1(id uint64) *CreateSwapPoolData {
	d.Coin1 = CoinID(id)
	return d
}

func (d *CreateSwapPoolData) SetVolume0(value0 *big.Int) *CreateSwapPoolData {
	d.Volume0 = value0
	return d
}
func (d *CreateSwapPoolData) SetVolume1(value1 *big.Int) *CreateSwapPoolData {
	d.Volume1 = value1
	return d
}

// Type returns Data type of the transaction.
func (d *CreateSwapPoolData) Type() Type {
	return TypeCreateSwapPool
}

// Fee returns commission of transaction Data
func (d *CreateSwapPoolData) Fee() Fee {
	return feeTypeAddLiquidityData
}

// Encode returns the byte representation of a transaction Data.
func (d *CreateSwapPoolData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

type AddLiquidityData struct {
	Coin0          CoinID
	Coin1          CoinID
	Volume0        *big.Int
	MaximumVolume1 *big.Int
}

func NewAddLiquidityData() *AddLiquidityData {
	return &AddLiquidityData{}
}

func (d *AddLiquidityData) SetCoin0(id uint64) *AddLiquidityData {
	d.Coin0 = CoinID(id)
	return d
}
func (d *AddLiquidityData) SetCoin1(id uint64) *AddLiquidityData {
	d.Coin1 = CoinID(id)
	return d
}

func (d *AddLiquidityData) SetVolume0(value0 *big.Int) *AddLiquidityData {
	d.Volume0 = value0
	return d
}
func (d *AddLiquidityData) SetMaximumVolume1(maximumVolume1 *big.Int) *AddLiquidityData {
	d.MaximumVolume1 = maximumVolume1
	return d
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

func NewRemoveLiquidityData() *RemoveLiquidityData {
	return &RemoveLiquidityData{}
}

func (d *RemoveLiquidityData) SetCoin0(id uint64) *RemoveLiquidityData {
	d.Coin0 = CoinID(id)
	return d
}
func (d *RemoveLiquidityData) SetCoin1(id uint64) *RemoveLiquidityData {
	d.Coin1 = CoinID(id)
	return d
}
func (d *RemoveLiquidityData) SetLiquidity(liquidity *big.Int) *RemoveLiquidityData {
	d.Liquidity = liquidity
	return d
}
func (d *RemoveLiquidityData) SetMinimumVolume0(minimumVolume0 *big.Int) *RemoveLiquidityData {
	d.MinimumVolume0 = minimumVolume0
	return d
}
func (d *RemoveLiquidityData) SetMinimumVolume1(minimumVolume1 *big.Int) *RemoveLiquidityData {
	d.MinimumVolume1 = minimumVolume1
	return d
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

func NewBuySwapPoolData() *BuySwapPoolData {
	return &BuySwapPoolData{}
}

// SetCoinToSell sets ID of a coin to get.
func (d *BuySwapPoolData) SetCoinToSell(id uint64) *BuySwapPoolData {
	d.CoinToSell = CoinID(id)
	return d
}

// SetCoinToBuy sets ID of a coin to give.
func (d *BuySwapPoolData) SetCoinToBuy(id uint64) *BuySwapPoolData {
	d.CoinToBuy = CoinID(id)
	return d
}

// SetValueToBuy sets amount of CoinToBuy to get.
func (d *BuySwapPoolData) SetValueToBuy(value *big.Int) *BuySwapPoolData {
	d.ValueToBuy = value
	return d
}

// SetMaximumValueToSell sets maximum value of coins to sell.
func (d *BuySwapPoolData) SetMaximumValueToSell(value *big.Int) *BuySwapPoolData {
	d.MaximumValueToSell = value
	return d
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

func NewSellAllSwapPoolData() *SellAllSwapPoolData {
	return &SellAllSwapPoolData{}
}

// SetCoinToSell sets ID of a coin to give.
func (d *SellAllSwapPoolData) SetCoinToSell(id uint64) *SellAllSwapPoolData {
	d.CoinToSell = CoinID(id)
	return d
}

// SetCoinToBuy sets ID of a coin to get.
func (d *SellAllSwapPoolData) SetCoinToBuy(id uint64) *SellAllSwapPoolData {
	d.CoinToBuy = CoinID(id)
	return d
}

// SetMinimumValueToBuy sets minimum value of coins to get
func (d *SellAllSwapPoolData) SetMinimumValueToBuy(value *big.Int) *SellAllSwapPoolData {
	d.MinimumValueToBuy = value
	return d
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

func NewSellSwapPoolData() *SellSwapPoolData {
	return &SellSwapPoolData{}
}

// SetCoinToSell sets ID of a coin to give
func (d *SellSwapPoolData) SetCoinToSell(id uint32) *SellSwapPoolData {
	d.CoinToSell = CoinID(id)
	return d
}

// SetCoinToBuy sets ID of a coin to get
func (d *SellSwapPoolData) SetCoinToBuy(id uint32) *SellSwapPoolData {
	d.CoinToBuy = CoinID(id)
	return d
}

// SetValueToSell sets amount of CoinToSell to give
func (d *SellSwapPoolData) SetValueToSell(value *big.Int) *SellSwapPoolData {
	d.ValueToSell = value
	return d
}

// SetMinimumValueToBuy sets minimum value of coins to get
func (d *SellSwapPoolData) SetMinimumValueToBuy(value *big.Int) *SellSwapPoolData {
	d.MinimumValueToBuy = value
	return d
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
