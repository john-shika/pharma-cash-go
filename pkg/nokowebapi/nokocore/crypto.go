package nokocore

import (
	rand "crypto/rand"
	"math/big"
)

func CryptoRandomRangeInt(min, max int) (int, error) {
	var n *big.Int
	var err error
	diff := max - min
	if n, err = rand.Int(rand.Reader, big.NewInt(int64(diff))); err != nil {
		return 0, err
	}
	return min + int(n.Int64()), nil
}

func CryptoRandomRangeInt32(min, max int32) (int32, error) {
	var n *big.Int
	var err error
	diff := max - min
	if n, err = rand.Int(rand.Reader, big.NewInt(int64(diff))); err != nil {
		return 0, err
	}
	return min + int32(n.Int64()), nil
}

func CryptoRandomRangeInt64(min, max int64) (int64, error) {
	var n *big.Int
	var err error
	diff := max - min
	if n, err = rand.Int(rand.Reader, big.NewInt(diff)); err != nil {
		return 0, err
	}
	return min + n.Int64(), nil
}
