package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/crypto"
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
)

func main() {
	app := cli.NewApp()
	app.Name = "Lagrange Client Utility CLI"

	app.Commands = []*cli.Command{
		{
			Name:  "generate-bls-signature",
			Usage: "Generate a BLS signature for the operator register",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "bls-private-key",
					Usage:   "BLS private key",
					Aliases: []string{"k"},
				},
				&cli.StringFlag{
					Name:    "digest",
					Usage:   "Digest to sign",
					Aliases: []string{"d"},
				},
			},
			Action: generateBLSSignature,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
}

func generateBLSSignature(c *cli.Context) error {
	scheme := new(crypto.BN254Scheme)

	blsPrivateKey := core.Hex2Bytes(c.String("bls-private-key"))
	digest := core.Hex2Bytes(c.String("digest"))

	pubKeyG2, err := scheme.GetPublicKey(blsPrivateKey, false, false)
	if err != nil {
		return err
	}
	pubKeyG1, err := scheme.GetPublicKey(blsPrivateKey, false, true)
	if err != nil {
		return err
	}
	fmt.Printf("%x\n%x\n", pubKeyG1[:32], pubKeyG1[32:])
	signature, err := scheme.Sign(blsPrivateKey, digest, false)
	if err != nil {
		return err
	}

	aggG2 := new(bn254.G2Affine)
	if _, err = aggG2.SetBytes(pubKeyG2); err != nil {
		return err
	}
	fmt.Printf("%x\n", aggG2.X.A1.BigInt(big.NewInt(0)).Bytes())
	fmt.Printf("%x\n", aggG2.X.A0.BigInt(big.NewInt(0)).Bytes())
	fmt.Printf("%x\n", aggG2.Y.A1.BigInt(big.NewInt(0)).Bytes())
	fmt.Printf("%x\n", aggG2.Y.A0.BigInt(big.NewInt(0)).Bytes())

	aggSig := new(bn254.G1Affine)
	if _, err = aggSig.SetBytes(signature); err != nil {
		return err
	}

	fmt.Printf("%x\n", aggSig.X.BigInt(big.NewInt(0)).Bytes())
	fmt.Printf("%x\n", aggSig.Y.BigInt(big.NewInt(0)).Bytes())

	return nil
}
