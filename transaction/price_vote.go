package transaction

import "github.com/ethereum/go-ethereum/rlp"

type PriceVoteData struct {
	Price uint
}

func NewPriceVoteData() *PriceVoteData {
	return &PriceVoteData{}
}

func (d *PriceVoteData) SetPrice(price uint) *PriceVoteData {
	d.Price = price
	return d
}

// Type returns Data type of the transaction.
func (d *PriceVoteData) Type() Type {
	return TypePriceVote
}

// Fee returns commission of transaction Data
func (d *PriceVoteData) fee() fee {
	return feePriceVote
}

// Encode returns the byte representation of a transaction Data.
func (d *PriceVoteData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
