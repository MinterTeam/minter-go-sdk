package transaction

import (
	"errors"
	"fmt"
	"math/big"
)

// BipToPip converts BIP to PIP (multiplies input by 1e18)
func BipToPip(bip *big.Int) *big.Int {
	return big.NewInt(0).Mul(bip, big.NewInt(1e18))
}

// FloatBipToPip converts BIP to PIP (multiplies input by 1e18)
func FloatBipToPip(bip float64) *big.Int {
	p := big.NewInt(10)
	p.Exp(p, big.NewInt(18), nil)
	result, _ := big.NewFloat(0).Mul(big.NewFloat(bip), big.NewFloat(0).SetInt(p)).Int(nil)

	return result
}

// StringToBigInt converts string to BigInt, panics on empty strings and errors
func StringToBigInt(s string) *big.Int {
	result, err := stringToBigInt(s)
	if err != nil {
		panic(err)
	}

	return result
}

func stringToBigInt(s string) (*big.Int, error) {
	if s == "" {
		return nil, errors.New("string is empty")
	}

	b, success := big.NewInt(0).SetString(s, 10)
	if !success {
		return nil, fmt.Errorf("cannot decode %s into big.Int", s)
	}

	return b, nil
}

// IsValidBigInt verifies that string is a valid int
func IsValidBigInt(s string) bool {
	b, err := stringToBigInt(s)
	if err != nil {
		return false
	}

	if b.Sign() == -1 {
		return false
	}

	return true
}
