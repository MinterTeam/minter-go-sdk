package transaction

import (
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type Check struct {
	Nonce    []byte
	ChainID  ChainID
	DueBlock uint64
	Coin     [10]byte
	Value    *big.Int
	Lock     *big.Int
	V        *big.Int
	R        *big.Int
	S        *big.Int
}

func (check *Check) Sign(prKey string) error {
	h, err := rlpHash([]interface{}{
		check.Nonce,
		check.ChainID,
		check.DueBlock,
		check.Coin,
		check.Value,
		check.Lock,
	})
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(prKey)
	if err != nil {
		return err
	}

	sig, err := crypto.Sign(h[:], privateKey)
	if err != nil {
		return err
	}

	check.R = new(big.Int).SetBytes(sig[:32])
	check.S = new(big.Int).SetBytes(sig[32:64])
	check.V = new(big.Int).SetBytes([]byte{sig[64] + 27})

	return nil
}
