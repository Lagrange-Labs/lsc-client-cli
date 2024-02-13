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
	LagrangeServiceSCAddr string `mapstructure:"LagrangeServiceSCAddress"`
	StakeManagerAddr      string `mapstructure:"StakeManagerAddress"`
	GRPCURL               string `mapstructure:"LagrangeServerGrpcURL"`
	RPCEndpoint           string `mapstructure:"ChainRPCEndpoint"`
	EthereumRPCURL        string `mapstructure:"EthereumRPCURL"`
	BLSCurve              string `mapstructure:"BLSCurve"`
	DockerImageTag        string `mapstructure:"DockerImageTag"`
	WETHAddr              string `mapstructure:"WETHAddr"`
}

// ClientConfig is the configuration for the lagrange client.
// This is used to generate the client.toml file.
type ClientConfig struct {
	ChainName          string `json:"chain_name"`
	ServerGrpcURL      string `json:"server_grpc_url"`
	RPCEndpoint        string `json:"rpc_endpoint"`
	EthereumRPCURL     string `json:"ethereum_rpc_url"`
	CommitteeSCAddress string `json:"committee_sc_address"`
	BLSPrivateKey      string `json:"bls_private_key"`
	ECDSAPrivateKey    string `json"ecdsa_private_key"`
	BLSCurve           string `json:"bls_curve"`
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
