package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
	"math/big"
)

// CreateSwapPoolData is a Data of Transaction for creating a liquidity pool for two coins, in volumes specified within this transaction. The volumes will be withdrawn from your balance according to the figure you've specified in the transaction. When a pool is established, a PL-number coin (example: PL-123) is created and issued in the amount equal to the amount of pool liquidity. The calculations related to that liquidity are described below.
type CreateSwapPoolData struct {
	Coin0   CoinID
	Coin1   CoinID
	Volume0 *big.Int
	Volume1 *big.Int
}

// NewCreateSwapPoolData returns CreateSwapPoolData transaction
func NewCreateSwapPoolData() *CreateSwapPoolData {
	return &CreateSwapPoolData{}
}

// SetCoin0 sets first ID of coin to pair
func (d *CreateSwapPoolData) SetCoin0(id uint64) *CreateSwapPoolData {
	d.Coin0 = CoinID(id)
	return d
}

// SetCoin1 sets second ID of coin to pair
func (d *CreateSwapPoolData) SetCoin1(id uint64) *CreateSwapPoolData {
	d.Coin1 = CoinID(id)
	return d
}

// SetVolume0 sets volume to add to reserve of the swap pool of first coin
func (d *CreateSwapPoolData) SetVolume0(value0 *big.Int) *CreateSwapPoolData {
	d.Volume0 = value0
	return d
}

// SetVolume1 sets volume to add to reserve of the swap pool of second coin
func (d *CreateSwapPoolData) SetVolume1(value1 *big.Int) *CreateSwapPoolData {
	d.Volume1 = value1
	return d
}

// Type returns Data type of the transaction.
func (d *CreateSwapPoolData) Type() Type {
	return TypeCreateSwapPool
}

