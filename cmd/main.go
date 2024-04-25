package main

import (
	"fmt"
	"html/template"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/client-cli/config"
	"github.com/Lagrange-Labs/client-cli/logger"
	"github.com/Lagrange-Labs/client-cli/utils"
	"github.com/Lagrange-Labs/lagrange-node/crypto"
	nutils "github.com/Lagrange-Labs/lagrange-node/utils"
)

const (
	clientConfigTemplate = `[Client]
GrpcURL = "{{.ServerGrpcURL}}"
Chain = "{{.ChainName}}"
EthereumURL = "{{.EthereumRPCURL}}"
OperatorAddress = "{{.OperatorAddress}}"
CommitteeSCAddress = "{{.CommitteeSCAddress}}"
BLSPrivateKey = "{{.BLSPrivateKey}}"
ECDSAPrivateKey = "{{.SignPrivateKey}}"
PullInterval = "1000ms"
BLSCurve = "{{.BLSCurve}}"

[RpcClient]

	[RpcClient.Optimism]
	RPCURL = "{{.L2RPCEndpoint}}"
	L1RPCURL = "{{.L1RPCEndpoint}}"
	BeaconURL = "{{.BeaconURL}}"
	BatchInbox = "{{.BatchInbox}}"
	BatchSender = "{{.BatchSender}}"
	ConcurrentFetchers = "{{.ConcurrentFetchers}}"

	[RpcClient.Mock]
	RPCURL = "http://localhost:8545"`

	dockerImageBase = "lagrangelabs/lagrange-node"
	configDir       = ".lagrange/config"

	flagKeyType     = "key-type"
	flagKeyPassword = "password"
	flagPrivateKey  = "private-key"
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

func registerOperator(c *cli.Context) error {
	network := c.String(flagNetwork)
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}
	cliCfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}
	// loads the operator private key
	if len(cliCfg.OperatorPrivateKey) == 0 {
		operatorPrivKey, err := utils.ReadPrivateKey("ECDSA", cliCfg.OperatorKeystorePassword, cliCfg.OperatorKeystorePath)
		if err != nil {
			return fmt.Errorf("failed to load Operator Key Stores: %w", err)
		}
		cliCfg.OperatorPrivateKey = nutils.Bytes2Hex(operatorPrivKey)
	}
	// loads the BLS private key
	blsPrivKey, err := utils.ReadPrivateKey("BN254", cliCfg.BLSKeystorePassword, cliCfg.BLSKeystorePath)
	if err != nil {
		return fmt.Errorf("failed to load BLS Key Stores: %w", err)
	}
	blsScheme := crypto.NewBLSScheme(crypto.BN254)
	blsPubKey, err := blsScheme.GetPublicKey(blsPrivKey, false)
	blsPubRawKeys := make([][2]*big.Int, 0)
	pubRawKey, err := utils.ConvertBLSKey(blsPubKey)
	if err != nil {
		return fmt.Errorf("failed to convert BLS Public Key: %w", err)
	}
	blsPubRawKeys = append(blsPubRawKeys, pubRawKey)
	// loads the Signer private key
	signerPrivKey, err := utils.ReadPrivateKey("ECDSA", cliCfg.SignerKeystorePassword, cliCfg.SignerKeystorePath)
	if err != nil {
		return fmt.Errorf("failed to load Signer Key Stores: %w", err)
	}
	signerPrivateKey, err := ecrypto.ToECDSA(signerPrivKey)
	if err != nil {
		return fmt.Errorf("failed to convert private key to ECDSA: %w", err)
	}
	signerAddr := ecrypto.PubkeyToAddress(signerPrivateKey.PublicKey)
	// register the operator to the committee
	logger.Infof("Registering with BLS public key: %s and sign address: %s", blsPubRawKeys, signerAddr.Hex())
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.Register(network, signerAddr.Hex(), blsPubRawKeys); err != nil {
		logger.Infof("Failed to register to the committee: %s", err)
	}

	return nil
}

