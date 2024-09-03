package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
)

// CLIConfig is the configuration for the lagrange CLI.
type CLIConfig struct {
	CertConfig              *core.CertConfig `mapstructure:"CertConfig"`
	SignerServerURL         string           `mapstructure:"SignerServerURL"`
	OperatorAddress         string           `mapstructure:"OperatorAddress"`
	OperatorKeyAccountID    string           `mapstructure:"OperatorKeyAccountID"`
	SignerKeyAccountID      string           `mapstructure:"SignerKeyAccountID"`
	BLSKeyAccountID         string           `mapstructure:"BLSKeyAccountID"`
	EthereumRPCURL          string           `mapstructure:"EthereumRPCURL"`
	L1RPCEndpoint           string           `mapstructure:"L1RPCEndpoint"`
	BeaconURL               string           `mapstructure:"BeaconURL"`
	L2RPCEndpoint           string           `mapstructure:"L2RPCEndpoint"`
	BLSCurve                string           `mapstructure:"BLSCurve"`
	ConcurrentFetchers      int              `mapstructure:"ConcurrentFetchers"`
	MetricsEnabled          bool             `mapstructure:"MetricsEnabled"`
	MetricsServerPort       string           `mapstructure:"MetricsServerPort"`
	HostBindingPort         string           `mapstructure:"HostBindingPort"`
	MetricsServiceName      string           `mapstructure:"MetricsServiceName"`
	PrometheusRetentionTime string           `mapstructure:"PrometheusRetentionTime"`
}

// NodeConfig is the configuration for the LSC node.
// This is used to generate the client.toml file.
type NodeConfig struct {
	CertConfig              *core.CertConfig
	ChainName               string
	ServerGrpcURL           string
	SignerServerURL         string
	SignerKeyAccountID      string
	BLSKeyAccountID         string
	OperatorAddress         string
	L1RPCEndpoint           string
	L2RPCEndpoint           string
	BeaconURL               string
	EthereumRPCURL          string
	CommitteeSCAddress      string
	BatchInbox              string
	BatchSender             string
	BLSCurve                string
	ConcurrentFetchers      int
	MetricsEnabled          bool
	MetricsServerPort       string
	MetricsServiceName      string
	PrometheusRetentionTime string
}

// DockerComposeConfig is the configuration for the docker-compose.yml file.
type DockerComposeConfig struct {
	Network         string           `json:"network"`
	ChainName       string           `json:"chain_name"`
	BLSPubKeyPrefix string           `json:"bls_pub_key"`
	CertConfig      *core.CertConfig `json:"cert_config"`
	DockerImage     string           `json:"docker_image"`
	ConfigFilePath  string           `json:"config_file_path"`
	PrometheusPort  string           `json:"prometheus_port"`
	HostBindingPort string           `json:"host_binding_port"`
}

// DockerSignerConfig is the configuration for the signer service in the docker-compose.yml file.
type DockerSignerConfig struct {
	DockerImage    string            `json:"docker_image"`
	ServerPort     string            `json:"server_port"`
	ConfigFilePath string            `json:"config_file_path"`
	KeyStorePaths  map[string]string `json:"keystore_paths"`
	PasswordPaths  map[string]string `json:"password_paths"`
	CertPaths      map[string]string `json:"cert_paths"`
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

	return &cfg, nil
}

// GenerateNodeConfig generates the client.toml file.
func GenerateNodeConfig(clientCfg *NodeConfig, network string) (string, error) {
	// Create the Client Config file
	tmplClient, err := template.New("client").Parse(nodeConfigTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse client config template: %s", err)
	}
	configFileName := fmt.Sprintf("client_%s_%s_%s.toml", network, clientCfg.ChainName, clientCfg.BLSKeyAccountID)
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
