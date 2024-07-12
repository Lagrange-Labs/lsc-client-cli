package utils

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestBN254(t *testing.T) {
	privateKeys := []string{
		"0x00000000000000000000000000000000000000000000000000000000499602d3",
		"0x00000000000000000000000000000000000000000000000000000000499602d4",
		"0x00000000000000000000000000000000000000000000000000000000499602d5",
		"0x00000000000000000000000000000000000000000000000000000000499602d6",
		"0x00000000000000000000000000000000000000000000000000000000499602d7",
		"0x00000000000000000000000000000000000000000000000000000000499602d8",
	}

	count := 2
	digest := "0xe7419ae945e234b771ebf93cbf198d6cfcdbd8447301f1c24500b95a22cf38ee"

	scheme := &BN254Scheme{}
	signatures := make([][]byte, count)
	pubKeyG2s := make([][]byte, count)
	for i := 0; i < count; i++ {
		privKey := common.FromHex(privateKeys[i])
		pubKey, err := scheme.GetPublicKeyG1(privKey, true)
		require.NoError(t, err)
		PrintG1(fmt.Sprintf("G1 PubKey %d: ", i), pubKey)
		signatures[i], err = scheme.Sign(privKey, common.FromHex(digest))
		require.NoError(t, err)
		pubKeyG2s[i], err = scheme.GetPublicKeyG2(privKey, true)
		require.NoError(t, err)
	}

	aggPubKeyG2, err := scheme.AggregateG2(pubKeyG2s)
	require.NoError(t, err)
	PrintG2("Aggregated PubKey G2: ", aggPubKeyG2)

	aggSig, err := scheme.AggregateG1(signatures)
	require.NoError(t, err)
	PrintG1("Aggregated Signature: ", aggSig)
}
