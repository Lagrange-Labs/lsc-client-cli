package utils

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Lagrange-Labs/client-cli/config"
	"github.com/Lagrange-Labs/client-cli/scinterface/avs"
	"github.com/Lagrange-Labs/client-cli/scinterface/committee"
	"github.com/Lagrange-Labs/client-cli/scinterface/lagrange"
	"github.com/Lagrange-Labs/lsc-node/core/crypto"
	"github.com/Lagrange-Labs/lsc-node/core/logger"
)

var lagrangeAVSSalt = []byte("lagrange-avs")

// ChainOps is a wrapper for Ethereum chain operations.
type ChainOps struct {
	chainID      *big.Int
	client       *ethclient.Client
	signerClient *SignerClient
	cliCfg       *config.CLIConfig
}

// NewChainOps creates a new ChainOps instance.
func NewChainOps(network string, cliCfg *config.CLIConfig) (*ChainOps, error) {
	client, err := ethclient.Dial(cliCfg.EthereumRPCURL)
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

	signerClient, err := NewSignerClient(cliCfg.SignerServerURL, cliCfg.CertConfig)
	if err != nil {
		return nil, err
	}

	return &ChainOps{
		chainID:      chainID,
		client:       client,
		signerClient: signerClient,
		cliCfg:       cliCfg,
	}, nil
}

func (c *ChainOps) getRawTxOpts() *bind.TransactOpts {
	return &bind.TransactOpts{
		From: common.HexToAddress(c.cliCfg.OperatorAddress),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
		NoSend: true,
	}
}

func (c *ChainOps) getSaltAndExpiry() ([32]byte, *big.Int) {
	var salt [32]byte
	copy(salt[:], lagrangeAVSSalt)
	// add timestamp to salt
	t := time.Now().UnixNano()
	copy(salt[len(lagrangeAVSSalt):], big.NewInt(t).Bytes())
	header, err := c.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return [32]byte{}, nil
	}
	expiry := header.Time + 300
	return salt, big.NewInt(int64(expiry))
}

func (c *ChainOps) getAVSDSignature(network string) (*lagrange.ISignatureUtilsSignatureWithSaltAndExpiry, error) {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return nil, err
	}
	avsAddr, err := lagrangeService.AvsDirectory(nil)
	if err != nil {
		return nil, err
	}
	avsDirectory, err := avs.NewAvs(avsAddr, c.client)
	if err != nil {
		return nil, err
	}
	salt, expiry := c.getSaltAndExpiry()
	digestHash, err := avsDirectory.CalculateOperatorAVSRegistrationDigestHash(nil, common.HexToAddress(c.cliCfg.OperatorAddress), common.HexToAddress(serviceAddr), salt, expiry)
	if err != nil {
		return nil, err
	}
	signature, err := c.signerClient.Sign(digestHash[:], c.cliCfg.OperatorKeyAccountID, true)
	if err != nil {
		return nil, err
	}
	signature[64] += 27

	return &lagrange.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature,
		Salt:      salt,
		Expiry:    expiry,
	}, nil
}

func (c *ChainOps) getCommitteeDigestHash(network string, blsKeyIDs ...string) (*lagrange.IBLSKeyCheckerBLSKeyWithProof, error) {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return nil, err
	}
	committeeAdr, err := lagrangeService.Committee(nil)
	if err != nil {
		return nil, err
	}
	committee, err := committee.NewCommittee(committeeAdr, c.client)
	if err != nil {
		return nil, err
	}
	salt, expiry := c.getSaltAndExpiry()
	digestHash, err := committee.CalculateKeyWithProofHash(nil, common.HexToAddress(c.cliCfg.OperatorAddress), salt, expiry)
	if err != nil {
		return nil, err
	}

	arg0, arg1, arg2, err := c.generateBLSSignature(digestHash[:], blsKeyIDs...)
	if err != nil {
		return nil, err
	}

	return &lagrange.IBLSKeyCheckerBLSKeyWithProof{
		BlsG1PublicKeys: arg0,
		AggG2PublicKey:  arg1,
		Signature:       arg2,
		Salt:            salt,
		Expiry:          expiry,
	}, nil
}

// Register registers a new validator.
func (c *ChainOps) Register(network, signerKeyID string, blsAccountIDs ...string) error {
	signAddr, err := c.signerClient.GetPublicKey(signerKeyID, false)
	if err != nil {
		return err
	}
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(config.NetworkConfigs[network].LagrangeServiceSCAddress), c.client)
	if err != nil {
		return err
	}

	avsSig, err := c.getAVSDSignature(network)
	if err != nil {
		return err
	}
	blsSig, err := c.getCommitteeDigestHash(network, blsAccountIDs...)
	if err != nil {
		return err
	}

	tx, err := lagrangeService.Register(c.getRawTxOpts(), common.BytesToAddress(signAddr), *blsSig, *avsSig)
	if err != nil {
		return fmt.Errorf("failed to register: %v", err)
	}

	return c.WaitForMined(tx)
}

