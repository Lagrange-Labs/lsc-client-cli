package main

import (
	"fmt"
	"html/template"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/client-cli/config"
	"github.com/Lagrange-Labs/client-cli/crypto"
	"github.com/Lagrange-Labs/client-cli/logger"
	"github.com/Lagrange-Labs/client-cli/utils"
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
)

var (
	configFileFlag = &cli.StringFlag{
		Name:    config.FlagCfg,
		Value:   "config.toml",
		Usage:   "Configuration `FILE`",
		Aliases: []string{"c"},
	}

	batchConfig = map[string]struct {
		BatchInbox  string
		BatchSender string
	}{
		"optimism": {
			BatchInbox:  "0xFF00000000000000000000000000000000000010",
			BatchSender: "0x6887246668a3b87F54DeB3b94Ba47a6f63F32985",
		},
		"base": {
			BatchInbox:  "0xFf00000000000000000000000000000000008453",
			BatchSender: "0x5050F69a9786F081509234F1a7F4684b5E5b76C9",
		},
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
			Name:  "run",
			Usage: "Run the CLI with the given configuration file",
			Flags: []cli.Flag{
				configFileFlag,
			},
			Action: run,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	cfg, err := config.Load(c)
	if err != nil {
		logger.Fatalf("Failed to load configuration: %s", err)
	}

	logger.Infof("Loaded configuration: %+v", cfg)

	for {
		choice, err := utils.StringPrompt("Enter the operation to perform (`r`un, re`g`ister, `s`ubscribe, `c`onfig, `e`xit): ")
		if err != nil {
			logger.Fatalf("Failed to get operation: %s", err)
		}
		switch strings.ToLower(choice) {
		case "g", "register":
			if err := registerOperator(cfg); err != nil {
				logger.Infof("Failed to register operator: %v, going back to the main menu!", err)
			}
		case "s", "subscribe":
			if err := subscribeChain(cfg); err != nil {
				logger.Infof("Failed to subscribe to the chain: %v, going back to the main menu!", err)
			}
		case "c", "config":
			if err := generateConfig(cfg); err != nil {
				logger.Infof("Failed to generate the config: %v, going back to the main menu!", err)
			}
		case "r", "run":
			if err := clientDeploy(cfg); err != nil {
				logger.Infof("Failed to deploy the client: %v, going back to the main menu!", err)
			}
		case "e", "exit":
			return nil
		default:
			logger.Info("Invalid operation, please try again")
		}
	}
}

func registerOperator(cfg *config.Config) error {
	// Generate new BLS Key Pairs
	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(cfg.BLSCurve))
	for {
		isConfirmed, err := utils.ConfirmPrompt("Do you want to add a new BLS Key Pair? (y/n): ")
		if err != nil {
			return fmt.Errorf("failed to get answer: %w", err)
		}
		if isConfirmed {
			var privKey []byte
			res, err := utils.StringPrompt("Do you want to generate a new Key Pair (y/bls_private_key): ")
			if err != nil {
				return fmt.Errorf("failed to get answer: %w", err)
			}
			if res == "y" {
				privKey, err = blsScheme.GenerateRandomKey()
				if err != nil {
					return fmt.Errorf("failed to generate BLS Key Pair: %w", err)
				}
			} else {
				privKey = crypto.Hex2Bytes(res)
			}
			pubRawKey, err := blsScheme.GetPublicKey(privKey, false)
			if err != nil {
				logger.Infof("Failed to get BLS Public Key: %s", err)
				continue
			}
			if err := utils.SaveKeyStore("bls", utils.NewKeyStore(privKey, pubRawKey)); err != nil {
				logger.Infof("Failed to save BLS Key Pair: %s", err)
				continue
			}
		} else {
			break
		}
	}

	// Generate a new signer ECDSA private key
	isConfirmed, err := utils.ConfirmPrompt("Do you want to generate a new signer private key? (y/n): ")
	if err != nil {
		return fmt.Errorf("failed to get answer: %w", err)
	}
	if isConfirmed {
		privKey, err := ecrypto.GenerateKey()
		if err != nil {
			return fmt.Errorf("failed to generate ECDSA Key Pair: %w", err)
		}
		addr := ecrypto.PubkeyToAddress(privKey.PublicKey)
		signAddress := addr.Hex()
		logger.Infof("Signer ECDSA Key Pair loaded address: %v", signAddress)
		if err := utils.SaveKeyStore("ecdsa", utils.NewKeyStore(privKey.D.Bytes(), addr.Bytes())); err != nil {
			logger.Infof("Failed to save ECDSA Key Pair: %s", err)
		}
	} else {
		privRawKey, err := utils.StringPrompt("Enter the Signer ECDSA Private Key, if you want to use the stored key, just press `n`: ")
		if err != nil {
			return fmt.Errorf("failed to get Signer ECDSA Address: %w", err)
		}
		if len(privRawKey) > 1 {
			privKey, err := ecrypto.HexToECDSA(privRawKey)
			if err != nil {
				return fmt.Errorf("failed to convert ECDSA Private Key: %w", err)
			}
			addr := ecrypto.PubkeyToAddress(privKey.PublicKey)
			signAddress := addr.Hex()
			logger.Infof("Signer ECDSA Key Pair loaded address: %v", signAddress)
			if err := utils.SaveKeyStore("ecdsa", utils.NewKeyStore(privKey.D.Bytes(), addr.Bytes())); err != nil {
				logger.Infof("Failed to save ECDSA Key Pair: %s", err)
			}
		}
	}

	// Register to the committee
	blsKeyStores, err := utils.LoadKeyStores("bls")
	if err != nil {
		return fmt.Errorf("failed to load BLS Key Stores: %w", err)
	}
	signKeyStores, err := utils.LoadKeyStores("ecdsa")
	if err != nil {
		return fmt.Errorf("failed to load ECDSA Key Stores: %w", err)
	}
	pubRawKeys := make([][2]*big.Int, 0)
	for _, blsKeyStore := range blsKeyStores {
		pubKey, err := crypto.ConvertBLSKey(crypto.Hex2Bytes(blsKeyStore.PubKey))
		if err != nil {
			logger.Infof("Failed to convert BLS Public Key: %s", err)
			continue
		}
		pubRawKeys = append(pubRawKeys, pubKey)
	}
	for {
		logger.Infof("Registering with BLS public key: %s and sign address: %s", pubRawKeys, signKeyStores[0].PubKey)
		isConfirmed, err := utils.ConfirmPrompt("Do you want to register to the committee? (y/n): ")
		if err != nil {
			return fmt.Errorf("failed to get answer: %w", err)
		}
		if !isConfirmed {
			break
		}
		operatorPrivKey, err := utils.PasswordPrompt("Enter the Operator ECDSA Private Key: ")
		if err != nil {
			return fmt.Errorf("failed to get Operator ECDSA Private Key: %w", err)
		}
		chainOps, err := utils.NewChainOps(cfg.EthereumRPCURL, operatorPrivKey)
		if err != nil {
			return fmt.Errorf("failed to create ChainOps instance: %s", err)
		}
		cfg.OperatorPrivKey = operatorPrivKey
		if err := chainOps.Register(cfg.LagrangeServiceSCAddr, signKeyStores[0].PubKey, pubRawKeys); err != nil {
			logger.Infof("Failed to register to the committee: %s", err)
		} else {
			break
		}
	}

	return nil
}

