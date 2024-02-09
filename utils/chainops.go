package utils

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Lagrange-Labs/client-cli/crypto"
	"github.com/Lagrange-Labs/client-cli/logger"
	"github.com/Lagrange-Labs/client-cli/scinterface/ERC20"
	"github.com/Lagrange-Labs/client-cli/scinterface/lagrange"
	"github.com/Lagrange-Labs/client-cli/scinterface/staking"
)

// ChainOps is a wrapper for Ethereum chain operations.
type ChainOps struct {
	client *ethclient.Client
	auth   *bind.TransactOpts
}

// NewChainOps creates a new ChainOps instance.
func NewChainOps(rpcEndpoint string, privateKey string) (*ChainOps, error) {
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	auth, err := crypto.GetSigner(context.Background(), client, privateKey)
	if err != nil {
		return nil, err
	}

	return &ChainOps{
		client: client,
		auth:   auth,
	}, nil
}

// Deposit deposits funds to the StakeManager contract.
func (c *ChainOps) Deposit(smAddr, tokenAddr string, amount *big.Int) error {
	// Approve token transfer
	token, err := ERC20.NewERC20(common.HexToAddress(tokenAddr), c.client)
	if err != nil {
		return err
	}
	tx, err := token.Approve(c.auth, common.HexToAddress(smAddr), amount)
	if err != nil {
		return fmt.Errorf("failed to approve token transfer: %v", err)
	}
	if err := c.WaitForMined(tx); err != nil {
		return err
	}

	// Deposit to StakeManager
	sm, err := staking.NewStaking(common.HexToAddress(smAddr), c.client)
	tx, err = sm.Deposit(c.auth, common.HexToAddress(tokenAddr), amount)
	if err != nil {
		return fmt.Errorf("failed to deposit to StakeManager: %v", err)
	}

	return c.WaitForMined(tx)
}

// Register registers a new validator.
func (c *ChainOps) Register(serviceAddr string, blsPubKey [2]*big.Int) error {
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	logger.Infof("Registering with BLS public key %s from %s", blsPubKey, c.auth.From.String())

	tx, err := lagrangeService.Register(c.auth, blsPubKey)
	if err != nil {
		return fmt.Errorf("failed to register: %v", err)
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
	receipt, err := bind.WaitMined(context.Background(), c.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}
	if receipt.Status != 1 {
		return fmt.Errorf("transaction failed: %v", receipt)
	}
	return nil
}

// GetOperatorShares returns the operator shares for a given operator and token.
func (c *ChainOps) GetOperatorShares(smAddr, tokenAddr string) (uint64, error) {
	sm, err := staking.NewStaking(common.HexToAddress(smAddr), c.client)
	if err != nil {
		return 0, err
	}

	votingPower, err := sm.OperatorShares(nil, c.auth.From, common.HexToAddress(tokenAddr))

	return votingPower.Uint64(), err
}
