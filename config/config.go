package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	nutils "github.com/Lagrange-Labs/lagrange-node/utils"
)

// CLIConfig is the configuration for the lagrange CLI.
type CLIConfig struct {
	OperatorPrivKey                 string
	OperatorAddress                 string
	OperatorKeystorePath            string `mapstructure:"OperatorKeyStorePath"`
	OperatorKeystorePassword        string
	OperatorKeystorePasswordPath    string `mapstructure:"OperatorKeyStorePasswordPath"`
	SignerAddress                   string
	SignerECDSAKeystorePath         string `mapstructure:"SignerECDSAKeystorePath"`
	SignerECDSAKeystorePassword     string
	SignerECDSAKeystorePasswordPath string `mapstructure:"SignerECDSAKeystorePasswordPath"`
	BLSPrivateKey                   string
	BLSPublicKey                    string
	BLSKeystorePath                 string `mapstructure:"BLSKeyStorePath"`
	BLSKeystorePassword             string
	BLSKeystorePasswordPath         string `mapstructure:"BLSKeyStorePasswordPath"`
	EthereumRPCURL                  string `mapstructure:"EthereumRPCURL"`
	L1RPCEndpoint                   string `mapstructure:"L1RPCEndpoint"`
	BeaconURL                       string `mapstructure:"BeaconURL"`
	L2RPCEndpoint                   string `mapstructure:"L2RPCEndpoint"`
	BLSCurve                        string `mapstructure:"BLSCurve"`
	ConcurrentFetchers              int    `mapstructure:"ConcurrentFetchers"`
	MetricsEnabled                  bool   `mapstructure:"MetricsEnabled"`
	MetricsServerPort               string `mapstructure:"MetricsServerPort"`
	HostBindingPort                 string `mapstructure:"HostBindingPort"`
	MetricsServiceName              string `mapstructure:"MetricsServiceName"`
	PrometheusRetentionTime         string `mapstructure:"PrometheusRetentionTime"`
}

// NodeConfig is the configuration for the LSC node.
// This is used to generate the client.toml file.
type NodeConfig struct {
	ChainName                       string
	ServerGrpcURL                   string
	OperatorAddress                 string
	L1RPCEndpoint                   string
	L2RPCEndpoint                   string
	BeaconURL                       string
	EthereumRPCURL                  string
	CommitteeSCAddress              string
	BLSPubKey                       string
	BLSKeystorePath                 string `mapstructure:"BLSKeystorePath"`
	BLSKeystorePasswordPath         string `mapstructure:"BLSKeystorePasswordPath"`
	SignerECDSAKeystorePath         string `mapstructure:"SignerECDSAKeystorePath"`
	SignerECDSAKeystorePasswordPath string `mapstructure:"SignerECDSAKeystorePasswordPath"`
	BatchInbox                      string
	BatchSender                     string
	BLSCurve                        string
	ConcurrentFetchers              int
	MetricsEnabled                  bool
	MetricsServerPort               string
	MetricsServiceName              string
	PrometheusRetentionTime         string
}

// DockerComposeConfig is the configuration for the docker-compose.yml file.
type DockerComposeConfig struct {
	Network                         string `json:"network"`
	ChainName                       string `json:"chain_name"`
	BLSPubKeyPrefix                 string `json:"bls_pub_key"`
	DockerImage                     string `json:"docker_image"`
	ConfigFilePath                  string `json:"config_file_path"`
	BLSKeystorePath                 string `json:"bls_keystore_path"`
	BLSKeystorePasswordPath         string `json:"bls_keystore_password_path"`
	SignerECDSAKeystorePath         string `json:"signer_ecdsa_keystore_path"`
	SignerECDSAKeystorePasswordPath string `json:"signer_ecdsa_keystore_password_path"`
	PrometheusPort                  string `json:"prometheus_port"`
	HostBindingPort                 string `json:"host_binding_port"`
}