// AddBlsPubKeys adds BLS Public keys to the validator.
func (c *ChainOps) AddBlsPubKeys(network string, blsKeyIDs ...string) error {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	blsSig, err := c.getCommitteeDigestHash(network, blsKeyIDs...)
	if err != nil {
		return err
	}
	tx, err := lagrangeService.AddBlsPubKeys(c.getRawTxOpts(), *blsSig)
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

	tx, err := lagrangeService.Subscribe(c.getRawTxOpts(), chainID)
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

	tx, err := lagrangeService.Unsubscribe(c.getRawTxOpts(), chainID)
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

	tx, err := lagrangeService.Deregister(c.getRawTxOpts())
	if err != nil {
		return fmt.Errorf("failed to deregister: %v", err)
	}

	return c.WaitForMined(tx)
}

// UpdateBlsPubKey updates the BLS public key at the given index.
func (c *ChainOps) UpdateBlsPubKey(network string, index uint32, blsKeyID string) error {
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	blsSig, err := c.getCommitteeDigestHash(network, blsKeyID)
	if err != nil {
		return err
	}
	tx, err := lagrangeService.UpdateBlsPubKey(c.getRawTxOpts(), index, *blsSig)
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

	tx, err := lagrangeService.RemoveBlsPubKeys(c.getRawTxOpts(), indices)
	if err != nil {
		return fmt.Errorf("failed to remove BLS keys: %v", err)
	}

	return c.WaitForMined(tx)
}

// UpdateSignerAddress updates the signer address.
func (c *ChainOps) UpdateSignerAddress(network, newSignerKeyID string) error {
	newSignerAddr, err := c.signerClient.GetPublicKey(newSignerKeyID, false)
	if err != nil {
		return err
	}
	serviceAddr := config.NetworkConfigs[network].LagrangeServiceSCAddress
	lagrangeService, err := lagrange.NewLagrange(common.HexToAddress(serviceAddr), c.client)
	if err != nil {
		return err
	}

	tx, err := lagrangeService.UpdateSignAddress(c.getRawTxOpts(), common.BytesToAddress(newSignerAddr))
	if err != nil {
		return fmt.Errorf("failed to update signer address: %v", err)
	}

	return c.WaitForMined(tx)
}

// WaitForMined send the transaction and wait for it to be mined.
func (c *ChainOps) WaitForMined(tx *types.Transaction) error {
	esigner := types.LatestSignerForChainID(c.chainID)
	txHash := esigner.Hash(tx)
	signature, err := c.signerClient.Sign(txHash[:], c.cliCfg.OperatorKeyAccountID, true)
	if err != nil {
		return err
	}
	tx, err = tx.WithSignature(esigner, signature)
	if err != nil {
		return err
	}
	if err = c.client.SendTransaction(context.Background(), tx); err != nil {
		return fmt.Errorf("failed to send transaction: %v", err)
	}

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

func (c *ChainOps) generateBLSSignature(digest []byte, blsKeyIDs ...string) (
	blsG1PublicKeys [][2]*big.Int,
	aggG2PublicKey [2][2]*big.Int,
	signature [2]*big.Int,
	err error,
) {
	pubKeyG2s := make([][]byte, len(blsKeyIDs))
	signatures := make([][]byte, len(blsKeyIDs))
	for i, blsKeyID := range blsKeyIDs {
		pubKeyG2s[i], err = c.signerClient.GetPublicKey(blsKeyID, false)
		if err != nil {
			return
		}
		var pubKey []byte
		pubKey, err = c.signerClient.GetPublicKey(blsKeyID, true)
		if err != nil {
			return
		}

		pubKeyX := new(big.Int).SetBytes(pubKey[:32])
		pubKeyY := new(big.Int).SetBytes(pubKey[32:])
		blsG1PublicKeys = append(blsG1PublicKeys, [2]*big.Int{pubKeyX, pubKeyY})

		signatures[i], err = c.signerClient.Sign(digest, blsKeyID, false)
		if err != nil {
			return
		}
	}

	aggG2Bytes, err := new(crypto.BN254Scheme).AggregatePublicKeys(pubKeyG2s, false)
	if err != nil {
		return blsG1PublicKeys, aggG2PublicKey, signature, err
	}
	aggG2 := new(bn254.G2Affine)
	if _, err = aggG2.SetBytes(aggG2Bytes); err != nil {
		return blsG1PublicKeys, aggG2PublicKey, signature, err
	}
	aggG2PublicKey[0][0] = aggG2.X.A1.BigInt(big.NewInt(0))
	aggG2PublicKey[0][1] = aggG2.X.A0.BigInt(big.NewInt(0))
	aggG2PublicKey[1][0] = aggG2.Y.A1.BigInt(big.NewInt(0))
	aggG2PublicKey[1][1] = aggG2.Y.A0.BigInt(big.NewInt(0))

	aggSigBytes, err := new(crypto.BN254Scheme).AggregateSignatures(signatures, true)
	if err != nil {
		return blsG1PublicKeys, aggG2PublicKey, signature, err
	}
	aggSig := new(bn254.G1Affine)
	if _, err = aggSig.SetBytes(aggSigBytes); err != nil {
		return blsG1PublicKeys, aggG2PublicKey, signature, err
	}
	signature[0] = aggSig.X.BigInt(big.NewInt(0))
	signature[1] = aggSig.Y.BigInt(big.NewInt(0))

	return
}
