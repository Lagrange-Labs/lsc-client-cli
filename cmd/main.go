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
	flagKeyType     = "key-type"
	flagKeyPassword = "password"
	flagPrivateKey  = "private-key"
	flagKeyPath     = "key-path"
	flagNetwork     = "network"
	flagChain       = "chain"
	flagRollupRPC   = "rollup-rpc"
	flagDockerImage = "docker-image"
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
	keyPasswordFlag = &cli.StringFlag{
		Name:    flagKeyPassword,
		Value:   "",
		Usage:   "Key `PASSWORD`",
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
		Usage:   "Chain `NAME` (optimism/base)",
		Aliases: []string{"h"},
	}
	rollupRPCFlag = &cli.StringFlag{
		Name:    flagRollupRPC,
		Value:   "",
		Usage:   "Rollup RPC `URL`",
		Aliases: []string{"r"},
	}
	dockerImageFlag = &cli.StringFlag{
		Name:    flagDockerImage,
		Value:   "lagrangelabs/lagrange-node:v0.3.13",
		Usage:   "Docker `IMAGE`",
		Aliases: []string{"i"},
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
				fmt.Fprintf(w, "Version:      %s\n", "v0.1.0")
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
				keyPasswordFlag,
			},
			Action: generateKeystore,
		},
		{
			Name:  "import-keystore",
			Usage: "Import a private key to the keystore for the given key type (bls/ecdsa)",
			Flags: []cli.Flag{
				keyTypeFlag,
				keyPasswordFlag,
				privateKeyFlag,
			},
			Action: importKeystore,
		},
		{
			Name:  "export-keystore",
			Usage: "Export a private key from the keystore",
			Flags: []cli.Flag{
				keyTypeFlag,
				keyPasswordFlag,
				keyPathFlag,
			},
			Action: exportKeystore,
		},
		{
			Name:  "register-operator",
			Usage: "Register the operator to the committee",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
				chainFlag,
			},
			Action: registerOperator,
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
				rollupRPCFlag,
			},
			Action: generateConfig,
		},
		{
			Name:  "deploy",
			Usage: "Deploy the Lagrange Node Client with the given config file",
			Flags: []cli.Flag{
				configFileFlag,
				dockerImageFlag,
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
	password := c.String(flagKeyPassword)
	return utils.GenerateKeystore(keyType, password)
}

func importKeystore(c *cli.Context) error {
	keyType := strings.ToLower(c.String(flagKeyType))
	password := c.String(flagKeyPassword)
	privKey := nutils.Hex2Bytes(c.String(flagPrivateKey))
	return utils.ImportFromPrivateKey(keyType, password, privKey)
}

func exportKeystore(c *cli.Context) error {
	keyType := strings.ToLower(c.String(flagKeyType))
	password := c.String(flagKeyPassword)
	keyPath := c.String(flagKeyPath)
	privKey, err := utils.ExportKeystore(keyType, password, keyPath)
	if err != nil {
		return err
	}
	logger.Infof("Private Key: %s", nutils.Bytes2Hex(privKey))
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

func subscribeChain(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	chain := c.String(flagChain)
	if _, ok := config.ChainBatchConfigs[chain]; !ok {
		return fmt.Errorf("invalid chain: %s, should be one of (optimism, base)", chain)
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
		return fmt.Errorf("invalid chain: %s, should be one of (optimism, base)", chain)
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
		return fmt.Errorf("invalid chain: %s, should be one of (optimism, base)", chain)
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
	clientCfg.L2RPCEndpoint = c.String(flagRollupRPC)
	clientCfg.BeaconURL = cfg.BeaconURL
	clientCfg.BatchInbox = config.ChainBatchConfigs[chain].BatchInbox
	clientCfg.BatchSender = config.ChainBatchConfigs[chain].BatchSender
	clientCfg.OperatorAddress = cfg.OperatorAddress
	clientCfg.BLSPubKey = cfg.BLSPublicKey
	clientCfg.BLSKeystorePath = cfg.BLSKeystorePath
	clientCfg.BLSKeystorePassword = cfg.BLSKeystorePassword
	clientCfg.BLSKeystorePasswordPath = cfg.BLSKeystorePasswordPath
	clientCfg.SignerECDSAKeystorePath = cfg.SignerECDSAKeystorePath
	clientCfg.SignerECDSAKeystorePassword = cfg.SignerECDSAKeystorePassword
	clientCfg.SignerECDSAKeystorePasswordPath = cfg.SignerECDSAKeystorePasswordPath

	configFilePath, err := config.GenerateClientConfig(clientCfg, network)
	if err != nil {
		return fmt.Errorf("failed to generate client config: %w", err)
	}

	logger.Infof("Client Config file created: %s", configFilePath)

	return nil
}

func clientDeploy(c *cli.Context) error {
	dockerImageName := c.String(flagDockerImage)
	logger.Infof("Deploying Lagrange Node Client with Image: %s", dockerImageName)
	// check if the docker image exists locally
	if err := utils.CheckDockerImageExists(dockerImageName); err != nil {
		return fmt.Errorf("failed to check docker image: %s", err)
	}

	return utils.RunDockerImage(dockerImageName, c.String(config.FlagCfg))
}
