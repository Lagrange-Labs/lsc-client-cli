package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/client-cli/config"
	"github.com/Lagrange-Labs/client-cli/utils"
	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	"github.com/Lagrange-Labs/lagrange-node/signer"
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
)

var (
	configFileFlag = &cli.StringFlag{
		Name:    config.FlagCfg,
		Value:   "./config.toml",
		Usage:   "Configuration `FILE`",
		Aliases: []string{"c"},
	}
	bulkConfigFileFlag = &cli.StringFlag{
		Name:    config.FlagBulkCfg,
		Value:   "./config_bulk.toml",
		Usage:   "Configuration `FILE` for bulk operations on multiple chains and nodes",
		Aliases: []string{"b"},
	}
	nodeConfigFileFlag = &cli.StringFlag{
		Name:     config.FlagNodeCfg,
		Usage:    "Node Configuration `FILE`",
		Aliases:  []string{"n"},
		Required: true,
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
		Value:   "lagrangelabs/lagrange-node:v1.1.5",
		Usage:   "Docker `IMAGE`",
		Aliases: []string{"i"},
	}
	keyIndexFlag = &cli.IntFlag{
		Name:    flagKeyIndex,
		Value:   0,
		Usage:   "BLS Key `INDEX`",
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
				fmt.Fprintf(w, "Version:      %s\n", "v1.1.2")
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
				nodeConfigFileFlag,
				dockerImageFlag,
			},
			Action: generateDockerCompose,
		},
		{
			Name:  "deploy",
			Usage: "Deploy the LSC node with the given node config file",
			Flags: []cli.Flag{
				configFileFlag,
				nodeConfigFileFlag,
				dockerImageFlag,
			},
			Action: clientDeploy,
		},
		{
			Name:  "generate-config-deploy",
			Usage: "Deploy the LSC node after generating the client config file and docker-compose file",
			Flags: []cli.Flag{
				configFileFlag,
				networkFlag,
				chainFlag,
				dockerImageFlag,
			},
			Action: deployWithConfig,
		},
		{
			Name:  "generate-signer-config",
			Usage: "Generate config and docker-compose files for the signer gRPC server",
			Flags: []cli.Flag{
				configFileFlag,
				dockerImageFlag,
			},
			Action: generateSignerConfig,
		},
		{
			Name:  "deploy-signer",
			Usage: "Deploy the LSC signer gRPC server with the given signer config file",
			Flags: []cli.Flag{
				configFileFlag,
				dockerImageFlag,
			},
			Action: deploySigner,
		},
		{
			Name:  "bulk-generate-config-deploy",
			Usage: "Deploy the LSC nodes after generating the client config file and docker-compose file for multiple chains",
			Flags: []cli.Flag{
				bulkConfigFileFlag,
				networkFlag,
				dockerImageFlag,
			},
			Action: bulkDeployWithConfig,
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
	privKey := core.Hex2Bytes(c.String(flagPrivateKey))
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
	logger.Infof("Private Key: %s", core.Bytes2Hex(privKey))
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
	// register the operator to the committee
	logger.Infof("Registering with BLS public key id: %s and sign key id: %s", cliCfg.BLSKeyAccountID, cliCfg.SignerKeyAccountID)
	chainOps, err := utils.NewChainOps(network, cliCfg)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.Register(network, cliCfg.SignerKeyAccountID, cliCfg.BLSKeyAccountID); err != nil {
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
	logger.Infof("Deregistering with sign key id: %s", cliCfg.SignerKeyAccountID)
	chainOps, err := utils.NewChainOps(network, cliCfg)
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
	// update the BLS public key
	logger.Infof("Updating BLS public key  at index: %d with BLS public key id: %s", index, cliCfg.BLSKeyAccountID)
	chainOps, err := utils.NewChainOps(network, cliCfg)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.UpdateBlsPubKey(network, index, cliCfg.BLSKeyAccountID); err != nil {
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
	logger.Infof("Updating signer address with sign key id: %s", cliCfg.SignerKeyAccountID)
	chainOps, err := utils.NewChainOps(network, cliCfg)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.UpdateSignerAddress(network, cliCfg.SignerKeyAccountID); err != nil {
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

	// add the BLS public key
	logger.Infof("Adding BLS public key with BLS public key id: %s", cliCfg.BLSKeyAccountID)
	chainOps, err := utils.NewChainOps(network, cliCfg)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}
	if err := chainOps.AddBlsPubKeys(network, cliCfg.BLSKeyAccountID); err != nil {
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
	chainOps, err := utils.NewChainOps(network, cliCfg)
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
	chainOps, err := utils.NewChainOps(network, cliCfg)
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
	chainOps, err := utils.NewChainOps(network, cliCfg)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}

	// subscribe chain
	logger.Infof("Unsubscribing from the dedicated chain: %s", chain)
	if err := chainOps.Unsubscribe(network, chain); err != nil {
		logger.Infof("Failed to unsubscribe from the dedicated chain: %s", err)
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
	nodeCfg := new(config.NodeConfig)
	nodeCfg.CertConfig = cfg.CertConfig
	nodeCfg.EthereumRPCURL = cfg.EthereumRPCURL
	nodeCfg.CommitteeSCAddress = config.NetworkConfigs[network].CommitteeSCAddress
	nodeCfg.BLSCurve = cfg.BLSCurve
	nodeCfg.ConcurrentFetchers = cfg.ConcurrentFetchers
	nodeCfg.ChainName = chain
	nodeCfg.ServerGrpcURL = config.NetworkConfigs[network].GRPCServerURLs[chain]
	nodeCfg.L1RPCEndpoint = cfg.L1RPCEndpoint
	nodeCfg.L2RPCEndpoint = cfg.L2RPCEndpoint
	nodeCfg.BeaconURL = cfg.BeaconURL
	nodeCfg.BatchInbox = config.ChainBatchConfigs[chain].BatchInbox
	nodeCfg.BatchSender = config.ChainBatchConfigs[chain].BatchSender
	nodeCfg.SignerServerURL = cfg.SignerServerURL
	nodeCfg.OperatorAddress = cfg.OperatorAddress
	nodeCfg.BLSKeyAccountID = cfg.BLSKeyAccountID
	nodeCfg.SignerKeyAccountID = cfg.SignerKeyAccountID
	nodeCfg.MetricsEnabled = cfg.MetricsEnabled
	nodeCfg.MetricsServerPort = cfg.MetricsServerPort
	nodeCfg.MetricsServiceName = cfg.MetricsServiceName
	nodeCfg.PrometheusRetentionTime = cfg.PrometheusRetentionTime

	nodeConfigFilePath, err := config.GenerateNodeConfig(nodeCfg, network)
	if err != nil {
		return fmt.Errorf("failed to generate client config: %w", err)
	}

	logger.Infof("Node Config file created: %s", nodeConfigFilePath)

	return c.Set(config.FlagCfg, nodeConfigFilePath)
}

func generateDockerCompose(c *cli.Context) error {
	logger.Infof("Generating Docker Compose file")
	cfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}
	if _, err := utils.GenerateDockerComposeFile(cfg, c.String(flagDockerImage), c.String(config.FlagNodeCfg)); err != nil {
		return fmt.Errorf("failed to generate docker compose file: %w", err)
	}

	return nil
}

func clientDeploy(c *cli.Context) error {
	dockerImageName := c.String(flagDockerImage)
	logger.Infof("Deploying Lagrange Node Client with Image: %s", dockerImageName)
	// check if the docker image exists locally
	if err := utils.CheckDockerImageExists(dockerImageName); err != nil {
		return fmt.Errorf("failed to check docker image: %s", err)
	}
	cfg, err := config.LoadCLIConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}
	return utils.RunClientNode(cfg, dockerImageName, c.String(config.FlagNodeCfg))
}

func deployWithConfig(c *cli.Context) error {
	cliConfigFilePath := c.String(config.FlagCfg)
	if err := generateConfig(c); err != nil {
		return fmt.Errorf("failed to generate client config: %w", err)
	}
	nodeConfigFilePath := c.String(config.FlagCfg)
	fs := flag.NewFlagSet("deploy", flag.ContinueOnError)
	fs.String(config.FlagCfg, cliConfigFilePath, "Configuration FILE")
	fs.String(config.FlagNodeCfg, nodeConfigFilePath, "Node Configuration FILE")
	fs.String(flagDockerImage, c.String(flagDockerImage), "Docker IMAGE")
	c = cli.NewContext(c.App, fs, nil)

	return clientDeploy(c)
}

func generateSignerConfig(c *cli.Context) error {
	cfg, err := signer.Load(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}
	dockerImageName := c.String(flagDockerImage)
	logger.Infof("Deploying Lagrange Signer with Image: %s", dockerImageName)
	// check if the docker image exists locally
	if err := utils.CheckDockerImageExists(dockerImageName); err != nil {
		return fmt.Errorf("failed to check docker image: %s", err)
	}
	dockerComposeFilePath, err := utils.GenerateSignerConfigFile(cfg, dockerImageName)
	if err != nil {
		return fmt.Errorf("failed to generate signer config file: %w", err)
	}

	return c.Set(flagDockerImage, dockerComposeFilePath)
}

func deploySigner(c *cli.Context) error {
	if err := generateSignerConfig(c); err != nil {
		return fmt.Errorf("failed to generate signer config: %w", err)
	}

	dockerComposeFilePath := c.String(flagDockerImage)
	return utils.RunDockerImage(dockerComposeFilePath)
}

func bulkDeployWithConfig(c *cli.Context) error {
	network := strings.ToLower(c.String(flagNetwork))
	if _, ok := config.NetworkConfigs[network]; !ok {
		return fmt.Errorf("invalid network: %s, should be one of (mainnet, holesky)", network)
	}

	cfg, err := config.LoadCLIBulkConfig(c)
	if err != nil {
		return fmt.Errorf("failed to load CLI config: %w", err)
	}

	dockerImageName := c.String(flagDockerImage)

	for _, chain := range cfg.Chains {
		if _, ok := config.ChainBatchConfigs[chain.ChainName]; !ok {
			return fmt.Errorf("invalid chain: %s, should be one of (optimism, base, arbitrum ...)", chain.ChainName)
		}
		for _, node := range chain.AttestationNodes {
			nodeCfg := new(config.NodeConfig)
			nodeCfg.CertConfig = cfg.CertConfig
			nodeCfg.EthereumRPCURL = cfg.EthereumRPCURL
			nodeCfg.CommitteeSCAddress = config.NetworkConfigs[network].CommitteeSCAddress
			nodeCfg.BLSCurve = cfg.BLSCurve
			nodeCfg.ConcurrentFetchers = node.ConcurrentFetchers
			nodeCfg.ChainName = chain.ChainName
			nodeCfg.ServerGrpcURL = config.NetworkConfigs[network].GRPCServerURLs[chain.ChainName]
			nodeCfg.L1RPCEndpoint = cfg.L1RPCEndpoint
			nodeCfg.L2RPCEndpoint = chain.L2RPCEndpoint
			nodeCfg.BeaconURL = cfg.BeaconURL
			nodeCfg.BatchInbox = config.ChainBatchConfigs[chain.ChainName].BatchInbox
			nodeCfg.BatchSender = config.ChainBatchConfigs[chain.ChainName].BatchSender
			nodeCfg.SignerServerURL = cfg.SignerServerURL
			nodeCfg.OperatorAddress = cfg.OperatorAddress
			nodeCfg.BLSKeyAccountID = node.BLSKeyAccountID
			nodeCfg.SignerKeyAccountID = cfg.SignerKeyAccountID
			nodeCfg.MetricsEnabled = node.MetricsEnabled
			nodeCfg.MetricsServerPort = node.MetricsServerPort
			nodeCfg.MetricsServiceName = node.MetricsServiceName
			nodeCfg.PrometheusRetentionTime = node.PrometheusRetentionTime

			nodeConfigFilePath, err := config.GenerateNodeConfig(nodeCfg, network)
			if err != nil {
				return fmt.Errorf("failed to generate client config: %w", err)
			}
			logger.Infof("Node Config file created: %s", nodeConfigFilePath)

			logger.Infof("Deploying Lagrange Node Client with Image: %s", dockerImageName)
			// check if the docker image exists locally
			if err := utils.CheckDockerImageExists(dockerImageName); err != nil {
				return fmt.Errorf("failed to check docker image: %s", err)
			}
			cliCfg, err := utils.CreateCLIConfigStruct(c, cfg, chain, node)
			if err != nil {
				return fmt.Errorf("failed to create CLI config for deployment: %w", err)
			}
			err = utils.RunClientNode(cliCfg, dockerImageName, nodeConfigFilePath)
			if err != nil {
				return fmt.Errorf("failed to deploy node for chain %s with BLS key account id %s: %w", chain.ChainName, node.BLSKeyAccountID, err)
			}
		}
	}
	return nil
}
