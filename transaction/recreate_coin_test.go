package transaction

import (
	"testing"
)

const txRecreateCoin = "0xf87d0102010110adec808a5350525445535400000089056bc75e2d631000008a043c33c19375648000000a893635c9adc5dea00000808001b845f8431ca04743e4b01fc1c8305bbe9e84f483fb4a7411c419f9ec73124e4e75579a6fd5e0a06d241ed5b6a8c1b9154e7e1cba57de520fc6b5681b2aaa28f578b4e5c071c36b"

func TestDecode_recreateCoin(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txRecreateCoin)
	if err != nil {
		t.Fatal(err)
	}

	if decode.Fee().String() != "10000000000000000000000" {
		t.Error("recreate coin transaction fee is invalid", decode.Fee().String())
	}
}
