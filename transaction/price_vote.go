package transaction

import "github.com/ethereum/go-ethereum/rlp"

// PriceVoteData is a Data of Transaction for
type PriceVoteData struct {
	Price uint
}

// NewPriceVoteData returns new PriceVoteData of Transaction for
func NewPriceVoteData() *PriceVoteData {
	return &PriceVoteData{}
}

// SetPrice sets price
func (d *PriceVoteData) SetPrice(price uint) *PriceVoteData {
	d.Price = price
	return d
}

// Type returns Data type of the transaction.
func (d *PriceVoteData) Type() Type {
	return TypePriceVote
}

// Fee returns commission of transaction Data
func (d *PriceVoteData) Fee() Fee {
	return feeTypePriceVote
}

// Encode returns the byte representation of a transaction Data.
func (d *PriceVoteData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
