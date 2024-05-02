package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Lagrange-Labs/client-cli/config"
	"github.com/Lagrange-Labs/client-cli/scinterface/avs"
	"github.com/Lagrange-Labs/client-cli/scinterface/lagrange"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	nutils "github.com/Lagrange-Labs/lagrange-node/utils"
)

var lagrangeAVSSalt = []byte("lagrange-avs")

// ChainOps is a wrapper for Ethereum chain operations.
type ChainOps struct {
	client     *ethclient.Client
	auth       *bind.TransactOpts
	privateKey *ecdsa.PrivateKey
}

// NewChainOps creates a new ChainOps instance.
func NewChainOps(network, rpcEndpoint string, privateKey string) (*ChainOps, error) {
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	if config.NetworkConfigs[network].ChainID != uint32(chainID.Int64()) {
		return nil, fmt.Errorf("chain ID mismatch: expected %d, got %d", config.ChainBatchConfigs[network].ChainID, chainID.Int64())
	}

	auth, err := nutils.GetSigner(context.Background(), client, privateKey)
	if err != nil {
		return nil, err
	}

	privateKey = strings.TrimPrefix(privateKey, "0x")
	privateKeyECDSA, err := ecrypto.HexToECDSA(privateKey)
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
func (c *ChainOps) Register(network, signAddr string, blsPubKeys [][2]*big.Int) error {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
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
	// add timestamp to salt
	t := time.Now().UnixNano()
	copy(salt[len(lagrangeAVSSalt):], big.NewInt(t).Bytes())
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
func (c *ChainOps) AddBlsPubKeys(network string, blsPubKeys [][2]*big.Int) error {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	tx, err := lagrangeService.AddBlsPubKeys(c.auth, blsPubKeys)
	if err != nil {
		return fmt.Errorf("failed to add BLS keys: %v", err)
	}

	return c.WaitForMined(tx)
}

// Suscribe subscribes the dedicated chain.
func (c *ChainOps) Subscribe(network, chain string) error {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	chainID := config.ChainBatchConfigs[chain].ChainID
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	tx, err := lagrangeService.Subscribe(c.auth, chainID)
	if err != nil {
		return fmt.Errorf("failed to subscribe: %v", err)
	}

	return c.WaitForMined(tx)
}

// Unsubscribe unsubscribes the dedicated chain.
func (c *ChainOps) Unsubscribe(network, chain string) error {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	chainID := config.ChainBatchConfigs[chain].ChainID
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	tx, err := lagrangeService.Unsubscribe(c.auth, chainID)
	if err != nil {
		return fmt.Errorf("failed to unsubscribe: %v", err)
	}

	return c.WaitForMined(tx)
}

// Deregsiter deregisters the validator.
func (c *ChainOps) Deregister(network string) error {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	tx, err := lagrangeService.Deregister(c.auth)
	if err != nil {
		return fmt.Errorf("failed to deregister: %v", err)
	}

	return c.WaitForMined(tx)
}

// UpdateBlsPubKey updates the BLS public key at the given index.
func (c *ChainOps) UpdateBlsPubKey(network string, index uint32, blsPubKey [2]*big.Int) error {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	tx, err := lagrangeService.UpdateBlsPubKey(c.auth, index, blsPubKey)
	if err != nil {
		return fmt.Errorf("failed to update BLS key: %v", err)
	}

	return c.WaitForMined(tx)
}

// RemoveBlsPubKeys removes the BLS public keys at the given indices.
func (c *ChainOps) RemoveBlsPubKeys(network string, indices []uint32) error {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	tx, err := lagrangeService.RemoveBlsPubKeys(c.auth, indices)
	if err != nil {
		return fmt.Errorf("failed to remove BLS keys: %v", err)
	}

	return c.WaitForMined(tx)
}

// UpdateSignerAddress updates the signer address.
func (c *ChainOps) UpdateSignerAddress(network, newSignerAddr string) error {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	tx, err := lagrangeService.UpdateSignAddress(c.auth, common.HexToAddress(newSignerAddr))
	if err != nil {
		return fmt.Errorf("failed to update signer address: %v", err)
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
