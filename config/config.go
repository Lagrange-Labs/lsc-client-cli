package config

import (
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

const (
	// FlagCfg is the flag for cfg.
	FlagCfg = "config"
)

// Config is the configuration for CLI.
type Config struct {
	CommitteeSCAddr       string `mapstructure:"CommitteeSCAddress"`
	OperatorPrivKey       string `mapstructure:"OperatorPrivateKey"`
	LagrangeServiceSCAddr string `mapstructure:"LagrangeServiceSCAddress"`
	EthereumRPCURL        string `mapstructure:"EthereumRPCURL"`
	BLSCurve              string `mapstructure:"BLSCurve"`
	DockerImageTag        string `mapstructure:"DockerImageTag"`
	ConcurrentFetchers    int    `mapstructure:"ConcurrentFetchers"`
}

// ClientConfig is the configuration for the lagrange client.
// This is used to generate the client.toml file.
type ClientConfig struct {
	ChainName          string `json:"chain_name"`
	ServerGrpcURL      string `json:"server_grpc_url"`
	OperatorAddress    string `json:"operator_address"`
	L1RPCEndpoint      string `json:"l1_rpc_endpoint"`
	L2RPCEndpoint      string `json:"l2_rpc_endpoint"`
	BeaconURL          string `json:"beacon_url"`
	EthereumRPCURL     string `json:"ethereum_rpc_url"`
	CommitteeSCAddress string `json:"committee_sc_address"`
	BLSPrivateKey      string `json:"bls_private_key"`
	SignPrivateKey     string `json:"sign_private_key"`
	BatchInbox         string `json:"batch_inbox"`
	BatchSender        string `json:"batch_sender"`
	BLSCurve           string `json:"bls_curve"`
	ConcurrentFetchers int    `json:"concurrent_fetchers"`
}

// DockerComposeConfig is the configuration for the docker-compose.yml file.
type DockerComposeConfig struct {
	ChainName      string `json:"chain_name"`
	BLSPubKey      string `json:"bls_pub_key"`
	DockerImage    string `json:"docker_image"`
	ConfigFilePath string `json:"config_file_path"`
}

// Load loads the configuration
func Load(ctx *cli.Context) (*Config, error) {
	var cfg Config
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
