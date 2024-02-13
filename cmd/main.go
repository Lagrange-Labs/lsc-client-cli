package main

import (
	"fmt"
	"html/template"
	"math/big"
	"os"
	"os/exec"
	"runtime"
	"strings"

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
RPCEndpoint = "{{.RPCEndpoint}}"
EthereumURL = "{{.EthereumRPCURL}}"
CommitteeSCAddress = "{{.CommitteeSCAddress}}"
BLSPrivateKey = "{{.BLSPrivateKey}}"
ECDSAPrivateKey = "{{.ECDSAPrivateKey}}"
PullInterval = "100ms"
BLSCurve = "{{.BLSCurve}}"`

	dockerImageBase = "lagrangelabs/lagrange-node"
)

var (
	configFileFlag = &cli.StringFlag{
		Name:    config.FlagCfg,
		Value:   "config.toml",
		Usage:   "Configuration `FILE`",
		Aliases: []string{"c"},
	}

	configFileName = "client_[chainName]_[blsPrivKey].toml"
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
		choice, err := utils.StringPrompt("Enter the operation to perform (`d`eposit, `r`un, `e`xit): ")
		if err != nil {
			logger.Fatalf("Failed to get operation: %s", err)
		}
		switch strings.ToLower(choice) {
		case "d", "deposit":
			if err := depositFunds(cfg); err != nil {
				logger.Infof("Failed to deposit funds: %v, going back to the main menu!", err)
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

func depositFunds(cfg *config.Config) error {
	privateKey, err := utils.PasswordPrompt("Enter the ECDSA Private Key: ")
	if err != nil {
		logger.Fatalf("Failed to get Private Key: %s", err)
	}
	chainOps, err := utils.NewChainOps(cfg.EthereumRPCURL, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %w", err)
	}
	amount, err := utils.IntegerPrompt("Enter the Amount: ")
	if err != nil {
		logger.Fatalf("Failed to get Amount: %s", err)
	}
	amountInt := new(big.Int).SetInt64(int64(amount))

	for {
		if err = chainOps.Deposit(cfg.StakeManagerAddr, cfg.WETHAddr, amountInt); err != nil {
			logger.Infof("Failed to deposit funds: %s", err)
			isRetry, err := utils.ConfirmPrompt("Do you want to retry? (y/n): ")
			if err != nil {
				logger.Fatalf("Failed to get answer: %s", err)
			}
			if isRetry {
				continue
			}
		}
		break
	}

	if err != nil {
		return fmt.Errorf("failed to deposit funds: %w", err)
	}

	logger.Infof("Funds deposited successfully")

	shares, err := chainOps.GetOperatorShares(cfg.StakeManagerAddr, cfg.WETHAddr)
	if err != nil {
		return fmt.Errorf("failed to get operator shares: %w", err)
	}

	logger.Infof("The total amount: %d", shares)

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

	clientCfg := new(config.ClientConfig)
	clientCfg.ServerGrpcURL = cfg.GRPCURL
	clientCfg.RPCEndpoint = cfg.RPCEndpoint
	clientCfg.EthereumRPCURL = cfg.EthereumRPCURL
	clientCfg.CommitteeSCAddress = cfg.CommitteeSCAddr
	clientCfg.BLSCurve = cfg.BLSCurve

	// Get the Chain Name and ECDSA Private Key
	var err error
	clientCfg.ChainName, err = utils.StringPrompt("Enter the Chain Name: ")
	if err != nil {
		logger.Fatalf("Failed to get Chain Name: %s", err)
	}
	clientCfg.ECDSAPrivateKey, err = utils.PasswordPrompt("Enter the ECDSA Private Key: ")
	if err != nil {
		logger.Fatalf("Failed to get ECDSA Private Key: %s", err)
	}

	// Create the BLS Key Pair
	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(cfg.BLSCurve))
	isConfirmed, err := utils.ConfirmPrompt("Do you want to generate a new BLS Key Pair? (y/n): ")
	if err != nil {
		return fmt.Errorf("failed to get answer: %w", err)
	}
	if !isConfirmed {
		clientCfg.BLSPrivateKey, err = utils.StringPrompt("Enter the BLS Private Key: ")
		if err != nil {
			return fmt.Errorf("failed to get BLS Private Key: %w", err)
		}
	} else {
		privKey, err := blsScheme.GenerateRandomKey()
		if err != nil {
			return fmt.Errorf("failed to generate BLS Key Pair: %w", err)
		}
		clientCfg.BLSPrivateKey = crypto.Bytes2Hex(privKey)

	}
	pubRawKey, err := blsScheme.GetPublicKey(crypto.Hex2Bytes(clientCfg.BLSPrivateKey), false)
	if err != nil {
		logger.Fatalf("Failed to get BLS Public Key: %s", err)
	}
	logger.Infof("BLS Key Pair loaded, private key: %v public key: %v", clientCfg.BLSPrivateKey, crypto.Bytes2Hex(pubRawKey))

	chainOps, err := utils.NewChainOps(cfg.EthereumRPCURL, strings.TrimLeft(clientCfg.ECDSAPrivateKey, "0x"))
	if err != nil {
		return fmt.Errorf("failed to create ChainOps instance: %s", err)
	}

	// yes/no prompt to ask the registration to the committee
	// if no, need to register within this flow
	isConfirmed, err = utils.ConfirmPrompt("Do you want to register to the committee now? (y/n): ")
	if err != nil {
		return fmt.Errorf("failed to get answer: %w", err)
	}
	if isConfirmed {
		// register to the committee
		pubKey, err := utils.ConvertBLSKey(pubRawKey)
		if err != nil {
			return fmt.Errorf("failed to convert BLS Public Key: %s", err)
		}
		for {
			if err := chainOps.Register(cfg.LagrangeServiceSCAddr, pubKey); err != nil {
				logger.Infof("Failed to register to the committee: %s", err)
				isRetry, err := utils.ConfirmPrompt("Do you want to retry? (y/n): ")
				if err != nil {
					logger.Fatalf("failed to get answer: %w", err)
				}
				if isRetry {
					continue
				}
			}
			break
		}
	}
	// Subscribe chain
	isConfirmed, err = utils.ConfirmPrompt("Do you want to subscribe to the dedicated chain? (y/n): ")
	if err != nil {
		return fmt.Errorf("failed to get answer: %w", err)
	}
	if isConfirmed {
		chainID, err := utils.IntegerPrompt("Enter the Chain ID: ")
		if err != nil {
			return fmt.Errorf("failed to get Chain ID: %s", err)
		}
		for {
			if err := chainOps.Subscribe(cfg.LagrangeServiceSCAddr, uint32(chainID)); err != nil {
				logger.Infof("Failed to subscribe to the dedicated chain: %s", err)
				isRetry, err := utils.ConfirmPrompt("Do you want to retry? (y/n): ")
				if err != nil {
					return fmt.Errorf("failed to get answer: %w", err)
				}
				if isRetry {
					continue
				}
			}
			break
		}
	}

	// Create the Client Config file
	tmplClient, err := template.New("client").Parse(clientConfigTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse client config template: %s", err)
	}
	configFileName = fmt.Sprintf("client_%s_%s.toml", clientCfg.ChainName, clientCfg.BLSPrivateKey)
	clientConfigFile, err := os.Create(configFileName)
	if err != nil {
		return fmt.Errorf("Failed to create client config file: %s", err)
	}
	defer clientConfigFile.Close()
	if err := tmplClient.Execute(clientConfigFile, clientCfg); err != nil {
		return fmt.Errorf("failed to execute client config template: %s", err)
	}

	logger.Infof("Client Config file created: %s", configFileName)

	return utils.RunDockerImage(dockerImageName, configFileName)
}
