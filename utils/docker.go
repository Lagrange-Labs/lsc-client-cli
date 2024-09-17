package utils

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Lagrange-Labs/client-cli/config"
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	"github.com/Lagrange-Labs/lagrange-node/signer"
)

const dockerComposeTemplate = `version: "3.7"

services:
  lagrange_client_{{.Network}}_{{.ChainName}}:
    container_name: lagrange_{{.Network}}_{{.ChainName}}_{{.BLSPubKeyPrefix}}
    image: {{.DockerImage}}
    restart: always
    ports:
      - "{{.HostBindingPort}}:{{.PrometheusPort}}"
    command:
      - "/bin/sh"
      - "-c"
      - "/app/lagrange-node run-client -c /app/config/client.toml"
    volumes:
      - {{.ConfigFilePath}}:/app/config/client.toml
      - {{.CertConfig.CACertPath}}:/app/config/ca.crt
      - {{.CertConfig.NodeKeyPath}}:/app/config/node.key
      - {{.CertConfig.NodeCertPath}}:/app/config/node.crt
      - lagrange_{{.Network}}_{{.ChainName}}_{{.BLSPubKeyPrefix}}:$HOME/.lagrange
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "10"

volumes:
  lagrange_{{.Network}}_{{.ChainName}}_{{.BLSPubKeyPrefix}}:
`

const dockerSignerTemplate = `version: "3.7"
services:
  lagrange_signer:
    container_name: lagrange_signer
    image: {{.DockerImage}}
    restart: always
    ports:
      - "{{.ServerPort}}:{{.ServerPort}}"
    volumes:
      - {{.ConfigFilePath}}:/app/config/config_signer.toml
      {{ range $key, $value := .KeyStorePaths }}
      - {{$key}}:{{$value}}
      {{ end }}
      {{ range $key, $value := .PasswordPaths }}
      - {{$key}}:{{$value}}
      {{ end }}
      {{ range $key, $value := .CertPaths }}
      - {{$key}}:{{$value}}
      {{ end }}
    command:
      - "/bin/sh"
      - "-c"
      - "/app/lagrange-node run-signer -c /app/config/config_signer.toml"
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "10"
`

// CheckDockerImageExists checks if a Docker image exists locally.
func CheckDockerImageExists(imageName string) error {
	cmd := exec.Command("docker", "image", "inspect", imageName)
	if err := cmd.Run(); err != nil {
		if err := pullDockerImage(imageName); err != nil {
			return fmt.Errorf("failed to pull docker image %s: %s", imageName, err)
		}
	}
	cmd = exec.Command("docker", "image", "inspect", imageName, "--format='{{index .RepoDigests 0}}'")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get docker image digest: %s", err)
	}
	logger.Infof("Docker image %s exists locally with digest: %s", imageName, output)

	return nil
}