// LoadCLIConfig loads the lagrange CLI configuration.
func LoadCLIConfig(ctx *cli.Context) (*CLIConfig, error) {
	var cfg CLIConfig
	viper.SetConfigType("toml")

	configFilePath := ctx.String(FlagCfg)
	if configFilePath != "" {
		dirName, fileName := filepath.Split(configFilePath)

		fileExtension := strings.TrimPrefix(filepath.Ext(fileName), ".")
		fileNameWithoutExtension := strings.TrimSuffix(fileName, "."+fileExtension)

		viper.AddConfigPath(dirName)
		viper.SetConfigName(fileNameWithoutExtension)
		viper.SetConfigType(fileExtension)
	}
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("LAGRANGE_CLI")
	if err := viper.ReadInConfig(); err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			return nil, err
		} else if len(configFilePath) > 0 {
			logger.Warnf("config file `%s` not found, the path should be absolute or relative to the current working directory like `./config.toml`", configFilePath)
			return nil, fmt.Errorf("config file not found: %s", err)
		}
	}

	decodeHooks := []viper.DecoderConfigOption{
		// this allows arrays to be decoded from env var separated by ",", example: MY_VAR="value1,value2,value3"
		viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(mapstructure.TextUnmarshallerHookFunc(), mapstructure.StringToSliceHookFunc(","))),
	}

	if err := viper.Unmarshal(&cfg, decodeHooks...); err != nil {
		return nil, err
	}

	var err error
	// loads the operator private key from the keystore
	if len(cfg.OperatorKeystorePath) == 0 {
		return nil, fmt.Errorf("operator keystore path is required")
	}
	if len(cfg.OperatorKeystorePasswordPath) > 0 {
		cfg.OperatorKeystorePassword, err = crypto.ReadKeystorePasswordFromFile(cfg.OperatorKeystorePasswordPath)
		if err != nil {
			return nil, err
		}
	}
	if len(cfg.OperatorKeystorePassword) == 0 {
		return nil, fmt.Errorf("operator keystore password is required")
	}
	privKey, err := crypto.LoadPrivateKey("ECDSA", cfg.OperatorKeystorePassword, cfg.OperatorKeystorePath)
	if err != nil {
		return nil, err
	}
	cfg.OperatorPrivKey = nutils.Bytes2Hex(privKey)
	privateKey, err := ecrypto.ToECDSA(privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to convert private key to ECDSA: %w", err)
	}
	cfg.OperatorAddress = ecrypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	// loads the signer private key from the keystore
	if len(cfg.SignerECDSAKeystorePath) == 0 {
		return nil, fmt.Errorf("signer ECDSA keystore path is required")
	}
	if len(cfg.SignerECDSAKeystorePasswordPath) > 0 {
		cfg.SignerECDSAKeystorePassword, err = crypto.ReadKeystorePasswordFromFile(cfg.SignerECDSAKeystorePasswordPath)
		if err != nil {
			return nil, err
		}
	}
	if len(cfg.SignerECDSAKeystorePassword) == 0 {
		return nil, fmt.Errorf("signer ECDSA keystore password is required")
	}
	privKey, err = crypto.LoadPrivateKey("ECDSA", cfg.SignerECDSAKeystorePassword, cfg.SignerECDSAKeystorePath)
	if err != nil {
		return nil, err
	}
	privateKey, err = ecrypto.ToECDSA(privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to convert private key to ECDSA: %w", err)
	}
	cfg.SignerAddress = ecrypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	// loads the BLS private key from the keystore
	if len(cfg.BLSKeystorePath) == 0 {
		return nil, fmt.Errorf("BLS keystore path is required")
	}
	if len(cfg.BLSKeystorePasswordPath) > 0 {
		cfg.BLSKeystorePassword, err = crypto.ReadKeystorePasswordFromFile(cfg.BLSKeystorePasswordPath)
		if err != nil {
			return nil, err
		}
	}
	if len(cfg.BLSKeystorePassword) == 0 {
		return nil, fmt.Errorf("BLS keystore password is required")
	}
	privKey, err = crypto.LoadPrivateKey(crypto.CryptoCurve(cfg.BLSCurve), cfg.BLSKeystorePassword, cfg.BLSKeystorePath)
	if err != nil {
		return nil, err
	}
	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(cfg.BLSCurve))
	pubKey, err := blsScheme.GetPublicKey(privKey, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get BLS public key: %w", err)
	}
	cfg.BLSPrivateKey = nutils.Bytes2Hex(privKey)
	cfg.BLSPublicKey = nutils.Bytes2Hex(pubKey)

	return &cfg, nil
}

// GenerateNodeConfig generates the client.toml file.
func GenerateNodeConfig(clientCfg *NodeConfig, network string) (string, error) {
	// Validate the client config
	if len(clientCfg.BLSKeystorePath) == 0 {
		return "", fmt.Errorf("BLS keystore path is required")
	}
	if len(clientCfg.SignerECDSAKeystorePath) == 0 {
		return "", fmt.Errorf("signer ECDSA keystore path is required")
	}
	// Create the Client Config file
	tmplClient, err := template.New("client").Parse(nodeConfigTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse client config template: %s", err)
	}
	configFileName := fmt.Sprintf("client_%s_%s_%s.toml", network, clientCfg.ChainName, clientCfg.BLSPubKey[:12])
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %s", err)
	}
	err = os.MkdirAll(filepath.Join(homeDir, configDir), 0700)
	if err != nil {
		return "", fmt.Errorf("failed to create config directory: %s", err)
	}
	configFilePath := filepath.Clean(filepath.Join(homeDir, configDir, configFileName))
	clientConfigFile, err := os.Create(configFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to create client config file: %s", err)
	}
	defer clientConfigFile.Close()
	if err := tmplClient.Execute(clientConfigFile, clientCfg); err != nil {
		return "", fmt.Errorf("failed to execute client config template: %s", err)
	}

	return configFilePath, nil
}

// LoadNodeConfig loads the node config from the client.toml file.
func LoadNodeConfig(nodeConfigFilePath string) (*NodeConfig, error) {
	var cfg struct {
		Client NodeConfig
	}

	viper.SetConfigFile(nodeConfigFilePath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	decodeHooks := []viper.DecoderConfigOption{
		// this allows arrays to be decoded from env var separated by ",", example: MY_VAR="value1,value2,value3"
		viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(mapstructure.TextUnmarshallerHookFunc(), mapstructure.StringToSliceHookFunc(","))),
	}
	if err := viper.Unmarshal(&cfg, decodeHooks...); err != nil {
		return nil, err
	}

	return &cfg.Client, nil
}
