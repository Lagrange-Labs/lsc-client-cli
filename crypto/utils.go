package crypto

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Hex2Bytes converts a hex string to bytes.
func Hex2Bytes(hex string) []byte {
	return common.FromHex(hex)
}

// Bytes2Hex converts bytes to a hex string.
func Bytes2Hex(bytes []byte) string {
	return common.Bytes2Hex(bytes)
}

// GetSigner returns a transaction signer.
func GetSigner(ctx context.Context, c *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainID, err := c.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactorWithChainID(privateKey, chainID)
}

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
