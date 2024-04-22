package utils

import (
	"fmt"
	"math/big"
)

// ConvertBLSKey converts a BLS public key from a string to a big.Int array.
func ConvertBLSKey(blsPubKey []byte) ([2]*big.Int, error) {
	// Parse BLS public key
	if len(blsPubKey) != 64 {
		return [2]*big.Int{}, fmt.Errorf("invalid BLS public key")
	}

	var pubKey [2]*big.Int
	pubKey[0] = new(big.Int).SetBytes(blsPubKey[:32])
	pubKey[1] = new(big.Int).SetBytes(blsPubKey[32:])

	return pubKey, nil
}
