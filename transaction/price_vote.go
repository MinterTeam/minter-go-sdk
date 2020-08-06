package transaction

import "github.com/ethereum/go-ethereum/rlp"

type PriceVoteData struct {
	Price uint
}

func (d *PriceVoteData) SetPrice(price uint) *PriceVoteData {
	d.Price = price
	return d
}

func (d *PriceVoteData) Type() Type {
	return TypePriceVote
}

func (d *PriceVoteData) Fee() Fee {
	return feePriceVote
}

func (d *PriceVoteData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
