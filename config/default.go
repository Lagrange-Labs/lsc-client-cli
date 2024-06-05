package config

const (
	// FlagCfg is the flag for cfg.
	FlagCfg = "config"

	clientConfigTemplate = `[Client]
GrpcURLs = "{{.ServerGrpcURL}}"
Chain = "{{.ChainName}}"
EthereumURL = "{{.EthereumRPCURL}}"
OperatorAddress = "{{.OperatorAddress}}"
CommitteeSCAddress = "{{.CommitteeSCAddress}}"
BLSKeystorePath = "{{.BLSKeystorePath}}"
BLSKeystorePasswordPath = "{{.BLSKeystorePasswordPath}}"
SignerECDSAKeystorePath = "{{.SignerECDSAKeystorePath}}"
SignerECDSAKeystorePasswordPath = "{{.SignerECDSAKeystorePasswordPath}}"
PullInterval = "1000ms"
BLSCurve = "{{.BLSCurve}}"

[RpcClient]
	{{ if eq .ChainName "optimism" "base" }}
	[RpcClient.Optimism]
	RPCURL = "{{.L2RPCEndpoint}}"
	L1RPCURL = "{{.L1RPCEndpoint}}"
	BeaconURL = "{{.BeaconURL}}"
	BatchInbox = "{{.BatchInbox}}"
	BatchSender = "{{.BatchSender}}"
	ConcurrentFetchers = "{{.ConcurrentFetchers}}"
	{{ else if eq .ChainName "arbitrum" }}
	[RpcClient.Arbitrum]
	RPCURL = "{{.L2RPCEndpoint}}"
	L1RPCURL = "{{.L1RPCEndpoint}}"
	BeaconURL = "{{.BeaconURL}}"
	BatchInbox = "{{.BatchInbox}}"
	ConcurrentFetchers = "{{.ConcurrentFetchers}}"
	{{ else if eq .ChainName "mock" }}
	[RpcClient.Mock]
	RPCURL = "{{.L2RPCEndpoint}}"
	{{ end }}
[Telemetry]
	ServiceName = "{{.MetricsServiceName}}"
	PrometheusRetentionTime = "{{.PrometheusRetentionTime}}"`

	configDir = ".lagrange/config"
)

var (
	NetworkConfigs = map[string]NetworkConfig{
		"holesky": {
			ChainID:                  17000,
			CommitteeSCAddress:       "0x57BAf26C77BBBa3D3A8Bd620c8d74B44Bfda8818",
			LagrangeServiceSCAddress: "0x18A74E66cc90F0B1744Da27E72Df338cEa0A542b",
			GRPCServerURLs: map[string]string{
				"optimism": "44.210.11.64:9090",
				"base":     "3.209.124.237:9090",
				"arbitrum": "18.211.62.223:9090",
			},
		},
		"mainnet": {
			ChainID:                  1,
			CommitteeSCAddress:       "0xECc22f3EcD0EFC8aD77A78ad9469eFbc44E746F5",
			LagrangeServiceSCAddress: "0x35F4f28A8d3Ff20EEd10e087e8F96Ea2641E6AA2",
			GRPCServerURLs: map[string]string{
				"optimism": "34.202.191.166:9090",
				"base":     "34.193.82.90:9090",
				"arbitrum": "44.208.119.151:9090",
			},
		},
		"local": {
			ChainID:                  1337,
			CommitteeSCAddress:       "0xBF4E09354df24900e3d2A1e9057a9F7601fbDD06",
			LagrangeServiceSCAddress: "0xBda41273d671bb33374F7b9C4Ae8746c712727f7",
			GRPCServerURLs: map[string]string{
				"mock": "localhost:9090",
			},
		},
	}

	ChainBatchConfigs = map[string]ChainBatchConfig{
		"optimism": {
			ChainID:     10,
			BatchInbox:  "0xFF00000000000000000000000000000000000010",
			BatchSender: "0x6887246668a3b87F54DeB3b94Ba47a6f63F32985",
		},
		"base": {
			ChainID:     8453,
			BatchInbox:  "0xFf00000000000000000000000000000000008453",
			BatchSender: "0x5050F69a9786F081509234F1a7F4684b5E5b76C9",
		},
		"arbitrum": {
			ChainID:    42161,
			BatchInbox: "0x1c479675ad559DC151F6Ec7ed3FbF8ceE79582B6",
		},
		"mock": {
			ChainID: 1337,
		},
	}
)

// NetworkConfig is the configuration for the network (Holesky/Mainnet).
type NetworkConfig struct {
	ChainID                  uint32            `json:"chain_id"`
	CommitteeSCAddress       string            `json:"committee_sc_address"`
	LagrangeServiceSCAddress string            `json:"lagrange_service_sc_address"`
	GRPCServerURLs           map[string]string `json:"grpc_server_urls"`
}

// ChainBatchConfig is the batch configuration for the rollup chain batch.
type ChainBatchConfig struct {
	ChainID     uint32 `json:"chain_id"`
	BatchInbox  string `json:"batch_inbox"`
	BatchSender string `json:"batch_sender"`
}