func subscribeChain(c *cli.Context) error {
	network := c.String(flagNetwork)
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
	// loads the operator private key
	if len(cliCfg.OperatorPrivateKey) == 0 {
		operatorPrivKey, err := utils.ReadPrivateKey("ECDSA", cliCfg.OperatorKeystorePassword, cliCfg.OperatorKeystorePath)
		if err != nil {
			return fmt.Errorf("failed to load Operator Key Stores: %w", err)
		}
		cliCfg.OperatorPrivateKey = nutils.Bytes2Hex(operatorPrivKey)
	}
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivateKey)
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
	network := c.String(flagNetwork)
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
	// loads the operator private key
	if len(cliCfg.OperatorPrivateKey) == 0 {
		operatorPrivKey, err := utils.ReadPrivateKey("ECDSA", cliCfg.OperatorKeystorePassword, cliCfg.OperatorKeystorePath)
		if err != nil {
			return fmt.Errorf("failed to load Operator Key Stores: %w", err)
		}
		cliCfg.OperatorPrivateKey = nutils.Bytes2Hex(operatorPrivKey)
	}
	chainOps, err := utils.NewChainOps(network, cliCfg.EthereumRPCURL, cliCfg.OperatorPrivateKey)
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
	network := c.String(flagNetwork)
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
	// loads the operator private key
	if len(cfg.OperatorPrivateKey) == 0 {
		operatorPrivKey, err := utils.ReadPrivateKey("ECDSA", cfg.OperatorKeystorePassword, cfg.OperatorKeystorePath)
		if err != nil {
			return fmt.Errorf("failed to load Operator Key Stores: %w", err)
		}
		cfg.OperatorPrivateKey = nutils.Bytes2Hex(operatorPrivKey)
	}
	opPrivateKey, err := ecrypto.ToECDSA(nutils.Hex2Bytes(cfg.OperatorPrivateKey))
	if err != nil {
		return fmt.Errorf("failed to convert private key to ECDSA: %w", err)
	}
	opAddr := ecrypto.PubkeyToAddress(opPrivateKey.PublicKey)
	clientCfg.OperatorAddress = opAddr.Hex()

	// loads the BLS private key
	if len(cfg.BLSPrivateKey) == 0 {
		blsPrivKey, err := utils.ReadPrivateKey("BN254", cfg.BLSKeystorePassword, cfg.BLSKeystorePath)
		if err != nil {
			return fmt.Errorf("failed to load BLS Key Stores: %w", err)
		}
		cfg.BLSPrivateKey = nutils.Bytes2Hex(blsPrivKey)
	}
	clientCfg.BLSPrivateKey = cfg.BLSPrivateKey
	// loads the Signer private key
	if len(cfg.SignerPrivateKey) == 0 {
		signerPrivKey, err := utils.ReadPrivateKey("ECDSA", cfg.SignerKeystorePassword, cfg.SignerKeystorePath)
		if err != nil {
			return fmt.Errorf("failed to load Signer Key Stores: %w", err)
		}
		cfg.SignerPrivateKey = nutils.Bytes2Hex(signerPrivKey)
	}
	clientCfg.SignPrivateKey = cfg.SignerPrivateKey

	// Create the Client Config file
	tmplClient, err := template.New("client").Parse(clientConfigTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse client config template: %s", err)
	}
	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(clientCfg.BLSCurve))
	pubRawKey, err := blsScheme.GetPublicKey(nutils.Hex2Bytes(clientCfg.BLSPrivateKey), false)
	if err != nil {
		logger.Fatalf("Failed to get BLS Public Key: %s", err)
	}
	configFileName := fmt.Sprintf("client_%s_%x.toml", clientCfg.ChainName, pubRawKey[:6])
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %s", err)
	}
	err = os.MkdirAll(filepath.Join(homeDir, configDir), 0700)
	if err != nil {
		return fmt.Errorf("failed to create config directory: %s", err)
	}
	configFilePath := filepath.Clean(filepath.Join(homeDir, configDir, configFileName))
	clientConfigFile, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to create client config file: %s", err)
	}
	defer clientConfigFile.Close()
	if err := tmplClient.Execute(clientConfigFile, clientCfg); err != nil {
		return fmt.Errorf("failed to execute client config template: %s", err)
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
