package main

import (
	"fmt"
	"math/big"
	"os"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/client-cli/config"
	"github.com/Lagrange-Labs/client-cli/utils"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	nutils "github.com/Lagrange-Labs/lagrange-node/utils"
)

const (
	flagKeyType         = "key-type"
	flagKeyPasswordPath = "password-path"
	flagPrivateKey      = "private-key"
	flagKeyPath         = "key-path"
	flagNetwork         = "network"
	flagChain           = "chain"
	flagDockerImage     = "docker-image"
	flagKeyIndex        = "key-index"
	flagPrometheusPort  = "prometheus-port"
)

var (
	configFileFlag = &cli.StringFlag{
		Name:    config.FlagCfg,
		Value:   "config.toml",
		Usage:   "Configuration `FILE`",
		Aliases: []string{"c"},
	}
	keyTypeFlag = &cli.StringFlag{
		Name:    flagKeyType,
		Value:   "bls",
		Usage:   "Key `TYPE` (bls/ecdsa)",
		Aliases: []string{"t"},
	}
	keyPasswordPathFlag = &cli.StringFlag{
		Name:    flagKeyPasswordPath,
		Value:   "",
		Usage:   "Keystore Password `PATH`",
		Aliases: []string{"p"},
	}
	keyPathFlag = &cli.StringFlag{
		Name:    flagKeyPath,
		Value:   "",
		Usage:   "Key `PATH`",
		Aliases: []string{"f"},
	}
	privateKeyFlag = &cli.StringFlag{
		Name:    flagPrivateKey,
		Value:   "",
		Usage:   "Private `KEY`",
		Aliases: []string{"k"},
	}
	networkFlag = &cli.StringFlag{
		Name:    flagNetwork,
		Value:   "mainnet",
		Usage:   "Network `NAME` (mainnet/holesky)",
		Aliases: []string{"n"},
	}
	chainFlag = &cli.StringFlag{
		Name:    flagChain,
		Value:   "optimism",
		Usage:   "Chain `NAME` (optimism/base/arbitrum ...)",
		Aliases: []string{"r"},
	}
	dockerImageFlag = &cli.StringFlag{
		Name:    flagDockerImage,
		Value:   "lagrangelabs/lagrange-node:v0.3.15",
		Usage:   "Docker `IMAGE`",
		Aliases: []string{"i"},
	}
	keyIndexFlag = &cli.IntFlag{
		Name:    flagKeyIndex,
		Value:   0,
		Usage:   "BLS Key `INDEX`",
		Aliases: []string{"i"},
	}
	prometheusPortFlag = &cli.StringFlag{
		Name:    flagPrometheusPort,
		Value:   "8080",
		Usage:   "Prometheus `PORT`",
		Aliases: []string{"p"},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "Lagrange Client CLI"

	app.Commands = []*cli.Command{
		{
			Name:  "version",
			Usage: "Prints the version of the Lagrange Client CLI",
			Action: func(c *cli.Context) error {
				w := os.Stdout
				fmt.Fprintf(w, "Version:      %s\n", "v0.2.0")
				fmt.Fprintf(w, "Go version:   %s\n", runtime.Version())
				fmt.Fprintf(w, "OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
				return nil
			},
		},
		{
			Name:  "generate-keystore",
			Usage: "Generate a new keystore file for the given key type (bls/ecdsa)",
			Flags: []cli.Flag{
				keyTypeFlag,
				keyPasswordPathFlag,
			},
			Action: generateKeystore,
		},
		{
			Name:  "import-keystore",
			Usage: "Import a private key to the keystore for the given key type (bls/ecdsa)",
			Flags: []cli.Flag{
				keyTypeFlag,
				keyPasswordPathFlag,
				privateKeyFlag,
			},
			Action: importKeystore,
		},
		{
			Name:  "export-keystore",
			Usage: "Export a private key from the keystore",
			Flags: []cli.Flag{
				keyTypeFlag,
				keyPasswordPathFlag,
				keyPathFlag,
			},
			Action: exportKeystore,
		},
		{
			Name:  "export-public-key",
			Usage: "Export the public key from the keystore",
			Flags: []cli.Flag{
				keyTypeFlag,
				keyPathFlag,
			},
			Action: exportPublicKey,
		},
		{
			Name:  "register-operator",
			Usage: "Register the operator to the committee",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
			},
			Action: registerOperator,
		},
		{
			Name:  "deregister-operator",
			Usage: "Deregister the operator from the committee",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
			},
			Action: deregisterOperator,
		},
		{
			Name:  "update-bls-pub-key",
			Usage: "Update the BLS public key at the given index",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
				keyIndexFlag,
			},
			Action: updateBlsPubKey,
		},
		{
			Name:  "update-signer-address",
			Usage: "Update the signer address",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
			},
			Action: updateSignerAddress,
		},
		{
			Name:  "add-bls-pub-key",
			Usage: "Add the BLS public key to the operator",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
			},
			Action: addBlsPubKey,
		},
		{
			Name:  "remove-bls-pub-key",
			Usage: "Remove the BLS public key at the given index",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
				keyIndexFlag,
			},
			Action: removeBlsPubKey,
		},
		{
			Name:  "subscribe-chain",
			Usage: "Subscribe to a chain for the given `CHAIN_ID`",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
				chainFlag,
			},
			Action: subscribeChain,
		},
		{
			Name:  "unsubscribe-chain",
			Usage: "Unsubscribe from a chain for the given `CHAIN_ID`",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
				chainFlag,
			},
			Action: unsubscribeChain,
		},
		{
			Name:  "generate-config",
			Usage: "Generate a client config file",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
				chainFlag,
			},
			Action: generateConfig,
		},
		{
			Name:  "generate-docker-compose",
			Usage: "Generate a docker-compose file for node deployment",
			Flags: []cli.Flag{
				configFileFlag,
				dockerImageFlag,
				prometheusPortFlag,
			},
			Action: generateDockerCompose,
		},
		{
			Name:  "deploy",
			Usage: "Deploy the Lagrange Node Client with the given config file",
			Flags: []cli.Flag{
				configFileFlag,
				dockerImageFlag,
				prometheusPortFlag,
			},
			Action: clientDeploy,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
}