func subscribeChain(cfg *config.Config) error {
	if len(cfg.OperatorPrivKey) == 0 {
		operatorPrivKey, err := utils.PasswordPrompt("Enter the Operator ECDSA Private Key: ")
		if err != nil {
			return fmt.Errorf("failed to get Operator ECDSA Private Key: %w", err)
		}
		cfg.OperatorPrivKey = operatorPrivKey
	}
	chainOps, err := utils.NewChainOps(cfg.EthereumRPCURL, cfg.OperatorPrivKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}

	// Subscribe chain
	for {
		chainID, err := utils.IntegerPrompt("Enter the Chain ID: ")
		if err != nil {
			return fmt.Errorf("failed to get Chain ID: %s", err)
		}
		if err := chainOps.Subscribe(cfg.LagrangeServiceSCAddr, uint32(chainID)); err != nil {
			logger.Infof("Failed to subscribe to the dedicated chain: %s", err)
		}
		isConfirmed, err := utils.ConfirmPrompt("Do you want to subscribe to another chain? (y/n): ")
		if err != nil {
			return fmt.Errorf("failed to get answer: %w", err)
		}
		if !isConfirmed {
			break
		}
	}

	return nil
}

func generateConfig(cfg *config.Config) error {
	logger.Infof("Generating Client Config file")

	clientCfg := new(config.ClientConfig)
	clientCfg.EthereumRPCURL = cfg.EthereumRPCURL
	clientCfg.CommitteeSCAddress = cfg.CommitteeSCAddr
	clientCfg.BLSCurve = cfg.BLSCurve
	clientCfg.ConcurrentFetchers = cfg.ConcurrentFetchers

	var err error
	// Get the Operator Address
	clientCfg.OperatorAddress, err = utils.StringPrompt("Enter the Operator Address: ")
	if err != nil {
		logger.Fatalf("Failed to get Operator Address: %s", err)
	}
	// Get the Chain Name
	clientCfg.ChainName, err = utils.StringPrompt("Enter the Chain Name: ")
	if err != nil {
		logger.Fatalf("Failed to get Chain Name: %s", err)
	}
	clientCfg.ChainName = strings.ToLower(clientCfg.ChainName)

	// Get the RPC Endpoints
	clientCfg.L1RPCEndpoint, err = utils.StringPrompt("Enter the L1 RPC Endpoint: ")
	if err != nil {
		logger.Fatalf("Failed to get L1 RPC Endpoint: %s", err)
	}
	clientCfg.L2RPCEndpoint, err = utils.StringPrompt("Enter the L2 RPC Endpoint: ")
	if err != nil {
		logger.Fatalf("Failed to get L2 RPC Endpoint: %s", err)
	}
	clientCfg.BeaconURL, err = utils.StringPrompt("Enter the Beacon URL: ")
	if err != nil {
		logger.Fatalf("Failed to get Beacon URL: %s", err)
	}

	// Get the Server gRPC URL
	clientCfg.ServerGrpcURL, err = utils.StringPrompt("Enter the Server gRPC URL: ")
	if err != nil {
		logger.Fatalf("Failed to get Server gRPC URL: %s", err)
	}

	// Select the BLS Key Pair
	blsKeyStores, err := utils.LoadKeyStores("bls")
	if err != nil {
		return fmt.Errorf("failed to load BLS Key Stores: %w", err)
	}
	if len(blsKeyStores) == 0 {
		clientCfg.BLSPrivateKey, err = utils.StringPrompt("No BLS Key Pair found, please enter the BLS Private Key: ")
		if err != nil {
			return fmt.Errorf("failed to get BLS Private Key: %w", err)
		}
	} else {
		for i, blsKeyStore := range blsKeyStores {
			logger.Infof("BLS Key Pair %d: %s", i+1, blsKeyStore.PubKey)
		}
		keyIndex, err := utils.IntegerPrompt("Select the BLS Key Pair (index): ")
		if err != nil {
			return fmt.Errorf("failed to get BLS Key Pair index: %w", err)
		}
		clientCfg.BLSPrivateKey = blsKeyStores[keyIndex-1].PrivKey
	}

	// Select the Signer ECDSA Private Key
	signKeyStores, err := utils.LoadKeyStores("ecdsa")
	if err != nil {
		return fmt.Errorf("failed to load ECDSA Key Stores: %w", err)
	}
	if len(signKeyStores) == 0 {
		clientCfg.SignPrivateKey, err = utils.StringPrompt("No Signer ECDSA Key Pair found, please enter the signer ECDSA Private Key: ")
		if err != nil {
			return fmt.Errorf("failed to get Signer ECDSA Private Key: %w", err)
		}
	} else {
		clientCfg.SignPrivateKey = signKeyStores[0].PrivKey
	}

	clientCfg.BatchInbox = batchConfig[clientCfg.ChainName].BatchInbox
	clientCfg.BatchSender = batchConfig[clientCfg.ChainName].BatchSender

	// Create the Client Config file
	tmplClient, err := template.New("client").Parse(clientConfigTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse client config template: %s", err)
	}
	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(clientCfg.BLSCurve))
	pubRawKey, err := blsScheme.GetPublicKey(crypto.Hex2Bytes(clientCfg.BLSPrivateKey), false)
	if err != nil {
		logger.Fatalf("Failed to get BLS Public Key: %s", err)
	}
	configFileName := fmt.Sprintf("client_%s_%x.toml", clientCfg.ChainName, pubRawKey)
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

