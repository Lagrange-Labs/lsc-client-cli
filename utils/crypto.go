package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

const (
	sizeFr         = fr.Bytes
	sizeFp         = fp.Bytes
	sizePublicKey  = sizeFp
	sizePrivateKey = sizeFr
	sizeSignature  = 2 * sizeFp
)

var (
	order = fr.Modulus()
	one   = new(big.Int).SetInt64(1)
	g     bn254.G1Affine
	g1    bn254.G1Affine
	g2    bn254.G2Affine
	g3    bn254.G2Affine
)

func init() {
	_, _, g, g2 = bn254.Generators()
	g1.Neg(&g)
	g3.Neg(&g2)
}

// Scheme is the crypto scheme implementation for BN254 curve.
type BN254Scheme struct {
}

func (s *BN254Scheme) GenerateRandomKey() ([]byte, error) {
	b := make([]byte, fr.Bits/8+8)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	k := new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(order, one)
	k.Mod(k, n)
	k.Add(k, one)

	privKey := make([]byte, sizeFr)
	k.FillBytes(privKey)

	return privKey, nil
}

func (s *BN254Scheme) GetPublicKeyG1(privKey []byte, isCompressed bool) ([]byte, error) {
	scalar := new(big.Int)
	scalar.SetBytes(privKey[:sizeFr])

	pubKey := new(bn254.G1Affine)
	pubKey.ScalarMultiplication(&g, scalar)
	if isCompressed {
		pubKeyRaw := pubKey.Bytes()
		return pubKeyRaw[:sizeFp], nil
	}
	pubKeyRaw := pubKey.RawBytes()
	return pubKeyRaw[:], nil
}

func (s *BN254Scheme) GetPublicKeyG2(privKey []byte, isCompressed bool) ([]byte, error) {
	scalar := new(big.Int)
	scalar.SetBytes(privKey[:sizeFr])

	pubKey := new(bn254.G2Affine)
	pubKey.ScalarMultiplication(&g2, scalar)
	if isCompressed {
		pubKeyRaw := pubKey.Bytes()
		return pubKeyRaw[:sizeSignature], nil
	}
	pubKeyRaw := pubKey.RawBytes()
	return pubKeyRaw[:], nil
}

// MapToCurve implements the simple hash-and-check (also sometimes try-and-increment) algorithm
// see https://hackmd.io/@benjaminion/bls12-381#Hash-and-check
// Note that this function needs to be the same as the one used in the contract:
// https://github.com/Layr-Labs/eigenlayer-middleware/blob/1feb6ae7e12f33ce8eefb361edb69ee26c118b5d/src/libraries/BN254.sol#L292
// we don't use the newer constant time hash-to-curve algorithms as they are gas-expensive to compute onchain
func MapToCurve(digest [32]byte) *bn254.G1Affine {
	one := new(big.Int).SetUint64(1)
	three := new(big.Int).SetUint64(3)
	x := new(big.Int)
	x.SetBytes(digest[:])
	for {
		// y = x^3 + 3
		xP3 := new(big.Int).Exp(x, big.NewInt(3), fp.Modulus())
		y := new(big.Int).Add(xP3, three)
		y.Mod(y, fp.Modulus())

		if y.ModSqrt(y, fp.Modulus()) == nil {
			x.Add(x, one).Mod(x, fp.Modulus())
		} else {
			var fpX, fpY fp.Element
			fpX.SetBigInt(x)
			fpY.SetBigInt(y)
			return &bn254.G1Affine{
				X: fpX,
				Y: fpY,
			}
		}
	}
}

func (s *BN254Scheme) Sign(privKey, message []byte) ([]byte, error) {
	// Convert the private key to a scalar
	scalar := new(big.Int)
	scalar.SetBytes(privKey[:sizeFr])

	// Hash the message into G1
	var digest [32]byte
	copy(digest[:], message)
	h := MapToCurve(digest)
	sig := new(bn254.G1Affine)
	sig.ScalarMultiplication(h, scalar)
	sigRaw := sig.RawBytes()
	return sigRaw[:], nil
}

func PrintG2(prefix string, point []byte) {
	fmt.Println(prefix)
	p := new(bn254.G2Affine)
	if _, err := p.SetBytes(point); err != nil {
		panic(err)
	}
	fmt.Printf("X.A0: %s, X.A1: %s\n", p.X.A0.String(), p.X.A1.String())
	fmt.Printf("Y.A0: %s, Y.A1: %s\n", p.Y.A0.String(), p.Y.A1.String())
}

func PrintG1(prefix string, point []byte) {
	fmt.Println(prefix)
	p := new(bn254.G1Affine)
	if _, err := p.SetBytes(point); err != nil {
		panic(err)
	}
	fmt.Printf("X: %s\n", p.X.String())
	fmt.Printf("Y: %s\n", p.Y.String())
}

func (s *BN254Scheme) AggregateG2(points [][]byte) ([]byte, error) {
	aggG2 := new(bn254.G2Affine)
	for i, point := range points {
		s := new(bn254.G2Affine)
		if _, err := s.SetBytes(point); err != nil {
			return nil, err
		}
		if i == 0 {
			aggG2 = s
		} else {
			aggG2.Add(aggG2, s)
		}
	}

	pRaw := aggG2.Bytes()
	return pRaw[:], nil
}

func (s *BN254Scheme) AggregateG1(points [][]byte) ([]byte, error) {
	aggG1 := new(bn254.G1Affine)
	for i, point := range points {
		p := new(bn254.G1Affine)
		if _, err := p.SetBytes(point); err != nil {
			return nil, err
		}
		if i == 0 {
			aggG1 = p
		} else {
			aggG1.Add(aggG1, p)
		}
	}

	pRaw := aggG1.Bytes()
	return pRaw[:], nil
}