func generateKeystore(c *cli.Context) error {
	keyType := strings.ToLower(c.String(flagKeyType))
	passwordPath := c.String(flagKeyPasswordPath)
	return utils.GenerateKeystore(keyType, passwordPath)
}

func importKeystore(c *cli.Context) error {
	keyType := strings.ToLower(c.String(flagKeyType))
	passwordPath := c.String(flagKeyPasswordPath)
	privKey := nutils.Hex2Bytes(c.String(flagPrivateKey))
	return utils.ImportFromPrivateKey(keyType, passwordPath, privKey)
}

func exportKeystore(c *cli.Context) error {
	keyType := strings.ToLower(c.String(flagKeyType))
	passwordPath := c.String(flagKeyPasswordPath)
	keyPath := c.String(flagKeyPath)
	privKey, err := utils.ExportKeystore(keyType, passwordPath, keyPath)
	if err != nil {
		return err
	}
	logger.Infof("Private Key: %s", nutils.Bytes2Hex(privKey))
	return nil
}

func exportPublicKey(c *cli.Context) error {
	keyType := strings.ToLower(c.String(flagKeyType))
	keyPath := c.String(flagKeyPath)
	if err := utils.ExportPublicKey(keyType, keyPath); err != nil {
		return err
	}
	return nil
}

func registerOperator(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	cliCfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}
	blsPubRawKeys := make([][2]*big.Int, 0)
	pubRawKey, err := utils.ConvertBLSKey(nutils.Hex2Bytes(cliCfg.BLSPublicKey))
	if err != nil {
		return fmt.Errorf("failed to convert BLS public key: %w", err)
	}
	blsPubRawKeys = append(blsPubRawKeys, pubRawKey)

	// register the operator to the committee
	logger.Infof("Registering with BLS public key: %s and sign address: %s", blsPubRawKeys, cliCfg.SignerAddress)
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.Register(network, cliCfg.SignerAddress, blsPubRawKeys); err != nil {
		logger.Infof("Failed to register to the committee: %s", err)
	}

	return nil
}

func deregisterOperator(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	cliCfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}

	// deregister the operator from the committee
	logger.Infof("Deregistering with sign address: %s", cliCfg.SignerAddress)
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.Deregister(network); err != nil {
		logger.Infof("Failed to deregister from the committee: %s", err)
	}

	return nil
}

func updateBlsPubKey(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	index := uint32(c.Int(flagKeyIndex))
	cliCfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}
	pubRawKey, err := utils.ConvertBLSKey(nutils.Hex2Bytes(cliCfg.BLSPublicKey))
	if err != nil {
		return fmt.Errorf("failed to convert BLS public key: %w", err)
	}

	// update the BLS public key
	logger.Infof("Updating BLS public key  at index: %d with BLS public key: %s", index, pubRawKey)
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.UpdateBlsPubKey(network, index, pubRawKey); err != nil {
		logger.Infof("Failed to update BLS public key: %s", err)
	}

	return nil
}

func updateSignerAddress(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	cliCfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}

	// update the signer address
	logger.Infof("Updating signer address with signer address: %s", cliCfg.SignerAddress)
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.UpdateSignerAddress(network, cliCfg.SignerAddress); err != nil {
		logger.Infof("Failed to update signer address: %s", err)
	}

	return nil
}

