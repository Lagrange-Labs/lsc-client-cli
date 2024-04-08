package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Lagrange-Labs/client-cli/crypto"
	"github.com/Lagrange-Labs/client-cli/logger"
	"github.com/Lagrange-Labs/client-cli/scinterface/avs"
	"github.com/Lagrange-Labs/client-cli/scinterface/lagrange"
)

var lagrangeAVSSalt = []byte("lagrange-avs")

// ChainOps is a wrapper for Ethereum chain operations.
type ChainOps struct {
	client     *ethclient.Client
	auth       *bind.TransactOpts
	privateKey *ecdsa.PrivateKey
}

// NewChainOps creates a new ChainOps instance.
func NewChainOps(rpcEndpoint string, privateKey string) (*ChainOps, error) {
	privateKey = strings.TrimPrefix(privateKey, "0x")
	privateKeyECDSA, err := ecrypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	auth, err := crypto.GetSigner(context.Background(), client, privateKeyECDSA)
	if err != nil {
		return nil, err
	}

	return &ChainOps{
		client:     client,
		auth:       auth,
		privateKey: privateKeyECDSA,
	}, nil
}

// Register registers a new validator.
func (c *ChainOps) Register(serviceAddr, signAddr string, blsPubKeys [][2]*big.Int) error {
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}
	avsAddr, err := lagrangeService.AvsDirectory(nil)
	if err != nil {
		return err
	}
	avsDirectory, err := avs.NewAvs(avsAddr, c.client)
	if err != nil {
		return err
	}
	var salt [32]byte
	copy(salt[:], lagrangeAVSSalt)
	header, err := c.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	expiry := header.Time + 300
	digestHash, err := avsDirectory.CalculateOperatorAVSRegistrationDigestHash(nil, c.auth.From, common.HexToAddress(serviceAddr), salt, big.NewInt(int64(expiry)))
	if err != nil {
		return err
	}
	signature, err := ecrypto.Sign(digestHash[:], c.privateKey)
	if err != nil {
		return err
	}
	signature[64] += 27

	tx, err := lagrangeService.Register(c.auth, common.HexToAddress(signAddr), blsPubKeys, lagrange.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature,
		Salt:      salt,
		Expiry:    big.NewInt(int64(expiry)),
	})
	if err != nil {
		return fmt.Errorf("failed to register: %v", err)
	}

	return c.WaitForMined(tx)
}

// AddBlsPubKeys adds BLS Public keys to the validator.
func (c *ChainOps) AddBlsPubKeys(serviceAddr string, blsPubKeys [][2]*big.Int) error {
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	logger.Infof("Adding BLS public keys %s from %s", blsPubKeys, c.auth.From.String())

	tx, err := lagrangeService.AddBlsPubKeys(c.auth, blsPubKeys)
	if err != nil {
		return fmt.Errorf("failed to add BLS keys: %v", err)
	}

	return c.WaitForMined(tx)
}

// Suscribe subscribes the dedicated chain.
func (c *ChainOps) Subscribe(serviceAddr string, chainID uint32) error {
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	logger.Infof("Subscribing to chain %d from %s", chainID, c.auth.From.String())

	tx, err := lagrangeService.Subscribe(c.auth, chainID)
	if err != nil {
		return fmt.Errorf("failed to subscribe: %v", err)
	}

	return c.WaitForMined(tx)
}

// WaitForMined waits for a transaction to be mined.
func (c *ChainOps) WaitForMined(tx *types.Transaction) error {
	logger.Infof("Waiting for transaction %s to be mined", tx.Hash().String())
	receipt, err := bind.WaitMined(context.Background(), c.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}
	if receipt.Status != 1 {
		return fmt.Errorf("transaction failed: %v", receipt)
	}
	logger.Infof("Transaction %s mined", tx.Hash().String())
	return nil
}