// Encode returns the byte representation of a transaction Data.
func (d *CreateSwapPoolData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

// AddLiquidityData is a Data of Transaction for to add reserves of a pair of coins to the pool. To create liquidity through this pool.
type AddLiquidityData struct {
	Coin0          CoinID   // ID of first coin to pair
	Coin1          CoinID   // ID of second coin to pair
	Volume0        *big.Int // Volume to add to reserve of the swap pool of first coin
	MaximumVolume1 *big.Int // Maximum volume to add to reserve of the swap pool of second coin
}

// NewAddLiquidityData creates AddLiquidityData
func NewAddLiquidityData() *AddLiquidityData {
	return &AddLiquidityData{}
}

// SetCoin0 sets first ID of coin to pair
func (d *AddLiquidityData) SetCoin0(id uint64) *AddLiquidityData {
	d.Coin0 = CoinID(id)
	return d
}

// SetCoin1 sets second ID of coin to pair
func (d *AddLiquidityData) SetCoin1(id uint64) *AddLiquidityData {
	d.Coin1 = CoinID(id)
	return d
}

// SetVolume0 sets volume to add to reserve of the swap pool of first coin
func (d *AddLiquidityData) SetVolume0(value0 *big.Int) *AddLiquidityData {
	d.Volume0 = value0
	return d
}

// SetMaximumVolume1 sets maximum volume to add to reserve of the swap pool of second coin
func (d *AddLiquidityData) SetMaximumVolume1(maximumVolume1 *big.Int) *AddLiquidityData {
	d.MaximumVolume1 = maximumVolume1
	return d
}

// Type returns Data type of the transaction.
func (d *AddLiquidityData) Type() Type {
	return TypeAddLiquidity
}

// Encode returns the byte representation of a transaction Data.
func (d *AddLiquidityData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

// RemoveLiquidityData is a Data of Transaction for withdrawing the reserves of a pair from the pool.
type RemoveLiquidityData struct {
	Coin0          CoinID   // ID of coin to pair
	Coin1          CoinID   // ID of coin to pair
	Liquidity      *big.Int // Volume of shares to be withdrawn from the pool
	MinimumVolume0 *big.Int // Minimum expected volume of coin0 to be returned to the account
	MinimumVolume1 *big.Int // Minimum expected volume of coin1 to be returned to the account
}

// NewRemoveLiquidityData returns RemoveLiquidityData
func NewRemoveLiquidityData() *RemoveLiquidityData {
	return &RemoveLiquidityData{}
}

// SetCoin0 sets first ID of coin to pair
func (d *RemoveLiquidityData) SetCoin0(id uint64) *RemoveLiquidityData {
	d.Coin0 = CoinID(id)
	return d
}

// SetCoin1 sets second ID of coin to pair
func (d *RemoveLiquidityData) SetCoin1(id uint64) *RemoveLiquidityData {
	d.Coin1 = CoinID(id)
	return d
}

// SetLiquidity sets volume of shares to be withdrawn from the pool
func (d *RemoveLiquidityData) SetLiquidity(liquidity *big.Int) *RemoveLiquidityData {
	d.Liquidity = liquidity
	return d
}

// SetMinimumVolume0 sets minimum expected volume of coin0 to be returned to the account
func (d *RemoveLiquidityData) SetMinimumVolume0(minimumVolume0 *big.Int) *RemoveLiquidityData {
	d.MinimumVolume0 = minimumVolume0
	return d
}

// SetMinimumVolume1 sets minimum expected volume of coin1 to be returned to the account
func (d *RemoveLiquidityData) SetMinimumVolume1(minimumVolume1 *big.Int) *RemoveLiquidityData {
	d.MinimumVolume1 = minimumVolume1
	return d
}

// Type returns Data type of the transaction.
func (d *RemoveLiquidityData) Type() Type {
	return TypeRemoveLiquidity
}

// Encode returns the byte representation of a transaction Data.
func (d *RemoveLiquidityData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

// BuySwapPoolData is a Data of Transaction for buying from the swap pool of the pair.
type BuySwapPoolData struct {
	Coins              []CoinID
	ValueToBuy         *big.Int
	MaximumValueToSell *big.Int
}

// NewBuySwapPoolData creates BuySwapPoolData
func NewBuySwapPoolData() *BuySwapPoolData {
	return &BuySwapPoolData{}
}

// AddCoin sets ID of a coin in exchanging route.
func (d *BuySwapPoolData) AddCoin(ids ...uint64) *BuySwapPoolData {
	for _, id := range ids {
		d.Coins = append(d.Coins, CoinID(id))
	}
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

// Encode returns the byte representation of a transaction Data.
func (d *BuySwapPoolData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

// SellAllSwapPoolData is a Data of Transaction for selling all existing coins from the swap pool of the pair.
// Coin to spend (Coins[0]) will be used as GasCoin to pay fee.
type SellAllSwapPoolData struct {
	Coins             []CoinID // List of coin IDs from given to received.
	MinimumValueToBuy *big.Int // Minimum value of coin to get.
}

// NewSellAllSwapPoolData creates SellAllSwapPoolData
func NewSellAllSwapPoolData() *SellAllSwapPoolData {
	return &SellAllSwapPoolData{}
}

// AddCoin sets ID of a coin in exchanging route.
func (d *SellAllSwapPoolData) AddCoin(ids ...uint64) *SellAllSwapPoolData {
	for _, id := range ids {
		d.Coins = append(d.Coins, CoinID(id))
	}
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

// Encode returns the byte representation of a transaction Data.
func (d *SellAllSwapPoolData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}

// SellSwapPoolData is a Data of Transaction for selling from the swap pool of the pair.
type SellSwapPoolData struct {
	Coins             []CoinID // List of coin IDs from given to received.
	ValueToSell       *big.Int // Amount of coin to spend (first coin in Coins list).
	MinimumValueToBuy *big.Int // Minimum value of coin to get.
}

// NewSellSwapPoolData creates SellSwapPoolData
func NewSellSwapPoolData() *SellSwapPoolData {
	return &SellSwapPoolData{}
}

// AddCoin sets ID of a coin in exchanging route.
func (d *SellSwapPoolData) AddCoin(ids ...uint64) *SellSwapPoolData {
	for _, id := range ids {
		d.Coins = append(d.Coins, CoinID(id))
	}
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

// Encode returns the byte representation of a transaction Data.
func (d *SellSwapPoolData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