func addBlsPubKey(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	cliCfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}
	blsPubRawKeys := make([][2]*big.Int, 0)
	pubRawKey, err := utils.ConvertBLSKey(nutils.Hex2Bytes(cliCfg.BLSPublicKey))
	if err != nil {
		return fmt.Errorf("failed to convert BLS public key: %w", err)
	}
	blsPubRawKeys = append(blsPubRawKeys, pubRawKey)

	// add the BLS public key
	logger.Infof("Adding BLS public key with BLS public key: %s", pubRawKey)
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.AddBlsPubKeys(network, blsPubRawKeys); err != nil {
		logger.Infof("Failed to add BLS public key: %s", err)
	}

	return nil
}

func removeBlsPubKey(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	index := uint32(c.Int(flagKeyIndex))
	cliCfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}

	// remove the BLS public key
	logger.Infof("Removing BLS public key at index: %d", index)
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.RemoveBlsPubKeys(network, []uint32{index}); err != nil {
		logger.Infof("Failed to remove BLS public key: %s", err)
	}

	return nil
}

func subscribeChain(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	chain := c.String(flagChain)
	if _, ok := config.ChainBatchConfigs[chain]; !ok {
		return fmt.Errorf("invalid chain: %s, should be one of (optimism, base, arbitrum ...)", chain)
	}
	cliCfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}

	// subscribe chain
	logger.Infof("Subscribing to the dedicated chain: %s", chain)
	if err := chainOps.Subscribe(network, chain); err != nil {
		logger.Infof("Failed to subscribe to the dedicated chain: %s", err)
	}

	return nil
}

func unsubscribeChain(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	chain := c.String(flagChain)
	if _, ok := config.ChainBatchConfigs[chain]; !ok {
		return fmt.Errorf("invalid chain: %s, should be one of (optimism, base, arbitrum ...)", chain)
	}
	cliCfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}

	// subscribe chain
	logger.Infof("Unsubscribing from the dedicated chain: %s", chain)
	if err := chainOps.Unsubscribe(network, chain); err != nil {
		logger.Infof("Failed to subscribe to the dedicated chain: %s", err)
	}

	return nil
}

func generateConfig(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	chain := c.String(flagChain)
	if _, ok := config.ChainBatchConfigs[chain]; !ok {
		return fmt.Errorf("invalid chain: %s, should be one of (optimism, base, arbitrum ...)", chain)
	}
	cfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}
	clientCfg := new(config.ClientConfig)
	clientCfg.EthereumRPCURL = cfg.EthereumRPCURL
	clientCfg.CommitteeSCAddress = config.NetworkConfigs[network].CommitteeSCAddress
	clientCfg.BLSCurve = cfg.BLSCurve
	clientCfg.ConcurrentFetchers = cfg.ConcurrentFetchers
	clientCfg.ChainName = chain
	clientCfg.ServerGrpcURL = config.NetworkConfigs[network].GRPCServerURLs[chain]
	clientCfg.L1RPCEndpoint = cfg.L1RPCEndpoint
	clientCfg.L2RPCEndpoint = cfg.L2RPCEndpoint
	clientCfg.BeaconURL = cfg.BeaconURL
	clientCfg.BatchInbox = config.ChainBatchConfigs[chain].BatchInbox
	clientCfg.BatchSender = config.ChainBatchConfigs[chain].BatchSender
	clientCfg.OperatorAddress = cfg.OperatorAddress
	clientCfg.BLSPubKey = cfg.BLSPublicKey
	clientCfg.BLSKeystorePath = cfg.BLSKeystorePath
	clientCfg.BLSKeystorePasswordPath = cfg.BLSKeystorePasswordPath
	clientCfg.SignerECDSAKeystorePath = cfg.SignerECDSAKeystorePath
	clientCfg.SignerECDSAKeystorePasswordPath = cfg.SignerECDSAKeystorePasswordPath

	configFilePath, err := config.GenerateClientConfig(clientCfg, network)
	if err != nil {
		return fmt.Errorf("failed to generate client config: %w", err)
	}

	logger.Infof("Client Config file created: %s", configFilePath)

	return nil
}

func generateDockerCompose(c *cli.Context) error {
	logger.Infof("Generating Docker Compose file")
	if _, err := utils.GenerateDockerComposeFile(c.String(flagDockerImage), c.String(flagPrometheusPort), c.String(config.FlagCfg)); err != nil {
		return fmt.Errorf("failed to generate docker compose file: %w", err)
	}

	return nil
}

func clientDeploy(c *cli.Context) error {
	dockerImageName := c.String(flagDockerImage)
	prometheusPort := c.String(flagPrometheusPort)
	logger.Infof("Deploying Lagrange Node Client with Image: %s", dockerImageName)
	// check if the docker image exists locally
	if err := utils.CheckDockerImageExists(dockerImageName); err != nil {
		return fmt.Errorf("failed to check docker image: %s", err)
	}

	return utils.RunDockerImage(dockerImageName, prometheusPort, c.String(config.FlagCfg))
}