func clientDeploy(cfg *config.Config) error {
	dockerImageName := fmt.Sprintf("%s:%s", dockerImageBase, cfg.DockerImageTag)
	logger.Infof("Deploying Lagrange Node Client with Image: %s", dockerImageName)

	// Check if the docker image exists locally
	if err := utils.CheckDockerImageExists(dockerImageName); err != nil {
		logger.Infof("Docker image %s does not exist locally, pulling started", dockerImageName)
		if err := utils.PullDockerImage(dockerImageName); err != nil {
			return fmt.Errorf("failed to pull docker image %s: %s", dockerImageName, err)
		}
	} else {
		cmd := exec.Command("docker", "image", "inspect", dockerImageName, "--format='{{index .RepoDigests 0}}'")
		output, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("failed to get docker image digest: %s", err)
		}
		logger.Infof("Docker image %s exists locally with digest: %s", dockerImageName, output)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %s", err)
	}
	configDirPath := filepath.Clean(filepath.Join(homeDir, configDir))
	files, err := os.ReadDir(configDirPath)
	if err != nil {
		return fmt.Errorf("failed to read config directory: %s", err)
	}
	configFileNames := make([]string, 0)
	for i, file := range files {
		if file.IsDir() {
			continue
		}
		logger.Infof("Config file %d: %s", i+1, file.Name())
		configFileNames = append(configFileNames, filepath.Join(configDirPath, file.Name()))
	}
	index, err := utils.IntegerPrompt("Select the Config file (index): ")
	if err != nil {
		return fmt.Errorf("failed to get Config file index: %s", err)
	}

	return utils.RunDockerImage(dockerImageName, configFileNames[index-1])
}