func pullDockerImage(imageName string) error {
	cmd := exec.Command("docker", "pull", imageName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// GenerateSignerConfigFile generates a signer service configuration file and docker-compose file.
func GenerateSignerConfigFile(cfg *signer.Config, imageName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %s", err)
	}
	workDir := filepath.Join(homeDir, ".lagrange")
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create work directory: %s", err)
	}

	var signerConfig config.DockerSignerConfig
	signerConfig.DockerImage = imageName
	signerConfig.ServerPort = cfg.GRPCPort

	signerConfig.KeyStorePaths = map[string]string{}
	signerConfig.PasswordPaths = map[string]string{}
	for _, provider := range cfg.ProviderConfigs {
		if provider.Type != "local" {
			return "", fmt.Errorf("unsupported provider type: %s", provider.Type)
		}
		keyPath := fmt.Sprintf("/app/config/%s.key", filepath.Base(provider.LocalConfig.AccountID))
		passPath := fmt.Sprintf("/app/config/%s.pass", filepath.Base(provider.LocalConfig.AccountID))
		signerConfig.KeyStorePaths[provider.LocalConfig.PrivateKeyPath] = keyPath
		signerConfig.PasswordPaths[provider.LocalConfig.PasswordKeyPath] = passPath
		provider.LocalConfig.PrivateKeyPath = keyPath
		provider.LocalConfig.PasswordKeyPath = passPath
	}

	signerConfig.CertPaths = map[string]string{}
	if cfg.TLSConfig != nil && len(cfg.TLSConfig.CACertPath) > 0 {
		certPath := "/app/config/ca-crt.pem"
		serverKeyPath := "/app/config/server-key.pem"
		serverCertPath := "/app/config/server-crt.pem"
		signerConfig.CertPaths[cfg.TLSConfig.CACertPath] = certPath
		signerConfig.CertPaths[cfg.TLSConfig.NodeKeyPath] = serverKeyPath
		signerConfig.CertPaths[cfg.TLSConfig.NodeCertPath] = serverCertPath
		cfg.TLSConfig.CACertPath = certPath
		cfg.TLSConfig.NodeKeyPath = serverKeyPath
		cfg.TLSConfig.NodeCertPath = serverCertPath
	}

	signerConfigFilePath := filepath.Join(workDir, "config/config_signer.toml")
	signerConfig.ConfigFilePath = signerConfigFilePath
	if err := config.WriteSignerConfig(cfg, signerConfigFilePath); err != nil {
		return "", err
	}

	tmpSigner, err := template.New("docker-signer").Parse(dockerSignerTemplate)
	if err != nil {
		return "", err
	}
	signerDockerFilePath := filepath.Join(workDir, "docker-compose_signer.yml")
	signerDockerFile, err := os.Create(signerDockerFilePath)
	if err != nil {
		return "", err
	}
	defer signerDockerFile.Close()
	if err := tmpSigner.Execute(signerDockerFile, signerConfig); err != nil {
		return "", err
	}

	logger.Infof("Generated Signer service configuration file: %s", signerConfigFilePath)
	logger.Infof("Generated Docker Compose file for Signer service: %s", signerDockerFilePath)
	return signerDockerFilePath, nil
}

// GenerateDockerComposeFile generates a Docker Compose file.
func GenerateDockerComposeFile(cfg *config.CLIConfig, imageName, nodeConfigFilePath string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %s", err)
	}
	workDir := filepath.Join(homeDir, ".lagrange")
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create work directory: %s", err)
	}

	var dockerConfig config.DockerComposeConfig
	// Get chain name and bls pub key from config file path
	configFileName := filepath.Base(nodeConfigFilePath)
	seps := strings.Split(configFileName, "_")
	dockerConfig.Network = seps[1]
	dockerConfig.ChainName = seps[2]
	dockerConfig.BLSPubKeyPrefix = strings.Split(seps[3], ".")[0]
	dockerConfig.DockerImage = imageName
	dockerConfig.CertConfig = cfg.CertConfig
	dockerConfig.ConfigFilePath = nodeConfigFilePath
	dockerConfig.PrometheusPort = cfg.MetricsServerPort
	dockerConfig.HostBindingPort = cfg.HostBindingPort

	tmpDocker, err := template.New("docker-compose").Parse(dockerComposeTemplate)
	if err != nil {
		return "", err
	}
	dockerComposeFilePath := filepath.Join(workDir, fmt.Sprintf("docker-compose_%s_%s_%s.yml", dockerConfig.Network, dockerConfig.ChainName, dockerConfig.BLSPubKeyPrefix))
	dockerConfigFile, err := os.Create(dockerComposeFilePath)
	if err != nil {
		return "", err
	}
	defer dockerConfigFile.Close()
	if err := tmpDocker.Execute(dockerConfigFile, dockerConfig); err != nil {
		return "", err
	}

	logger.Infof("Generated Docker Compose file: %s", dockerComposeFilePath)
	return dockerComposeFilePath, nil
}

// RunClientNode runs a client node image.
func RunClientNode(cfg *config.CLIConfig, imageName, clientConfigFilePath string) error {
	dockerComposeFilePath, err := GenerateDockerComposeFile(cfg, imageName, clientConfigFilePath)
	if err != nil {
		return err
	}

	return RunDockerImage(dockerComposeFilePath)
}

// RunDockerImage runs a Docker image.
func RunDockerImage(dockerComposeFilePath string) error {
	cmd := exec.Command("docker", "compose", "-f", dockerComposeFilePath, "up", "-d")
	if err := cmd.Start(); err != nil {
		return err
	}
	logger.Infof("Started Docker container with docker-compose file: %s", dockerComposeFilePath)
	// Detach from the container
	if err := cmd.Process.Release(); err != nil {
		return err
	}
	return nil
}
