package transaction

import (
	"testing"
)

const txPriceVote = "0xf852010201011382c101808001b845f8431ba00e6ceba5074a56661daf2099872627973e9ee09f82519893a1fda16717b4eec1a00664a550774a27d6f6a56c58d53d39ff46719ddd53423a371339314a65857196"

func TestDecode_priceVote(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txPriceVote)
	if err != nil {
		t.Fatal(err)
	}

	if decode.GetTransaction().Type != TypePriceVote {
		t.Error("price vote transaction type is invalid", decode.GetTransaction().Type)
	}
}
